package routes

import (
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/controllers"
	"github.com/manabie-com/togo/middlewares"
	"github.com/manabie-com/togo/repositories"
	"github.com/manabie-com/togo/services"
	"gorm.io/gorm"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middlewares []middlewares.Middleware
}

type Routes []Route

func getRoutes(db *gorm.DB) Routes {
	var userRepo = repositories.NewUserRepository(db)
	var userService = services.NewUserService(&userRepo)

	var taskRepo = repositories.NewTaskRepository(db)
	var taskService = services.NewTaskService(&taskRepo)

	var authController = controllers.NewAuthController(&userService)
	var taskController = controllers.NewTaskController(&userService, &taskService)

	return Routes{
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
			HandlerFunc: taskController.GetTasks,
			Middlewares: []middlewares.Middleware{
				middlewares.LoggingMiddleware,
				middlewares.AuthMiddleware,
			},
		},

		Route{
			Name:        "API CREATE A TASK",
			Method:      "POST",
			Pattern:     "/api/tasks",
			HandlerFunc: taskController.AddTask,
			Middlewares: []middlewares.Middleware{
				middlewares.LoggingMiddleware,
				middlewares.AuthMiddleware,
			},
		},
	}
}

func NewRouter(db *gorm.DB) *mux.Router {
	routes := getRoutes(db)

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
