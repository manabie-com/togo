package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteConnector(dbPath string) Connector {
	return &sqliteConnector{
		dbPath: dbPath,
	}
}

type sqliteConnector struct {
	db     *sql.DB
	dbPath string
}

func (p *sqliteConnector) Open(cfg *Config) error {
	db, err := sql.Open("sqlite3", p.dbPath)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

func (p *sqliteConnector) Close() error {
	return p.db.Close()
}

func (p *sqliteConnector) GetDB() *sql.DB {
	return p.db
}
