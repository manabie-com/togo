package routes

import (
	"lntvan166/togo/controller/user"

	"github.com/gorilla/mux"
)

func HandleUser(route *mux.Router) {
	route.HandleFunc("/user", user.GetAllUsers).Methods("GET")
}
