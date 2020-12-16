package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "togo"
	password = "togo"
	dbname   = "datatogo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	//db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB:        db,
			DriveName: "postgres",
		},
	})
}
