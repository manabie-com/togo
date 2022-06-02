package router

import (
	"net/http"

	"togo/controller"
	"togo/middleware"

	"github.com/go-chi/chi/v5"
)

type routes struct {
	Tasks controller.TaskController
	Users controller.UserController
}

func NewChiRouter(tasks controller.TaskController, users controller.UserController) RouterInterface {
	return &routes{
		Tasks: tasks,
		Users: users,
	}
}

func (route *routes) Router() http.Handler {
	r := chi.NewRouter()

	r.Post("/registeration", route.Users.Register)
	r.Put("/login", route.Users.Login)

	r.Route("/tasks", func(r chi.Router) {
		r.Use(middleware.AuthenticateRequest)
		r.Post("/", route.Tasks.CreateTask)
	})

	return r
}
