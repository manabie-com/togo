package routes

import (
	"lntvan166/togo/internal/controller"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleTask(route *mux.Router) {
	taskRouter := route.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.Authorization)
	taskRouter.HandleFunc("/{id}", controller.HandlerInstance.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("", controller.HandlerInstance.GetAllTaskOfUser).Methods("GET")
	taskRouter.HandleFunc("", controller.HandlerInstance.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", controller.HandlerInstance.CompleteTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", controller.HandlerInstance.DeleteTask).Methods("DELETE")
}
