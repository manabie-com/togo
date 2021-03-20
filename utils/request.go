package utils

type AddTaskRequest struct {
	Content string `json:"content" validate:"required"`
}

type LoginRequest struct {
	Username string `validate:"required,min=4,max=100" json:"username"`
	Password string `validate:"required,min=6,max=200" json:"password"`
}
