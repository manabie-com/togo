package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// //
	// // sql-lite
	// //
	// db, err := sql.Open("sqlite3", "./data.db")
	// if err != nil {
	// 	log.Fatal("error opening db", err)
	// }
	// log.Println("SQL database opened")
	// store := &sqllite.LiteDB{
	// 	DB: db,
	// }

	//
	// postgres
	//
	// connection string
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "postgres"
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	store := &postgres.PostgresDB{
		DB: db,
	}

	err = store.InitTables()
	if err != nil {
		log.Fatal("error initializing tables", err)
	}

	//
	// Service Setup
	//

	svc := &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  store,
	}

	log.Println("Database Setup & Check...")

	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging db", err)
	}
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
