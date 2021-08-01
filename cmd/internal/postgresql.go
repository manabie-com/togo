package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/url"
	"togo/config"
)

func NewPostgresql(config *config.Config) (*sql.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.DbUsername, config.DbPassword),
		Host:   fmt.Sprintf("%s:%d", config.DbHost, config.DbPort),
		Path:   config.DbName,
	}

	q := dsn.Query()
	q.Add("sslmode", config.SslMode)

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping %w", err)
	}

	return db, nil
}
