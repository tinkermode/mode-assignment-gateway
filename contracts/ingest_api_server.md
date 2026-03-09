# Ingest API Contract

## Endpoint
`POST /ingest`

Default runtime target in this assignment:
- `INGEST_HOST=ingest-api-server`
- `INGEST_PORT=8080`

The IoT Gateway should call:
- `http://${INGEST_HOST}:${INGEST_PORT}/ingest`

## Request Body (JSON)
Required fields:
- `seq` (int)
- `temp` (float)
- `humid` (float)
- `timestamp` (RFC3339 string)

Example:
```json
{
  "seq": 1042,
  "temp": 23.1,
  "humid": 52.0,
  "timestamp": "2026-03-05T08:00:00Z"
}
```

## Headers
Recommended and expected in this environment:
- `Idempotency-Key`: stable per record, e.g. `session_id:seq` (for example `7:1042`)

## Responses
- `200 OK`: accepted
- `400 Bad Request`: malformed JSON or unknown fields
- `422 Unprocessable Entity`: syntactically valid JSON but semantically invalid payload
- `429 Too Many Requests`: includes `Retry-After: <seconds>`
- `503 Service Unavailable`: transient overload/chaos condition

## Rate Limit Specification
- Rate limiting is enforced by the Ingest API Server using a token-bucket model.
- The limiter applies at server level (shared capacity), not per-client.
- Default behavior in this assignment:
  - steady-state capacity is about `10 requests/second`
  - short bursts are allowed up to about `10 requests`
  - when rate-limited, `Retry-After` is `1` second
- When capacity is exhausted, the Ingest API Server returns:
  - `429 Too Many Requests`
  - `Retry-After: <seconds>`

## Validation Rules (for `422`)
- `seq` must be a positive integer (`> 0`)
- `temp` must be a finite number (not `NaN`/`Inf`)
- `humid` must be a finite number in range `0..100`
- `timestamp` must be non-empty and parseable as RFC3339/RFC3339Nano

The IoT Gateway should retry transient failures with bounded backoff.
