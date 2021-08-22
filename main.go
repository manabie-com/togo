package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/api"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "docker"
	password = "docker"
	dbname   = "todo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	todoApi := api.NewToDoApi("wqGyEBBfPK9w3Lxw", db)
	err = http.ListenAndServe(":5050", &todoApi)
	if err != nil {
		log.Fatal("error listen and serve api", err)
	}
}
