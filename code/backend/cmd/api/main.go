package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-2fc37493/backend/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
)

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
