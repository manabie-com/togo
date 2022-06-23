package routes

import (
	"lntvan166/togo/controller/task"
	"lntvan166/togo/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleTask(route *mux.Router) {
	// route.HandleFunc("/task", task.GetAllTasks).Methods("GET")
	// route.HandleFunc("/task/{id}", task.GetTask).Methods("GET")
	route.Handle("/task", middleware.Authorization(http.HandlerFunc(task.GetTaskForUser))).Methods("GET")
	route.Handle("/task", middleware.Authorization(http.HandlerFunc(task.CreateTask))).Methods("POST")
	// route.HandleFunc("/task/{id}", task.UpdateTask).Methods("PUT")
	// route.HandleFunc("/task/{id}", task.DeleteTask).Methods("DELETE")
}
