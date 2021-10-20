package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//receiver of app pointer to the application
func (app *ToDoService) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	return router
}
