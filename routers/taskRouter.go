package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)

func taskRouter(r *mux.Router, Repo *controllers.BaseHandler) {
	taskRouting := r.PathPrefix("/tasks").Subrouter()
	taskRouting.Use(middlewares.Logging) // only logging can check task
	taskRouting.HandleFunc("", Repo.ResponseAllTask).Methods("GET")
	taskRouting.HandleFunc("", Repo.CreateTask).Methods("POST")

	taskRoutingid := r.PathPrefix("/tasks/{id}").Subrouter()
	taskRoutingid.Use(middlewares.Logging, middlewares.MiddlewareID)// only logging and ID valid can check task
	taskRoutingid.HandleFunc("", Repo.DeleteTask).Methods("DELETE")
	taskRoutingid.HandleFunc("", Repo.ResponseOneTask).Methods("GET")
	taskRoutingid.HandleFunc("", Repo.UpdateEntireTask).Methods("PUT")
	taskRoutingid.HandleFunc("", Repo.UpdateEntireTask).Methods("PATCH")

}
