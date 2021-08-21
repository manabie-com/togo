package dto

import "context"

type ListTaskRequest struct {
	CreatedDate string `json:"created_date"`
}

type ListTaskResponse struct {
	Data []Task `json:"data"`
}

func (ltr *ListTaskResponse) ToRes() interface{} {
	return ltr
}

type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

type AddTaskRequest struct {
	Content string `json:"content"`
}

type AddTaskResponse struct {
	Data Task `json:"data"`
}

func (atr *AddTaskResponse) ToRes() interface{} {
	return atr
}

type ITaskApi interface {
	ListTasksByUserAndDate(ctx context.Context, request ListTaskRequest) (*ListTaskResponse, error)
	AddTask(ctx context.Context, request AddTaskRequest) (*AddTaskResponse, error)
}
