package intergration_test

import (
	"github.com/jmoiron/sqlx"
	"github.com/manabie-com/togo/internal/api"
	"github.com/manabie-com/togo/internal/tools"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	config, err := tools.LoadConfig("../deploy/test_todo/")
	if err != nil {
		log.Fatal("You should run application with valid file path", err)
	}
	db, err := sqlx.Open("postgres", config.PsqlInfo())
	if err != nil {
		log.Fatal("error opening db", err)
	}
	todoApi := api.NewToDoApi("wqGyEBBfPK9w3Lxw", db)
	err = http.ListenAndServe(config.ServerPort(), &todoApi)
	if err != nil {
		log.Fatal("error listen and serve api", err)
	}
	return m.Run()
}
