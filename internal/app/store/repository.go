package store

import (
	"github.com/karavzeka/http-rest-api-by-tutorial/model"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}