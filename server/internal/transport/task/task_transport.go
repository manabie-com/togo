package task

import (
	"context"
	"encoding/json"
	taskHandler "github.com/HoangVyDuong/togo/internal/handler/task"
	"github.com/HoangVyDuong/togo/internal/kit"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Endpoint interface {
	GetTasks(ctx context.Context, request interface{}) (response interface{}, err error)
	CreateTask(ctx context.Context, request interface{}) (response interface{}, err error)
}

type taskEndpoint struct {
	taskHandler taskHandler.Handler
}

func WithEndpoint(taskHandler taskHandler.Handler) Endpoint{
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

func WithHandler(router *httprouter.Router, taskHandler taskHandler.Handler) {
	router.Handler("GET", "/tasks", kit.NewServer(
		WithEndpoint(taskHandler).GetTasks,
		decodeGetTaskRequest,
	))

	router.Handler("CREATE", "/tasks", kit.NewServer(
		WithEndpoint(taskHandler).CreateTask,
		decodeCreateTaskRequest,
	))

}

func decodeGetTaskRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return dtos.EmptyRequest{}, nil
}

func decodeCreateTaskRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req taskDTO.CreateTaskRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}
