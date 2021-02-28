package mock

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
)

func SetupMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Panicf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TeardownMock(db *sql.DB) {
	db.Close()
}
