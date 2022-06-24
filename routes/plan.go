package routes

import (
	"lntvan166/togo/controller/plan"
	"lntvan166/togo/middleware"

	"github.com/gorilla/mux"
)

func HandlePlan(route *mux.Router) {
	planRouter := route.PathPrefix("/plan").Subrouter()

	planRouter.Use(middleware.Authorization)
	planRouter.HandleFunc("", plan.GetPlan).Methods("GET")
	planRouter.HandleFunc("/upgrade/{id}", plan.UpgradePlan).Methods("POST")
}
