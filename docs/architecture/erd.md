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
