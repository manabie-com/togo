// +build integration

package postgres_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/laghodessa/togo/infra/postgres"
	"github.com/laghodessa/togo/test/pgtest"
)

var db *sql.DB
var dbURL string

func TestMain(m *testing.M) {
	os.Exit(setup(m))
}

func setup(m *testing.M) int {
	var cleanup func()
	db, dbURL, cleanup = pgtest.NewContainerDB()
	defer cleanup()
	return m.Run()
}

func migrate(t *testing.T) {
	t.Helper()
	if err := postgres.Migrate(dbURL); err != nil {
		t.Errorf("migrate db: %s", err)
	}
}

func clearDB() {
	pgtest.ClearDB(db)
}
