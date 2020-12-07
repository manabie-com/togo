package task

import (
	"context"
	"encoding/json"
	taskHandler "github.com/HoangVyDuong/togo/internal/handler/task"
	"github.com/HoangVyDuong/togo/internal/transport/middleware"
	taskService "github.com/HoangVyDuong/togo/internal/usecase/task"
	userService "github.com/HoangVyDuong/togo/internal/usecase/user"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/HoangVyDuong/togo/pkg/kit"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHandler(router *httprouter.Router, taskHandler taskHandler.Handler, userService userService.Service, taskService taskService.Service) {
	convertAuthorization := func(ctx context.Context, req *http.Request) context.Context {
		return context.WithValue(ctx, define.ContextKeyAuthorization, req.Header.Get("Authorization"))
	}
	router.Handler("GET", "/tasks", kit.WithCORS(kit.NewServer(
		middleware.Authenticate(
			Endpoint(taskHandler).GetTasks),
		decodeGetTaskRequest,
		kit.ServerBefore(convertAuthorization),
	)))

	router.Handler("POST", "/tasks", kit.WithCORS(kit.NewServer(
		middleware.Authenticate(
			middleware.LimitCreateTask(taskService, userService)(
				Endpoint(taskHandler).CreateTask)),
		decodeCreateTaskRequest,
		kit.ServerBefore(convertAuthorization),
	)))

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
