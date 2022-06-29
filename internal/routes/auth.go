package routes

import (
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleAuthentication(router *mux.Router, handler *delivery.Handler) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", handler.Register).Methods("POST")
	authRouter.HandleFunc("/login", handler.Login).Methods("POST")

	passwordRouter := authRouter.PathPrefix("/password").Subrouter()

	passwordRouter.Use(middleware.Authorization)
	// passwordRouter.HandleFunc("", handler.UpdatePassword).Methods("POST")
}
