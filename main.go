package main

import (
	"database/sql"
	"fmt"
	"github.com/manabie-com/togo/internal/http"
	"github.com/manabie-com/togo/internal/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	fmt.Println("Setting up database...")
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	fmt.Println("Database setup done.")

	userRepo := &sqlite.UserRepo{DB:db}
	taskRepo := &sqlite.TaskRepo{DB:db}
	fmt.Println("Listening on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", http.JSONServer("wqGyEBBfPK9w3Lxw", userRepo, taskRepo)))
}
