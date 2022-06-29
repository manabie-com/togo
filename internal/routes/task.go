package routes

import (
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleTask(route *mux.Router, handler *delivery.Handler) {
	taskRouter := route.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.Authorization)
	taskRouter.HandleFunc("/{id}", handler.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("", handler.GetAllTaskOfUser).Methods("GET")
	taskRouter.HandleFunc("", handler.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", handler.CompleteTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", handler.DeleteTask).Methods("DELETE")
}
