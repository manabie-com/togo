package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/manabie-com/togo/internal/configurations"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := configurations.LoadConfig("../../../resources")

	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to Db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
