package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"togo/config"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalln("can not load config:", err)
	}

	databaseSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", conf.DbHost, conf.DbUsername, conf.DbPassword, conf.DbName, conf.SslMode)
	testDB, err := sql.Open("postgres", databaseSourceName)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
