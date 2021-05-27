package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	host     = "db"
	port     = 5432
	user     = "todo"
	password = "todo"
	dbname   = "todo"
)

func main() {
	dbType := os.Getenv("DB_TYPE")
	var dbName, dbInfo string
	if dbType == "postgres" {
		dbName = "postgres"
		dbInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	} else {
		dbName = "sqlite3"
		dbInfo = "./data.db"
	}
	db, err := sql.Open(dbName, dbInfo)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})
}
