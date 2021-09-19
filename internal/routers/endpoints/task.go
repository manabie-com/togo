package endpoints

import (
	"context"
	"time"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/services"
)

type TaskEndpoint struct {
	ListTasks Endpoint
	AddTask   Endpoint
}

func MakeListTasksEndpoint(s *services.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		createDate, _ := request.(time.Time)
		res, err := s.TaskService.ListTasks(ctx, createDate)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func MakeAddTaskEndpoint(s *services.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		task, _ := request.(*models.Task)
		res, err := s.TaskService.AddTask(ctx, task)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
