package routes

import (
	"lntvan166/togo/internal/controller"
	"lntvan166/togo/internal/middleware"

	"github.com/gorilla/mux"
)

func HandlePlan(route *mux.Router) {
	planRouter := route.PathPrefix("/plan").Subrouter()

	planRouter.Use(middleware.Authorization)
	planRouter.HandleFunc("", controller.GetPlan).Methods("GET")
	planRouter.HandleFunc("/upgrade/{id}", controller.UpgradePlan).Methods("POST")
}
