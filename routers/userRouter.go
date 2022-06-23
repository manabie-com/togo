package routers

import (
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
)
func userRouter(r *mux.Router){
    userRouting := r.PathPrefix("/users").Subrouter()
    userRouting.HandleFunc("/", controllers.ResponeAllUser).Methods("GET")
	userRouting.HandleFunc("/", controllers.CreateUser).Methods("POST")
    
    userRoutingid := r.PathPrefix("/users/{id}").Subrouter()
    userRoutingid.Use(middlewares.MiddlewareID)
    userRoutingid.HandleFunc("/", controllers.ResponeOneUser).Methods("GET")
	userRoutingid.HandleFunc("/", controllers.DeleteUser).Methods("DELETE")
    userRoutingid.HandleFunc("/", controllers.UpdateUser).Methods("PUT")   
}