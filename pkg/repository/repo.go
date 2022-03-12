package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Insert()
}
type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (t *repository) Insert() {

}
