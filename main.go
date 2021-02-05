package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"togo/internal/services"
	pg "togo/internal/storages/postgres"
	sqllite "togo/internal/storages/sqlite"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	pgHost     = "localhost"
	pgPort     = 5432
	pgUser     = "postgres"
	pgPassword = "changeme"
	pgDbname   = "pg_db"
)

type userAuthKey int8

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgDbname)
	dbPg, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("connect postgres: %s\n", err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
		StorePg: &pg.ProstgresDB{
			DB: dbPg,
		},
	})
}
