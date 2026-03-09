FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY src ./src
RUN go build -o /bin/gateway ./src

FROM alpine:3.20
RUN adduser -D -u 10001 app
USER app
COPY --from=build /bin/gateway /bin/gateway
ENTRYPOINT ["/bin/gateway"]
