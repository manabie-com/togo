// +build integration

package services

import (
	"testing"
	"context"

	"github.com/manabie-com/togo/internal/storages/postgres"

	"github.com/jackc/pgx/v4"
)

const postgresURL = "postgres://test:test@localhost:5432/test"

// TestLoginOKPostgres call testLoginOK with PostgresDB
func TestLoginOKPostgres(t *testing.T) {
	db, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(context.TODO())

	pg := &postgres.PostgresDB{
		DB: db,
	}

	var (
		user = "firstUser"
		pass = "example"
	)

	testLoginOK(t, pg, user, pass)
}
