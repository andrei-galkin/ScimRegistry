package api

import (
	"encoding/json"
	"net/http"

	"github.com/andrei-galkin/ScimRegistry/internal/schema"
)

type schemasListResponse struct {
	Schemas      []string        `json:"schemas"`
	TotalResults int             `json:"totalResults"`
	StartIndex   int             `json:"startIndex"`
	ItemsPerPage int             `json:"itemsPerPage"`
	Resources    json.RawMessage `json:"Resources"`
}

// SchemasHandler returns the collection of schemas supported by ScimRegistry
func (s *Server) SchemasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	response := schemasListResponse{
		Schemas:      []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"},
		TotalResults: 2,
		StartIndex:   1,
		ItemsPerPage: 2,
		Resources:    json.RawMessage(schema.RawSchemasJSON),
	}

	payload, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal schemas response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/scim+json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(payload)
}
