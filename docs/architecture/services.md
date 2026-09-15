# Service Contracts

Base path reaches backend after proxy strips `/api`. Paths below deliberately omit `/api`.

## Error envelope

All non-2xx JSON responses:

```json
{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}
```

| HTTP | Code | Message | When |
|---|---|---|---|
| 400 | `MALFORMED_REQUEST` | `Request body is malformed.` | Bad JSON, wrong JSON type, or unknown field. |
| 422 | `VALIDATION_FAILED` | `Greeting must not be empty.` | `text` trims to empty. |
| 500 | `INTERNAL` | `Internal server error.` | Query or unexpected server failure. |
| 503 | `UNAVAILABLE` | `Service unavailable.` | Database dependency unavailable. |

## Greeting

Mock review: `code/frontend/lib/mock/persisted-editable-greeting.ts` exposes `Greeting` as `{ "text": string }`; its localStorage mechanics are UI-only. Contract retains this sound response shape, so frontend backend swap changes data source only. Empty client validation remains UI behavior; API repeats validation at trust boundary.

### `GET /v1/greeting`

No authentication. Returns current singleton greeting.

Success `200`:

```json
{"text":"Hello, World!"}
```

Errors: `INTERNAL`, `UNAVAILABLE`.

### `PUT /v1/greeting`

No authentication. Replaces greeting. Request must be JSON object with only `text` string. Server trims leading and trailing whitespace before validation and storage. Last successful save wins.

```json
{"text":"Hi pipeline"}
```

Success `200`:

```json
{"text":"Hi pipeline"}
```

Errors: `MALFORMED_REQUEST`, `VALIDATION_FAILED`, `INTERNAL`, `UNAVAILABLE`.
