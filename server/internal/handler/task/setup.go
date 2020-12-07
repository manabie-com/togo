package task

import (
	"context"
	taskService "github.com/HoangVyDuong/togo/internal/usecase/task"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
)

type Handler interface {
	GetTasks(ctx context.Context, request dtos.EmptyRequest) (response taskDTO.Tasks, err error)
	CreateTask(ctx context.Context, request taskDTO.CreateTaskRequest) (response taskDTO.CreateTaskResponse, err error)
}

type taskHandler struct {
	taskService taskService.Service
}

func NewHander(taskService taskService.Service) Handler{
	return &taskHandler{taskService}
}
