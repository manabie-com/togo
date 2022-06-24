package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)

func taskRouter(r *mux.Router) {
	taskRouting := r.PathPrefix("/tasks").Subrouter()
	taskRouting.Use(middlewares.Logging) // only logging can check task
	taskRouting.HandleFunc("", controllers.ResponeAllTask).Methods("GET")
	taskRouting.HandleFunc("", controllers.CreateTask).Methods("POST")

	taskRoutingid := r.PathPrefix("/tasks/{id}").Subrouter()
	taskRoutingid.Use(middlewares.Logging, middlewares.MiddlewareID)// only logging and ID valid can check task
	taskRoutingid.HandleFunc("", controllers.DeleteTask).Methods("DELETE")
	taskRoutingid.HandleFunc("", controllers.ResponeOneTask).Methods("GET")
	taskRoutingid.HandleFunc("", controllers.UpdateEntireTask).Methods("PUT")
	taskRoutingid.HandleFunc("", controllers.UpdateEntireTask).Methods("PATCH")

}
