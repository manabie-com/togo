package repository

import (
	"database/sql"
	"togo/internal/postgresql"
)

type Repo struct {
	q *postgresql.Queries
	*sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		q: postgresql.New(db),
	}
}
