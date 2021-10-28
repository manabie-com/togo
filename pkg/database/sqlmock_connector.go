package database

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewSqlMockConnector() (Connector, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	return &mockConnector{
		db:    db,
		ready: false,
	}, mock, nil
}

type mockConnector struct {
	db    *sql.DB
	ready bool
}

func (p *mockConnector) Open(cfg *Config) error {
	p.ready = true
	return nil
}

func (p *mockConnector) Close() error {
	return p.db.Close()
}

func (p *mockConnector) GetDB() *sql.DB {
	if p.ready == false {
		return nil
	}
	return p.db
}
