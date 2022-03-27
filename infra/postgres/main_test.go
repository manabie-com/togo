// +build integration

package postgres_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/laghodessa/togo/infra/postgres"
	"github.com/ory/dockertest/v3"
)

var db *sql.DB
var dbURL string

func TestMain(m *testing.M) {
	os.Exit(setup(m))
}

func setup(m *testing.M) int {
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
	defer pool.Purge(resource)

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
		log.Fatalf("connect to database: %s", err)
	}

	return m.Run()
}

func migrate(t *testing.T) {
	t.Helper()
	if err := postgres.Migrate(dbURL); err != nil {
		t.Errorf("migrate db: %s", err)
	}
}
func clearDB() {
	dbx := sqlx.NewDb(db, "postgres")
	_ = dbx.MustExec(`DELETE FROM "user"`)
	_ = dbx.MustExec(`DELETE FROM task`)
	_ = dbx.MustExec(`DROP TABLE schema_migrations`)
}
