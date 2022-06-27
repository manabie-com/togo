package routes

import (
	"lntvan166/togo/internal/controller"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleTask(route *mux.Router) {
	taskRouter := route.PathPrefix("/task").Subrouter()
	taskRouter.Use(middleware.Authorization)
	taskRouter.HandleFunc("/{id}", controller.GetTaskByID).Methods("GET")
	taskRouter.HandleFunc("", controller.GetAllTaskOfUser).Methods("GET")
	taskRouter.HandleFunc("", controller.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id}", controller.CheckTask).Methods("PUT")
	taskRouter.HandleFunc("/{id}", controller.DeleteTask).Methods("DELETE")
}
