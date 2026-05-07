package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andrei-galkin/ScimRegistry/internal/api"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", api.HealthHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Starting ScimRegistry on :8080...")
	fmt.Println("Health check available at http://localhost:8080/health")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
