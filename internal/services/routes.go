package services

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//receiver of app pointer to the application where routes of the application will be listed
func (app *ToDoService) Routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/login", app.getAuthToken)
	router.HandlerFunc(http.MethodGet, "/tasks", app.listTasks)
	//get the status of the application
	return app.enableCORS(router)
}
