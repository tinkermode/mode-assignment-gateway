# IoT Gateway Take-Home Assignment

## Overview

```
Sensor TCP Simulator ──TCP stream──► IoT Gateway ──HTTP POST──► Ingest API Server
(provided by MODE)                   (you implement)            (provided by MODE)
```

The **Sensor TCP Simulator** emits newline-delimited sensor readings over a persistent TCP connection.
Your **IoT Gateway** must read that stream and forward each record to the **Ingest API Server** via HTTP.

## Objective

Build a resilient IoT Gateway service that reads newline-delimited sensor data from a TCP stream and forwards it to an HTTP Ingest API.

## Timebox

Target: **1-6 hours**.

## What You Implement

Implement only the `iot-gateway` service in `iot-gateway/src/`.

Go is recommended, but any language is fine as long as it runs in the provided Docker environment.

The `sensor-tcp-simulator` and `ingest-api-server` services are provided by MODE and should not be modified. You can run them locally via `docker compose` using the provided `compose.yml`.

You can use any AI tools or resources to help with the implementation, but you must document your usage in `AI_USAGE.md`.

## How to submit

1. Clone or download this repository to your local environment.
2. Implement the IoT Gateway service in your forked repository.
3. Push your local repository to your GitHub account. Make sure the repository is **private**.
4. Invite the `tinkermode` GitHub account as a collaborator to your repository.

**Do not fork this repository directly, because it will be public**

## Inputs and Contracts

- Sensor TCP stream protocol: `contracts/sensor_stream.md`
- Ingest API: `contracts/ingest_api.md`
- Examples: `contracts/examples/`

Notes on runtime behavior:

- The Sensor TCP Simulator is a TCP byte stream and uses chunked delivery by default (partial lines / multiple lines per read are possible).
- The Sensor TCP Simulator may produce random bursts and random disconnects.
- The Ingest API Server enforces constant rate limiting by default and may return `429` with `Retry-After`.
- The Ingest API Server may also return intermittent `5xx`.

## Run

```bash
docker compose up --build
```

## Requirements

- The IoT Gateway must read newline-delimited sensor data from the Sensor TCP Simulator and forward it to the Ingest API Server via HTTP POST.
- The IoT Gateway process should run by calling `docker compose up`.
- Read the following env vars and use them to connect to the Sensor TCP Simulator and Ingest API Server:
  - `SENSOR_HOST` (default `sensor-tcp-simulator`)
  - `SENSOR_PORT` (default `7000`)
  - `INGEST_HOST` (default `ingest-api-server`)
  - `INGEST_PORT` (default `8080`)
- Handle error cases from both services gracefully (e.g. sensor disconnects, bursts, `429`, `5xx`). Document your decisions about how you handle these cases in `DESIGN_DECISIONS.md`.

## Deliverables

1. IoT Gateway implementation with a `Dockerfile` in `iot-gateway/` that can be built and run with `docker compose up`.
2. Completed `DESIGN_DECISIONS.md`.
3. Completed `AI_USAGE.md`.
