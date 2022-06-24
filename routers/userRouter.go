package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)

func userRouter(r *mux.Router) {
	userRouting := r.PathPrefix("/users").Subrouter()
	userRouting.Use(middlewares.AdminVerified) // middleware admin, only admin can be modified user
	userRouting.HandleFunc("", controllers.ResponeAllUser).Methods("GET")
	userRouting.HandleFunc("", controllers.CreateUser).Methods("POST")
	
	userRoutingGetme := r.PathPrefix("/users/info").Subrouter()
	userRoutingGetme.Use(middlewares.Logging)
	userRoutingGetme.HandleFunc("", controllers.ResponeOneUser).Methods("GET")

	userRoutingid := r.PathPrefix("/users/{id}").Subrouter()
	userRoutingid.Use(middlewares.AdminVerified, middlewares.MiddlewareID)// middleware admin, only admin can be modified user and check ID
	userRoutingid.HandleFunc("", controllers.ResponeOneUser).Methods("GET")
	userRoutingid.HandleFunc("", controllers.DeleteUser).Methods("DELETE")
	userRoutingid.HandleFunc("", controllers.UpdateUser).Methods("PUT")
	userRoutingid.HandleFunc("", controllers.UpdateUser).Methods("PATCH")
}
