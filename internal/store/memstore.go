package store

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/andrei-galkin/ScimRegistry/internal/schema"
)

type MemStore struct {
	mu     sync.RWMutex
	users  map[string]*schema.User
	groups map[string]*schema.Group
}

func NewMemStore() *MemStore {
	return &MemStore{
		users:  make(map[string]*schema.User),
		groups: make(map[string]*schema.Group),
	}
}

// CreateUser saves a user and ensures no duplicate UserNames (SCIM requirement)
func (s *MemStore) CreateUser(ctx context.Context, user *schema.User) (*schema.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// SCIM specific: check for uniqueness
	for _, existing := range s.users {
		if existing.UserName == user.UserName {
			return nil, errors.New("conflict: userName already exists")
		}
	}

	s.users[user.ID] = user
	return user, nil
}

func (s *MemStore) GetUser(ctx context.Context, id string) (*schema.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %s not found", id)
	}
	return user, nil
}

func (s *MemStore) DeleteUser(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return fmt.Errorf("user %s not found", id)
	}
	delete(s.users, id)
	return nil
}

// Group implementations follow the same pattern...
func (s *MemStore) CreateGroup(ctx context.Context, g *schema.Group) (*schema.Group, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.groups[g.ID] = g
	return g, nil
}

func (s *MemStore) GetGroup(ctx context.Context, id string) (*schema.Group, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	g, ok := s.groups[id]
	if !ok {
		return nil, fmt.Errorf("group %s not found", id)
	}
	return g, nil
}
