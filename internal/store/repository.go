package store

import (
	"context"

	"github.com/andrei-galkin/ScimRegistry/internal/schema"
)

type Repository interface {
	// User operations
	CreateUser(ctx context.Context, user *schema.User) (*schema.User, error)
	GetUser(ctx context.Context, id string) (*schema.User, error)
	DeleteUser(ctx context.Context, id string) error

	// Group operations
	CreateGroup(ctx context.Context, group *schema.Group) (*schema.Group, error)
	GetGroup(ctx context.Context, id string) (*schema.Group, error)
}
