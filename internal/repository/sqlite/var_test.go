package sqlite

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	username      = "khxingn"
	password      = "Qq@1234567"
	wrongUsername = "wrong_username"
	wrongPassword = "wrong_password"
	userID        = 1
)

func setupMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error '%s' occured when opening a mock db connection", err)
	}

	return db, mock
}
