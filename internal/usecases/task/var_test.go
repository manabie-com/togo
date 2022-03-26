package task

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	userID        = 1
	maxTaskPerDay = 5
)

func setupMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error '%s' occured when opening a mock db connection", err)
	}

	return db, mock
}
