package routes

import (
	"lntvan166/togo/controller/task"
	"lntvan166/togo/middleware"

	"github.com/gorilla/mux"
)

func HandleTask(route *mux.Router) {
	taskRouter := route.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.Authorization)
	// route.HandleFunc("/task", task.GetAllTasks).Methods("GET")
	// route.HandleFunc("/task/{id}", task.GetTask).Methods("GET")
	taskRouter.HandleFunc("", task.GetTaskForUser).Methods("GET")
	taskRouter.HandleFunc("", task.CreateTask).Methods("POST")
	// route.HandleFunc("/task/{id}", task.UpdateTask).Methods("PUT")
	// route.HandleFunc("/task/{id}", task.DeleteTask).Methods("DELETE")
}
