package routes

import (
	"lntvan166/togo/controller/auth"
	"lntvan166/togo/middleware"

	"github.com/gorilla/mux"
)

func HandleAuthentication(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", auth.Register).Methods("POST")
	authRouter.HandleFunc("/login", auth.Login).Methods("POST")

	passwordRouter := authRouter.PathPrefix("/password").Subrouter()

	passwordRouter.Use(middleware.Authorization)
	passwordRouter.HandleFunc("", auth.UpdatePassword).Methods("POST")
}
