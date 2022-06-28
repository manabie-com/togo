package routes

import (
	"lntvan166/togo/internal/controller"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandleAuthentication(router *mux.Router) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/register", controller.HandlerInstance.Register).Methods("POST")
	authRouter.HandleFunc("/login", controller.HandlerInstance.Login).Methods("POST")

	passwordRouter := authRouter.PathPrefix("/password").Subrouter()

	passwordRouter.Use(middleware.Authorization)
	// passwordRouter.HandleFunc("", controller.HandlerInstance.UpdatePassword).Methods("POST")
}
