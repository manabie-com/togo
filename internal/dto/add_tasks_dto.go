package dto

// AddTasksDTO struct
type AddTasksDTO struct {
	Contents []string `json:"contents" validate:"required"`
}
