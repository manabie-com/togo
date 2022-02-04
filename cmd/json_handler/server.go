package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeServer() {

	// Create the server
	taskHandler := TaskHandler{}
	router := mux.NewRouter()

	// RESTful API for creating a task. A user is required to make a task.
	router.HandleFunc("/user/{id}/tasks", taskHandler.CreateTask).Methods(http.MethodPost)
}
