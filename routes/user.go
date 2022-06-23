package routes

import (
	"lntvan166/togo/controller/user"

	"github.com/gorilla/mux"
)

func HandleUser(route *mux.Router) {
	userRouter := route.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("", user.GetAllUsers).Methods("GET")
}
