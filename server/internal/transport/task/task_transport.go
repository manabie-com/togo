package task

import (
	"context"
	"encoding/json"
	taskHandler "github.com/HoangVyDuong/togo/internal/handler/task"
	"github.com/HoangVyDuong/togo/internal/transport/middleware"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/HoangVyDuong/togo/pkg/kit"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func MakeHandler(router *httprouter.Router, taskHandler taskHandler.Handler) {
	router.Handler("GET", "/tasks", kit.WithCORS(kit.NewServer(
		middleware.Authenticate(
			Endpoint(taskHandler).GetTasks),
		decodeGetTaskRequest,
	)))

	router.Handler("CREATE", "/tasks", kit.WithCORS(kit.NewServer(
		middleware.Authenticate(
			middleware.LimitCreateTask(
				Endpoint(taskHandler).CreateTask)),
		decodeCreateTaskRequest,
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
