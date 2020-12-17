package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	host = "host.docker.internal"
	//host     = "localhost"
	port     = 5432
	user     = "togo"
	password = "togo"
	dbname   = "datatogo"
)

func main() {
	var selectDB string

	flag.StringVar(&selectDB, "db", "", "")
	flag.Parse()

	var db *sql.DB
	var err error
	if selectDB == "postgres" {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal("error opening db", err)
		}
	} else {
		selectDB = "sqlite3"
		db, err = sql.Open("sqlite3", "./data.db")
		if err != nil {
			log.Fatal("error opening db", err)
		}
	}
	fmt.Println("App running with db " + selectDB + " on port 5050...")

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB:        db,
			DriveName: selectDB,
		},
	})
}
