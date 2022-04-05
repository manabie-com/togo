package dtos

// CreateTodoRequest struct
type CreateTodoRequest struct {
	Task    string  `json:"task" binding:"required"`
	DueDate *string `json:"due_date"`
}
