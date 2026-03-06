# Sensor Stream Contract

## Transport
- Protocol: TCP
- Address: `${SENSOR_HOST}:${SENSOR_PORT}`
- Data format: newline-delimited text (`\n`)

Default runtime endpoint in this assignment:
- `SENSOR_HOST=sensor`
- `SENSOR_PORT=7000`

Notes:
- In `compose.yml`, `sensor` is the Docker Compose service name of the simulator.
- Gateway should connect to `sensor:7000` from inside the Compose network.

## Semantics
- TCP is a byte stream: message boundaries are not preserved.
- A read can contain:
  - a partial line,
  - one full line,
  - multiple full lines.
- Sender may disconnect; client should reconnect.

## Record Format
Each line is comma-separated `key=value` pairs.

Example:
```text
seq=1042,temp=23.10,humid=52.00
```

Fields:
- `seq` (int): monotonically increasing per sensor session.
- `temp` (float)
- `humid` (float)

## Malformed Inputs
Malformed lines may appear, e.g.:
- missing keys
- non-numeric values
- truncated lines

Gateway should log parse errors and continue.
