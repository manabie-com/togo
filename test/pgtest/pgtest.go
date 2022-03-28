package pgtest

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
)

func NewContainerDB() (db *sql.DB, dbURL string, cleanup func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("docker.io/postgres", "14.2-alpine3.15", []string{
		"POSTGRES_USER=togo",
		"POSTGRES_PASSWORD=togo",
		"POSTGRES_DB=togo_test",
		"TZ=utc",
	})
	if err != nil {
		log.Fatalf("start postgres: %s", err)
	}
	purge := func() {
		pool.Purge(resource)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		dbPort := resource.GetPort("5432/tcp")
		dbURL = fmt.Sprintf("postgres://togo:togo@localhost:%s/togo_test?sslmode=disable", dbPort)
		db, err = sql.Open("postgres", dbURL)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		purge()
		log.Fatalf("connect to database: %s", err)
	}
	return db, dbURL, purge
}

func ClearDB(db *sql.DB) {
	dbx := sqlx.NewDb(db, "postgres")
	_ = dbx.MustExec(`TRUNCATE "user" CASCADE`)
	_ = dbx.MustExec(`TRUNCATE task CASCADE`)
	_ = dbx.MustExec(`DROP TABLE IF EXISTS schema_migrations`)
}
