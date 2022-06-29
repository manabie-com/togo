package routes

import (
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleTask(route *mux.Router) {
	taskRouter := route.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.Authorization)
	taskRouter.HandleFunc("/{id}", delivery.HandlerInstance.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("", delivery.HandlerInstance.GetAllTaskOfUser).Methods("GET")
	taskRouter.HandleFunc("", delivery.HandlerInstance.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", delivery.HandlerInstance.CompleteTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", delivery.HandlerInstance.DeleteTask).Methods("DELETE")
}
