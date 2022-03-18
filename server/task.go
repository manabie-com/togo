package server

import (
	"github.com/go-chi/chi"
	"github.com/manabie-com/togo/internal/task/handler"
	userhandler "github.com/manabie-com/togo/internal/user/handler"
	"github.com/manabie-com/togo/registry"
)

type TaskDomain struct {
	taskHandler    *handler.TaskHandler
	userHandler    *userhandler.UserHandler
	userMiddleware *userhandler.UserMiddleware
}

func newTaskDomain(r *registry.Registry) *TaskDomain {
	return &TaskDomain{
		taskHandler:    handler.New(r.RegisterTaskService()),
		userMiddleware: userhandler.NewUserMiddleware(r.RegisterUserService()),
	}
}

// Base path: /api/task
func (s *TaskDomain) router() chi.Router {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(s.userMiddleware.UserOnly)
		r.Get("/", s.taskHandler.GetTasks)
		r.Get("/{task_id}", s.taskHandler.GetTask)
		r.Post("/", s.taskHandler.CreateTask)
		r.Delete("/{task_id}", s.taskHandler.DeleteTask)
		r.Put("/{task_id}", s.taskHandler.UpdateTask)
	})

	return r
}
