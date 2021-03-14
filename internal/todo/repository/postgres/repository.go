package postgres

import "github.com/jmoiron/sqlx"

type PGRepository struct {
	DBConn *sqlx.DB
}
