package dto

import "time"

type CreateTaskDTO struct {
	Content string `form:"content" validate:"required,min=6,max=32"`
}

type UpdateTaskDTO struct {
	IsDone bool `json:"is_done" xml:"is_done" form:"is_done"`
}

type SearchTasksRequest struct {
	IsDone      bool      `query:"is_done"`
	CreatedDate time.Time `query:"created_date"`
}
