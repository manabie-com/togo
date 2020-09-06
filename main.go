package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
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
		Store: &postgres.DBStore{
			DB: db,
		},
	}
	// login
	todoService.Router.HandleFunc("/login", todoService.LoginHandler)
	// get tasks by date
	todoService.Router.HandleFunc("/tasks", todoService.Validate(todoService.GetTasksHandler)).Methods(http.MethodGet)
	// create a task
	todoService.Router.HandleFunc("/tasks", todoService.Validate(todoService.CreateTaskHandler)).Methods(http.MethodPost)
	// delete tasks by date
	todoService.Router.HandleFunc("/tasks", todoService.Validate(todoService.DeleteTasksHandler)).Methods(http.MethodDelete)
	// delete task by id
	todoService.Router.HandleFunc("/tasks/{id}", todoService.Validate(todoService.DeleteTaskHandler)).Methods(http.MethodDelete)
	// update task status by id
	todoService.Router.HandleFunc("/tasks/{id}", todoService.Validate(todoService.UpdateTaskStatusHandler)).Methods(http.MethodPatch)
	// update all task status by id
	todoService.Router.HandleFunc("/tasks", todoService.Validate(todoService.UpdateAllTaskStatusHandler)).Methods(http.MethodPatch)
	http.ListenAndServe(":5050", &todoService)
}
