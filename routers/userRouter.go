package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)

func userRouter(r *mux.Router, Repo *controllers.BaseHandler) {
	userRouting := r.PathPrefix("/users").Subrouter()
	userRouting.Use(middlewares.AdminVerified) // middleware admin, only admin can be modified user
	userRouting.HandleFunc("", Repo.ResponseAllUser).Methods("GET")
	userRouting.HandleFunc("", Repo.CreateUser).Methods("POST")
	
	userRoutingGetme := r.PathPrefix("/users/info").Subrouter()
	userRoutingGetme.Use(middlewares.Logging)
	userRoutingGetme.HandleFunc("", Repo.ResponseOneUser).Methods("GET")

	userRoutingid := r.PathPrefix("/users/{id}").Subrouter()
	userRoutingid.Use(middlewares.AdminVerified, middlewares.MiddlewareID)// middleware admin, only admin can be modified user and check ID
	userRoutingid.HandleFunc("", Repo.ResponseOneUser).Methods("GET")
	userRoutingid.HandleFunc("", Repo.DeleteUser).Methods("DELETE")
	userRoutingid.HandleFunc("", Repo.UpdateUser).Methods("PUT")
	userRoutingid.HandleFunc("", Repo.UpdateUser).Methods("PATCH")
}
