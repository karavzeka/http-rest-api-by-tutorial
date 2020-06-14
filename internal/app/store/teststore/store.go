package teststore

import (
	"github.com/karavzeka/http-rest-api-by-tutorial/model"
	"github.com/karavzeka/http-rest-api-by-tutorial/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if (s.userRepository == nil) {
		s.userRepository = &UserRepository{
			store: s,
			users: make(map[string]*model.User),
		}
	}
	return s.userRepository
}