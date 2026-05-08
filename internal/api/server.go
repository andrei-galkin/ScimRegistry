package api

import (
	"github.com/andrei-galkin/ScimRegistry/internal/store"
)

// Server holds dependencies for our HTTP handlers
type Server struct {
	Storage store.Repository
}

func NewServer(s store.Repository) *Server {
	return &Server{Storage: s}
}
