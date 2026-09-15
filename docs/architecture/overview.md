# Architecture Overview

## Stack

| Part | Choice | Reason |
|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Required UI stack; standalone output fits supplied image. |
| Backend | Go 1.22, `net/http`, pgx | Small HTTP service; pgx provides PostgreSQL driver. |
| Database | PostgreSQL 16 | Required persisted shared greeting. |

## Layout

- `code/backend/cmd/api`: only executable; starts migrations and HTTP server.
- `code/backend/migrations`: ordered SQL migrations embedded with backend binary.
- `code/frontend/app`: App Router shell and immutable shared design tokens.
- `code/frontend/components`: one default-export component per story.
- `code/frontend/lib`: story API clients; temporary fixtures live in `lib/mock`.

## Contracts and conventions

- Backend reads `DATABASE_URL`, migrates database, then listens on `PORT`, `APP_PORT`, or `8080`.
- `/healthz` returns 200 only after migrations and database `SELECT 1` succeed.
- Migrations use timestamped `*.up.sql`/`*.down.sql`; `schema_migrations` records applied versions. Concurrent-index migrations run outside transaction.
- API routes use `/v1/...`; deployment proxy owns `/api` prefix stripping.
- Go names use standard Go casing. React components use PascalCase default exports. CSS modules use tokens from `app/globals.css`; no visual literals or token fallbacks.
- `app/page.tsx` remains server composition root. Interactive story components start with literal first line `"use client"`.

## Environment

| Location | Keys |
|---|---|
| `code/backend/.env.example` | `DATABASE_URL`, `PORT`, `APP_PORT` |
| `code/frontend/.env.example` | `NEXT_PUBLIC_API_URL`, `API_ORIGIN` |
| `.env.example` | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, ports, API origins |

No external services or secrets beyond local PostgreSQL credentials. Validate request JSON and greeting rules at API boundary; use parameterized queries. Public greeting has no user data.

## Run

1. Copy `.env.example` to `.env` and choose non-production local values.
2. Run `docker compose --profile local up --build` from repository root.
3. Open `http://localhost:3000`; backend health endpoint is `http://localhost:8080/healthz`.

## Decisions

| Decision | Rejected | Tradeoff |
|---|---|---|
| One shared greeting row | Accounts, history, multiple rows | Matches approved scope; no per-user isolation or audit trail. |
| Boot-time migrations | Manual migration command | Empty runtime DB works automatically; startup depends on DB availability. |
| SQL migrations + pgx | ORM | One-table scope needs no abstraction; future complex queries remain explicit SQL. |

Unknown: greeting maximum length is not specified. First story should preserve valid non-empty trimmed text without inventing a limit. Rollout: migration seed uses id `1`; repeated boots are idempotent.
