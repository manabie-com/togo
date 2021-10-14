package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries


const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/todo?sslmode=disable"
)

func TestMain(m *testing.M) {
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}

	testQueries = New(connection)
	
	os.Exit(m.Run())
}