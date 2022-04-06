package driver

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const (
	maxOpenDbConns  = 10
	maxIdleConns    = 5
	connMaxLifeTime = 5 * time.Minute
)

var dbConn = &DB{}

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

// ConnectDB creates a database connection pool for PostgreSql
func ConnectDB(dsn string) (*DB, error) {
	d, err := newPostgresDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConns)
	d.SetMaxIdleConns(maxIdleConns)
	d.SetConnMaxLifetime(connMaxLifeTime)

	dbConn.SQL = d

	return dbConn, nil
}

func newPostgresDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
