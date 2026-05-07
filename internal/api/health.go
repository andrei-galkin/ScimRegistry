package api

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status  string            `json:"status"`
	Version string            `json:"version"`
	Checks  map[string]string `json:"checks"`
}

// HealthHandler returns a simple UP status for liveness/readiness checks.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "UP",
		Version: "v1.0.0",
		Checks: map[string]string{
			"check": "healthy", // Placeholder for actual other checks
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
