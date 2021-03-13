package main

import (
	"database/sql"
	server2 "github.com/banhquocdanh/togo/internal/server"
	"github.com/banhquocdanh/togo/internal/services"
	sqllite "github.com/banhquocdanh/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	server := server2.NewToDoHttpServer("wqGyEBBfPK9w3Lxw",
		services.NewToDoService(services.WithStore(&sqllite.LiteDB{DB: db})))

	if err := server.Listen(5050); err != nil {
		panic(err)
	}

}
