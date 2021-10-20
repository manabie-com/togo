package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//receiver of app pointer to the application where routes of the application will be listed
func (app *ToDoService) routes() *httprouter.Router {
	router := httprouter.New()
	//get the status of the application
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	return router
}
