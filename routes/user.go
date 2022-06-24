package routes

import (
	"lntvan166/togo/controller/user"
	"lntvan166/togo/middleware"

	"github.com/gorilla/mux"
)

func HandleUser(route *mux.Router) {
	userRouter := route.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.AdminAuthorization)
	userRouter.HandleFunc("", user.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", user.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id}", user.DeleteUserByID).Methods("DELETE")
}
