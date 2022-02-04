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
	conn, err := sql.Open("postgres", "user=postgres password=postgres dbname=todo_app sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	testQueries = sqlc.New(conn)
	os.Exit(m.Run())
}
