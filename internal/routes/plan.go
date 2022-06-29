package routes

import (
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandlePlan(route *mux.Router, handler *delivery.Handler) {
	planRouter := route.PathPrefix("/plan").Subrouter()

	planRouter.Use(middleware.Authorization)
	planRouter.HandleFunc("", handler.GetPlan).Methods("GET")
	planRouter.HandleFunc("/upgrade/{id}", handler.UpgradePlan).Methods("POST")
}
