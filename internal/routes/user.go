package routes

import (
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleUser(route *mux.Router) {
	userRouter := route.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.AdminAuthorization)
	userRouter.HandleFunc("", delivery.HandlerInstance.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", delivery.HandlerInstance.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id}", delivery.HandlerInstance.DeleteUserByID).Methods("DELETE")
}
