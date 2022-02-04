package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeServer() {

	// Create the server
	taskHandler := TaskHandler{}
	router := mux.NewRouter()

	// RESTful API for creating a task. A user is required to make a task.
	router.HandleFunc("/users/{id}/tasks", taskHandler.CreateTask).Methods(http.MethodPost)

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
