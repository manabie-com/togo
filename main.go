package main

import (
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
)

func main() {
	db, err := storages.Initialize("postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.DB.Close()

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  db,
	})
}
