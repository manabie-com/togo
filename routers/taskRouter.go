package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)

func taskRouter(r *mux.Router, bh *controllers.BaseHandler) {
	taskRouting := r.PathPrefix("/tasks").Subrouter()
	taskRouting.Use(middlewares.Logging) // only logging can check task
	taskRouting.HandleFunc("", bh.ResponseAllTask).Methods("GET")
	taskRouting.HandleFunc("", middlewares.CheckLimitTaskUserMiddleware(bh, bh.CreateTask)).Methods("POST")

	taskRoutingid := r.PathPrefix("/tasks/{id}").Subrouter()
	taskRoutingid.Use(middlewares.Logging, middlewares.MiddlewareID)// only logging and ID valid can check task
	taskRoutingid.HandleFunc("", bh.DeleteFromTask).Methods("DELETE")
	taskRoutingid.HandleFunc("", bh.ResponseOneTask).Methods("GET")
	taskRoutingid.HandleFunc("", bh.UpdateToTask).Methods("PUT")
	taskRoutingid.HandleFunc("", bh.UpdateToTask).Methods("PATCH")

}
