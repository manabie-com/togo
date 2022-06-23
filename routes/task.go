package routes

import (
	"lntvan166/togo/controller/task"
	"lntvan166/togo/middleware"

	"github.com/gorilla/mux"
)

func HandleTask(route *mux.Router) {
	taskRouter := route.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.Authorization)
	taskRouter.HandleFunc("/{id}", task.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("", task.GetAllTaskOfUser).Methods("GET")
	taskRouter.HandleFunc("", task.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", task.CheckTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", task.DeleteTask).Methods("DELETE")
}
