package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres",
		dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error ping database: %w", err)
	}

	return &Store{db}, nil
}
