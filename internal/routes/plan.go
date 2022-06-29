package routes

import (
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandlePlan(route *mux.Router) {
	planRouter := route.PathPrefix("/plan").Subrouter()

	planRouter.Use(middleware.Authorization)
	planRouter.HandleFunc("", delivery.HandlerInstance.GetPlan).Methods("GET")
	planRouter.HandleFunc("/upgrade/{id}", delivery.HandlerInstance.UpgradePlan).Methods("POST")
}
