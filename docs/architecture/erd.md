# ERD

## `greetings`

One shared persisted greeting.

| Column | PostgreSQL type | Constraints | Notes |
|---|---|---|---|
| `id` | `smallint` | primary key, check (`id = 1`) | Enforces one row. |
| `text` | `text` | not null, check (`btrim(text) <> ''`) | Stored trimmed greeting. |
| `updated_at` | `timestamptz` | not null, default `now()` | Last successful save. |

Relationships: none. `greetings` is intentionally singleton. Initial migration inserts `(1, 'Hello, World!')` with conflict-safe seed logic.

Migration compatibility: add columns with safe defaults or nullable first. Do not change singleton constraint without product scope change.

## Migration plan — Persisted Editable Greeting

Forward: startup migration creates `greetings` if absent with all constraints, then seeds `(1, 'Hello, World!')` using `INSERT ... ON CONFLICT (id) DO NOTHING`. Existing populated singleton rows remain unchanged. Safe on populated databases: yes; schema creation is conditional and seed is non-destructive.

Backward: drop `greetings` only when reverting a fresh acceptance environment. This deletes saved public greeting content and is unsafe on populated databases; production rollback must deploy prior application code while retaining table and row.
