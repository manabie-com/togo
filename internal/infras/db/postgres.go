package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPGDB(dbUrl string) (*sql.DB, error) {
	database, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	return database, nil
}
