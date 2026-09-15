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

### `GET /v1/greeting`

Returns current singleton greeting.

Success `200`:

```json
{"text":"Hello, World!"}
```

Errors: `INTERNAL`, `UNAVAILABLE`.

### `PUT /v1/greeting`

Replaces greeting. Request must be JSON object with only `text` string. Server trims leading and trailing whitespace before validation and storage.

```json
{"text":"Hi pipeline"}
```

Success `200`:

```json
{"text":"Hi pipeline"}
```

Errors: `MALFORMED_REQUEST`, `VALIDATION_FAILED`, `INTERNAL`, `UNAVAILABLE`.

No authentication, pagination, caching, or conflict contract. Last successful save wins.
