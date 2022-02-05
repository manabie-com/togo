package api

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	sqlc "github.com/roandayne/togo/db/sqlc"
)

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	database_url := "postgres://postgres:postgres@db:5432/todo_app?sslmode=disable"
	conn, err := sql.Open("postgres", database_url)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	testQueries = sqlc.New(conn)
	os.Exit(m.Run())
}
