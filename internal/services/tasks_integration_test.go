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

	testLoginOK(t, pg, testUser, testPass)
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

// TestListTasksInvalidTokenPostgres calls testListTasksInvalidToken with PostgreSQL
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

// TestListTasksOKPostgres calls testListTasksOK with PostgreSQL
func TestListTasksOKPostgres(t *testing.T) {
	db, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(context.TODO())

	pg := &postgres.PostgresDB{
		DB: db,
	}

	testListTasksOK(t, pg, testUser, testPass, testDate)
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

// TestAddTasksOKPostgres calls testAddTasksOK with PostgreSQL
func TestAddTasksOKPostgres(t *testing.T) {
	db, err := pgx.Connect(context.TODO(), postgresURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close(context.TODO())

	pg := &postgres.PostgresDB{
		DB: db,
	}

	testAddTasksOK(t, pg, testUser, testPass)
}
