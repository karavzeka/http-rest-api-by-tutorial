package sqlstore

import (
	"database/sql"
	"github.com/karavzeka/http-rest-api-by-tutorial/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db 	           *sql.DB
	userRepository *UserRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store {
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if (s.userRepository == nil) {
		s.userRepository = &UserRepository{
			store: s,
		}
	}
	return s.userRepository
}