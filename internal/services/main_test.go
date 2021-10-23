package services

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/storages"
)

var testQueries storages.Models

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:root@localhost/todos?sslmode=disable"
)

/**
* Test for connecting database postgres
**/
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = storages.NewModels(conn)

	os.Exit(m.Run())
}
