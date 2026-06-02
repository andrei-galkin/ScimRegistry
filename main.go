package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andrei-galkin/ScimRegistry/internal/api"
	"github.com/andrei-galkin/ScimRegistry/internal/store"
)

func main() {
	// Initialize the In-Memory Storage
	storage := store.NewMemStore()

	// Initialize the API Server with the storage dependency
	scimApi := api.NewServer(storage)

	// Setup the Request Multiplexer
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/health", scimApi.HealthHandler)
	mux.HandleFunc("/Schemas", scimApi.SchemasHandler)

	// Example of a route that needs storage (we'll build this handler next)
	// mux.HandleFunc("POST /v2/Users", scimApi.CreateUserHandler)

	// Configure the HTTP Server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("SCIM Registry starting on :8080...")
	fmt.Println("Health check: http://localhost:8080/health")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Critical: Server failed to start: %v", err)
	}
}
