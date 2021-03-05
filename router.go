package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/services"
)

// NewRouter setup app router
func NewRouter(ToDoService *services.ToDoService,
	userService *services.UserService) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", userService.Login).Methods("POST")
	r.Handle("/profile", userService.ValidToken(http.HandlerFunc(userService.GetProfile)))

	r.Handle("/tasks", userService.ValidToken(http.HandlerFunc(ToDoService.AddTask))).Methods("POST")
	r.Handle("/tasks", userService.ValidToken(http.HandlerFunc(ToDoService.ListTasks))).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))

	return r
}
