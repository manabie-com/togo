package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"net/http"
	"os"
)

var started = false

func main() {
	println("Start TODO service")

	//db, err := sql.Open("sqlite3", "./data.db")
	//if err != nil {
	//	log.Fatal("error opening db", err)
	//}
	//
	//http.ListenAndServe(":5050", &services.ToDoService{
	//	JWTKey: "wqGyEBBfPK9w3Lxw",
	//	Store: &sqllite.LiteDB{
	//		DB: db,
	//	},
	//})

	/* --- Change to using Postgres --- */
	connStr := config.GetConfig().GetConnString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: os.Getenv("ENCRYPTION"),
		Store:  &postgres.Postgres{DB: db},
	})
}
