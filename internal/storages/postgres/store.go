package postgres

import (
	"database/sql"
)

// Store defines all functions to execute db queries
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL Queries
type SQLStore struct {
	DB *sql.DB
	*Queries
}

// NewStore creates new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		DB:      db,
		Queries: New(db),
	}
}
