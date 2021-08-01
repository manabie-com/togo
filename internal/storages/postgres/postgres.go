package postgres

import (
	"database/sql"
)

const (
	Postgres = "postgres"
)

func NewPostgresClient(address string) (*sql.DB, error) {
	return sql.Open(Postgres, address)
}
