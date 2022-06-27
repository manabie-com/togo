package routes

import (
	"lntvan166/togo/internal/controller"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleAuthentication(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", controller.Register).Methods("POST")
	authRouter.HandleFunc("/login", controller.Login).Methods("POST")

	passwordRouter := authRouter.PathPrefix("/password").Subrouter()

	passwordRouter.Use(middleware.Authorization)
	passwordRouter.HandleFunc("", controller.UpdatePassword).Methods("POST")
}
