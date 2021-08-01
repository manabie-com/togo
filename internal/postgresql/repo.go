package postgresql

import "database/sql"

type Repo struct {
	*Queries
	*sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB:      db,
		Queries: New(db),
	}
}
