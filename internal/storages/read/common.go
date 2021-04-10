package read

import "github.com/go-pg/pg"

const (
	defaultLimit = 20
)

type readRepo struct {
	db *pg.DB
}

func NewReadRepo(db *pg.DB) *readRepo {
	return &readRepo{db: db}
}
