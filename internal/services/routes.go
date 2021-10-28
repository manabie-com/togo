package services

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//receiver of app pointer to the application where routes of the application will be listed
func (app *ToDoService) Routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/login", app.getAuthToken)
	router.HandlerFunc(http.MethodGet, "/tasks", app.listTasks)
	router.HandlerFunc(http.MethodPost, "/tasks", app.addTask)
	router.HandlerFunc(http.MethodPost, "/task/delete", app.deleteTask)
	router.HandlerFunc(http.MethodPost, "/task/update", app.updateTask)

	//use the middlerware for header setters
	return app.enableCORS(router)
}
