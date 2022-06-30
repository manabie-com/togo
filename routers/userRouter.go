package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)

func userRouter(r *mux.Router, bh *controllers.BaseHandler) {
	UserRouting := r.PathPrefix("/users").Subrouter()
	UserRouting.Use(middlewares.AdminVerified) // middleware admin, only admin can be modified user
	UserRouting.HandleFunc("", bh.ResponseAllUser).Methods("GET")
	UserRouting.HandleFunc("", middlewares.ValidUsernameAndHashPassword(bh, bh.CreateUser)).Methods("POST")
	
	userRoutingGetme := r.PathPrefix("/users/info").Subrouter()
	userRoutingGetme.Use(middlewares.LoggingVerified)
	userRoutingGetme.HandleFunc("", bh.ResponseOneUser).Methods("GET")

	userRoutingid := r.PathPrefix("/users/{id}").Subrouter()
	userRoutingid.Use(middlewares.AdminVerified, middlewares.MiddlewareID)// middleware admin, only admin can be modified user and check ID
	userRoutingid.HandleFunc("", bh.ResponseOneUser).Methods("GET")
	userRoutingid.HandleFunc("", bh.DeleteFromUser).Methods("DELETE")
	userRoutingid.HandleFunc("", middlewares.ValidUsernameAndHashPassword(bh, bh.UpdateToUser)).Methods("PUT")
	userRoutingid.HandleFunc("", middlewares.ValidUsernameAndHashPassword(bh, bh.UpdateToUser)).Methods("PATCH")
}
