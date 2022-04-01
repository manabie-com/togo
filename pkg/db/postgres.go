package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// New creates new PostgreSQL database connection
func New(dbPsn string, enableLog bool) (*gorm.DB, error) {
	return new("postgres", dbPsn, enableLog)
}

// new creates supported database connection (PostgreSQL currently)
func new(dialect, dbPsn string, enableLog bool) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, dbPsn)
	if err != nil {
		return nil, err
	}

	db.LogMode(enableLog == true)

	return db, nil
}
