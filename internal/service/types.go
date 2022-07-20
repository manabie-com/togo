package service

import "github.com/tonghia/togo/internal/store"

type CreateTodoTaskRequest struct {
	Name string `json:"name"`
}

type CreateTodoTaskResponse struct {
	Message string `json:"message"`
}

type GetTodoTaskResponse struct {
	Message string           `json:"message"`
	Data    []store.TodoTask `json:"data"`
}
