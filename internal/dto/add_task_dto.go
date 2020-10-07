package dto

// AddTaskDTO struct
type AddTaskDTO struct {
	Content string `json:"content" validate:"required"`
}
