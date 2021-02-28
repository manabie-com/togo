package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func SetupPostgres(connection string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}