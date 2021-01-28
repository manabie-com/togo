package postgres

import (
	"github.com/manabie-com/togo/config"

	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries PostgresDB

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSourceTest)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = PostgresDB{
		DB: testDB,
	}

	os.Exit(m.Run())
}
