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
		Name:        "Docs",
		Method:      "GET",
		Pattern:     "/v2/docs",
		HandlerFunc: controllers.AA,
	},

	Route{
		Name:        "GetUserByName",
		Method:      "GET",
		Pattern:     "/v2/user/{username}",
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
