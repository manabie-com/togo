package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/configurations"
	"github.com/manabie-com/togo/internal/services"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"

	// sqllite "github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/mattn/go-sqlite3"

	_ "github.com/lib/pq"
)

func main() {
	config, err := configurations.LoadConfig("./resources")

	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	http.ListenAndServe(":5050", &services.ServiceController{Config: config, Store: postgres.NewStore(conn)})
}
