package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbconnecter, err := config.GetDBConnecter()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := dbconnecter.Connect()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	log.Println("connect database successfully")

	todoService := services.ToDoService{
		Router: mux.NewRouter(),
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: storages.DBStore{
			DB: db,
		},
	}

	todoService.Router.HandleFunc("/login", todoService.LoginHandler)
	todoService.Router.HandleFunc("/tasks", todoService.Validate(todoService.GetTasksHandler)).Methods(http.MethodGet)
	todoService.Router.HandleFunc("/tasks", todoService.Validate(todoService.CreateTaskHandler)).Methods(http.MethodPost)
	todoService.Router.HandleFunc("/tasks/{id}", todoService.Validate(todoService.DeleteTaskHandler)).Methods(http.MethodDelete)
	http.ListenAndServe(":5050", &todoService)
}
