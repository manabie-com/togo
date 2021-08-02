package dto

type CreateTaskDTO struct {
	Content string `form:"content" validate:"required,min=6,max=32"`
}

type UpdateTaskDTO struct {
	IsDone bool `json:"is_done" xml:"is_done" form:"is_done"`
}
