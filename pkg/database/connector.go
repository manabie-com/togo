package database

import (
	"database/sql"
)

type Connector interface {
	Open(cfg *Config) error
	Close() error
	GetDB() *sql.DB
}
