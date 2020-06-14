package teststore_test

import (
	"github.com/karavzeka/http-rest-api-by-tutorial/internal/app/store"
	"github.com/karavzeka/http-rest-api-by-tutorial/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"github.com/karavzeka/http-rest-api-by-tutorial/model"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}