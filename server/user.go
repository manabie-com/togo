package server

import (
	"github.com/go-chi/chi"
	userhandler "github.com/manabie-com/togo/internal/user/handler"
	"github.com/manabie-com/togo/registry"
)

type UserDomain struct {
	userHandler *userhandler.UserHandler
}

func newUserDomain(r *registry.Registry) *UserDomain {
	return &UserDomain{
		userHandler: userhandler.New(r.RegisterUserService()),
	}
}

// Base path: /api/user
func (s *UserDomain) router() chi.Router {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		//r.Get("/", s.taskHandler.GetTasks)
		//r.Get("/{task_id}", s.taskHandler.GetTask)
		r.Post("/", s.userHandler.CreateUser)
		r.Post("/login", s.userHandler.Login)
		//r.Delete("/{task_id}", s.taskHandler.DeleteTask)
		r.Put("/{user_id}", s.userHandler.UpdateUser)
	})

	return r
}
