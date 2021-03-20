package routes

import (
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/controllers"
	"github.com/manabie-com/togo/db"
	"github.com/manabie-com/togo/middlewares"
	"github.com/manabie-com/togo/repositories"
	"github.com/manabie-com/togo/services"
	"net/http"
)

var userRepo = repositories.NewUserRepository(db.DB)
var userService = services.NewUserService(&userRepo)

var taskRepo = repositories.NewTaskRepository(db.DB)
var taskService = services.NewTaskService(&taskRepo)

var authController = controllers.NewAuthController(&userService)

var taskController = controllers.NewTaskController(&userService, &taskService)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middlewares []middlewares.Middleware
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "API LOGIN",
		Method:      "POST",
		Pattern:     "/api/auth/login",
		HandlerFunc: authController.Login,
		Middlewares: []middlewares.Middleware{
			middlewares.LoggingMiddleware,
		},
	},

	Route{
		Name:        "API GET ALL TASKS",
		Method:      "GET",
		Pattern:     "/api/tasks",
		HandlerFunc: taskController.AA,
		Middlewares: []middlewares.Middleware{
			middlewares.AuthMiddleware,
			middlewares.LoggingMiddleware,
		},
	},

	Route{
		Name:        "API CREATE A TASK",
		Method:      "POST",
		Pattern:     "/api/tasks",
		HandlerFunc: taskController.AA,
		Middlewares: []middlewares.Middleware{
			middlewares.AuthMiddleware,
			middlewares.LoggingMiddleware,
		},
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	router.MethodNotAllowedHandler = http.HandlerFunc(controllers.NotAllowed)

	for _, route := range routes {
		var handler http.Handler = middlewares.Middlewares(route.HandlerFunc, route.Middlewares...)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
