// +build integration

package services

import (
	"context"
	"testing"

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

// TestLoginUnauthorizedPostgres calls testLoginUnauthorized with PostgresDB
func TestLoginUnauthorizedPostgres(t *testing.T) {
	db, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(context.TODO())

	pg := &postgres.PostgresDB{
		DB: db,
	}

	testLoginUnauthorized(t, pg)
}

// TestListTasksInvalidTokenPostgres calls testListTasksInvalidToken with a mock DB
func TestListTasksInvalidTokenPostgres(t *testing.T) {
	db, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(context.TODO())

	pg := &postgres.PostgresDB{
		DB: db,
	}

	testListTasksInvalidToken(t, pg)
}

// TestListTasksOKPostgres calls testListTasksOK with a mock DB
func TestListTasksOKPostgres(t *testing.T) {
	var (
		user = "firstUser"
		pass = "example"
		date = "2006-01-02"
	)

	db, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(context.TODO())

	pg := &postgres.PostgresDB{
		DB: db,
	}

	testListTasksOK(t, pg, user, pass, date)
}

// TestAddTasksInvalidTokenPostgres calls testAddTasksInvalidToken with PostgreSQL
func TestAddTasksInvalidTokenPostgres(t *testing.T) {
	db, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(context.TODO())

	pg := &postgres.PostgresDB{
		DB: db,
	}

	testAddTasksInvalidToken(t, pg)
}
