package main

import (
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
)

func main() {
	log.Println("Start to-do service...")
	db := &postgres.DataBase{}
	err := db.Init("todo", "postgres", "postgres")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Finalize()

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: db,
	})
}
