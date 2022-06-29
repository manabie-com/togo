package routes

import (
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleAuthentication(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", delivery.HandlerInstance.Register).Methods("POST")
	authRouter.HandleFunc("/login", delivery.HandlerInstance.Login).Methods("POST")

	passwordRouter := authRouter.PathPrefix("/password").Subrouter()

	passwordRouter.Use(middleware.Authorization)
	// passwordRouter.HandleFunc("", delivery.HandlerInstance.UpdatePassword).Methods("POST")
}
