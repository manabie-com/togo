package sqlstore

import (
	"database/sql"
)

// Store for working with postgres
type Store struct {
	DB *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}

func (s *Store) GetDB() *sql.DB {
	return s.DB
}

func (s *Store) Close() error {
	return s.DB.Close()
}
