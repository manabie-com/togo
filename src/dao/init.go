package dao

import (
	"github.com/HoangMV/togo/lib/pgsql"
	"github.com/jmoiron/sqlx"
)

type DAO struct {
	db *sqlx.DB
}

func New() *DAO {
	return &DAO{pgsql.Get()}
}
