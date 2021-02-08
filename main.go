package main

import (
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/services"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	service, err := services.NewToDoServices(config.Jwt, config.DBType.Postgres, config.GetPostgresDBConfig().ToString())
	if err != nil {
		log.Fatal("error opening db", err)
	}
	http.ListenAndServe(":5050", service)

}
