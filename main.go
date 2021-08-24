package main

import (
	"github.com/manabie-com/togo/internal/tools"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/api"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("You should run application with config file")
	}
	config, err := tools.LoadConfig(args[1])
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
}
