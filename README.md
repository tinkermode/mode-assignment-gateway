# Gateway Take-Home Assignment

## Objective

Build a resilient gateway service that reads newline-delimited sensor data from a TCP stream and forwards it to an HTTP ingest API.

## Timebox

Target: **1-6 hours**.

## What You Implement

Implement only the `gateway` service in `gateway/src/`.

Go is recommended, but any language is fine as long as it runs in the provided Docker environment.

The `sensor` and `server` services are provided by MODE and should not be modified. You can run them locally via `docker compose` using the provided `compose.yml`.

You can use any AI tools or resources to help with the implementation, but you must document your usage in `AI_USAGE.md`.

## How to submit

1. Fork this repository to your own GitHub account. Make sure to make your forked repository is **private**.
2. Implement the gateway service in your forked repository.
3. Once you are done, share the repository with the `tinkermode` GitHub account (read access is sufficient).

## Deliverables

1. Gateway implementation.
2. Short README section describing architecture/tradeoffs.
3. Completed `AI_USAGE.md`.

## Inputs and Contracts

- Sensor protocol: `contracts/sensor_stream.md`
- Ingest API: `contracts/ingest_api.md`
- Examples: `contracts/examples/`

## Run

```bash
docker compose up --build
```

## Requirements

- Read the following env vars and use them to connect to the sensor and ingest API server:
  - `SENSOR_HOST` (default `sensor`)
  - `SENSOR_PORT` (default `7000`)
  - `INGEST_HOST` (default `server`)
  - `INGEST_PORT` (default `8080`)
- The gateway process should run by calling `docker compose up` from the candidate repo root.

## Notes

- Sensor emulator is a TCP byte stream and uses chunked delivery by default (partial lines / multiple lines per read are possible).
- Sensor emulator may produce random bursts and random disconnects.
- Ingestion server enforces constant rate limiting by default and may return `429` with `Retry-After`.
- Ingestion server may also return intermittent `5xx`.
