package routes

import (
	"lntvan166/togo/internal/controller"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleUser(route *mux.Router) {
	userRouter := route.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.AdminAuthorization)
	userRouter.HandleFunc("", controller.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", controller.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id}", controller.DeleteUserByID).Methods("DELETE")
}
