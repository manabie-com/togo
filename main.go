package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

const (
	host = "localhost" // when run go run main.go on local machine
	// host = "database" // when run in container
	// PostgreSQL container is running on local machine, so we need to connect with localhost. If it is running on a specific server, use your server IP.
	port     = 5432
	user     = "togoapp"
	password = "togoapp"
	dbname   = "togodb"
)

func main() {
	// db, err := sql.Open("sqlite3", "./data.db")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	log.Println("connect database successfully")

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PqDB{
			DB: db,
		},
	})
}
