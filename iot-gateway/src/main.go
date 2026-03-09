package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func getenv(key, def string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return def
	}
	return v
}

func main() {
	log.SetFlags(0)
	sensorHost := getenv("SENSOR_HOST", "sensor-tcp-simulator")
	sensorPort := getenv("SENSOR_PORT", "7000")
	ingestHost := getenv("INGEST_HOST", "ingest-api-server")
	ingestPort := getenv("INGEST_PORT", "8080")
	ingestURL := fmt.Sprintf("http://%s:%s/ingest", ingestHost, ingestPort)

	log.Printf(`{"event":"startup","sensor_host":"%s","sensor_port":"%s","ingest_url":"%s"}`,
		sensorHost, sensorPort, ingestURL)
	log.Printf(`{"event":"todo","message":"implement stream ingestion, buffering, retry, and observability"}`)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		log.Printf(`{"event":"heartbeat","status":"template_running"}`)
	}
}
