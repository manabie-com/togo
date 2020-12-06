package task

import (
	"context"
	taskHandler "github.com/HoangVyDuong/togo/internal/handler/task"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
)

type EndpointSample interface {
	GetTasks(ctx context.Context, request interface{}) (response interface{}, err error)
	CreateTask(ctx context.Context, request interface{}) (response interface{}, err error)
}

type taskEndpoint struct {
	taskHandler taskHandler.Handler
}

func Endpoint(taskHandler taskHandler.Handler) EndpointSample{
	return &taskEndpoint{taskHandler}
}

func (te *taskEndpoint) GetTasks(ctx context.Context, request interface{}) (response interface{}, err error) {
	req := request.(dtos.EmptyRequest)
	return te.taskHandler.GetTasks(ctx, req)
}

func (te *taskEndpoint) CreateTask(ctx context.Context, request interface{}) (response interface{}, err error) {
	req := request.(taskDTO.CreateTaskRequest)
	return te.taskHandler.CreateTask(ctx, req)
}