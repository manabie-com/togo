package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

func main() {
	db, err := sql.Open("pgx", sqllite.DBConnectionURL())
	if err != nil {
		log.Fatal("error opening db: ", err)
	}
	defer db.Close()

	log.Println("Starting server on port 5050")
	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})
}
