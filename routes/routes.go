package routes

import (
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/controllers"
	"github.com/manabie-com/togo/middlewares"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secure      bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "API LOGIN",
		Method:      "GET",
		Pattern:     "/auth/login",
		HandlerFunc: controllers.AA,
	},

	Route{
		Name:        "API GET ALL TASKS",
		Method:      "GET",
		Pattern:     "/tasks",
		HandlerFunc: controllers.AA,
		Secure:      true,
	},

	Route{
		Name:        "API CREATE A TASK",
		Method:      "POST",
		Pattern:     "/tasks",
		HandlerFunc: controllers.AA,
		Secure:      true,
	},
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	router.MethodNotAllowedHandler = http.HandlerFunc(controllers.NotAllowed)

	for _, route := range routes {
		var handler http.Handler

		if route.Secure {
			handler = middlewares.AuthMiddleware(middlewares.LoggingMiddleware(route.HandlerFunc))
		} else {
			handler = middlewares.LoggingMiddleware(route.HandlerFunc)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var Router = middlewares.Middleware(newRouter())
