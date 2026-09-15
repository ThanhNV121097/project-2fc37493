package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-2fc37493/backend/migrations"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type greetingResponse struct {
	Text string `json:"text"`
}

type errorResponse struct {
	Error errorBody `json:"error"`
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer pool.Close()
	if err := migrate(ctx, pool); err != nil {
		log.Fatalf("migrate database: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("check database: %v", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := pool.Ping(ctx); err != nil {
			http.Error(w, "service unavailable", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/v1/greeting", greetingHandler(pool))
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}
	server := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	log.Fatal(server.ListenAndServe())
}

func greetingHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getGreeting(w, r, pool)
		case http.MethodPut:
			putGreeting(w, r, pool)
		default:
			w.Header().Set("Allow", "GET, PUT")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func getGreeting(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	var text string
	if err := pool.QueryRow(ctx, `SELECT text FROM greetings WHERE id = 1`).Scan(&text); err != nil {
		writeStoreError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, greetingResponse{Text: text})
}

func putGreeting(w http.ResponseWriter, r *http.Request, pool *pgxpool.Pool) {
	var raw map[string]string
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&raw); err != nil {
		writeError(w, http.StatusBadRequest, "MALFORMED_REQUEST", "Request body is malformed.")
		return
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		writeError(w, http.StatusBadRequest, "MALFORMED_REQUEST", "Request body is malformed.")
		return
	}
	text, ok := raw["text"]
	if !ok || len(raw) != 1 {
		writeError(w, http.StatusBadRequest, "MALFORMED_REQUEST", "Request body is malformed.")
		return
	}
	text = strings.TrimSpace(text)
	if text == "" {
		writeError(w, http.StatusUnprocessableEntity, "VALIDATION_FAILED", "Greeting must not be empty.")
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	if err := pool.QueryRow(ctx, `UPDATE greetings SET text = $1, updated_at = now() WHERE id = 1 RETURNING text`, text).Scan(&text); err != nil {
		writeStoreError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, greetingResponse{Text: text})
}

func writeStoreError(w http.ResponseWriter, err error) {
	var pgErr *pgconn.ConnectError
	if errors.As(err, &pgErr) || errors.Is(err, context.Canceled) {
		writeError(w, http.StatusServiceUnavailable, "UNAVAILABLE", "Service unavailable.")
		return
	}
	writeError(w, http.StatusInternalServerError, "INTERNAL", "Internal server error.")
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, errorResponse{Error: errorBody{Code: code, Message: message}})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func migrate(ctx context.Context, pool *pgxpool.Pool) error {
	if _, err := pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (version text PRIMARY KEY, applied_at timestamptz NOT NULL DEFAULT now())`); err != nil {
		return err
	}
	entries, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		var applied bool
		if err := pool.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)`, name).Scan(&applied); err != nil {
			return err
		}
		if applied {
			continue
		}
		sql, err := migrations.Files.ReadFile(name)
		if err != nil {
			return err
		}
		if strings.Contains(strings.ToUpper(string(sql)), "CREATE INDEX CONCURRENTLY") {
			if _, err := pool.Exec(ctx, string(sql)); err != nil { return err }
			if _, err := pool.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, name); err != nil { return err }
			continue
		}
		tx, err := pool.Begin(ctx)
		if err != nil { return err }
		if _, err = tx.Exec(ctx, string(sql)); err == nil {
			_, err = tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, name)
		}
		if err != nil { _ = tx.Rollback(ctx); return fmt.Errorf("apply %s: %w", name, err) }
		if err := tx.Commit(ctx); err != nil { return err }
	}
	return nil
}
