package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

func GetDBInfo() string {
	host := os.Getenv("HOST")
	port := 5432
	database := os.Getenv("DATABASE")
	username := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	if "" == host {
		host = "127.0.0.1"
	}
	if "" != os.Getenv("PORT") {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	if "" == database {
		database = "ocsenDB"
	}
	if "" == username {
		username = "ocsen"
	}
	if "" == password {
		password = "ocsen-hoc-code"
	}

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	return dbInfo
}

func main() {
	dbInfo := GetDBInfo()
	db, err := sql.Open("postgres", dbInfo)
	db.SetMaxOpenConns(10)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	serv := &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}
	serv.Store.CreateDatabase()
	http.ListenAndServe(":5050", serv)
}
