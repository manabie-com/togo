package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/services"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Get global config
	conf := config.GetConfig()
	conf.Print()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.PGHost, conf.PGPort, conf.PGUser, conf.PGPassword, conf.PGDBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":"+conf.ServerPort, &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PostgresDB{
			DB: db,
		},
	})
}
