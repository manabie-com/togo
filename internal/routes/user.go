package routes

import (
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleUser(route *mux.Router, handler *delivery.Handler) {
	userRouter := route.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.Authorization)
	userRouter.HandleFunc("", handler.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("/{id}", handler.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id}", handler.DeleteUserByID).Methods("DELETE")
}
