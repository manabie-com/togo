package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	log.Println("SQL database opened")
	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging db", err)
	}

	svc := &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}

	log.Println("Database Setup...")
	// just for fun
	log.Println("Database ping succeeded")

	// update the example user password so that it gets hashed in the DB
	err = svc.Store.SetUserPassword(context.Background(), "firstUser", "example")
	if err != nil {
		log.Fatal("error setting firstUser password", err)
	}
	log.Println("firstUser password hashed")

	max, err := svc.Store.MaxTodo(context.Background(), "firstUser")
	if err != nil {
		log.Fatal("error setting firstUser password", err)
	}
	log.Println("firstUser max todos:", max)

	log.Println("Starting Server")
	err = http.ListenAndServe(":5050", svc)
	log.Fatalln("http server error", err)
}
