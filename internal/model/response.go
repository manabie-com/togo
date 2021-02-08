package model

import "github.com/manabie-com/togo/internal/storages"

type LoginSuccessResponse struct {
	Data *string `json:"data"`
}

type ErrorResponse struct {
	Error *string `json:"error"`
}

type GetTaskResponse struct {
	Data []*storages.Task `json:"data"`
}

type AddTaskResponse struct {
	Data *storages.Task `json:"data"`
}
