package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=restapi_dev password='12345' host=localhost dbname=restapi_dev sslmode=disable"
	}

	os.Exit(m.Run())
}