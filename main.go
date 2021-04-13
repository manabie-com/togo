package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"

	_ "github.com/lib/pq"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "my_postgres"
	password = "my_password"
	dbname   = "manabie"
)

func main() {
	pqConnectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", pqConnectionString)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PostgresDB{
			DB: db,
		},
	})
}
