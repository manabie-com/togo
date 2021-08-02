package dto

type CreateTaskDTO struct {
	Content string `form:"content" validate:"required,min=6,max=32"`
}
