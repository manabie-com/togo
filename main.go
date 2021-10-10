package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

const (
	host = "localhost"
	port = ":5050"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	log.Printf("Listening and serving in %s%s", host, port)
	http.ListenAndServe(port, &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})
}
