package routes

import (
	"lntvan166/togo/internal/controller"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleUser(route *mux.Router) {
	userRouter := route.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.AdminAuthorization)
	userRouter.HandleFunc("", controller.HandlerInstance.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", controller.HandlerInstance.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id}", controller.HandlerInstance.DeleteUserByID).Methods("DELETE")
}
