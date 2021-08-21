package main

import (
	"github.com/manabie-com/togo/internal/storages/postgreSQL"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
)

func main() {
	db := postgreSQL.GetConnection()
	log.Println("Server is start port: 5000")

	http.ListenAndServe(":5000", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: db,
	})
}
