package utils

type AddTaskRequest struct {
	Content string `validate:"required,min=1,max=200" json:"content"`
}

type LoginRequest struct {
	Username string `validate:"required,min=4,max=100" json:"username"`
	Password string `validate:"required,min=6,max=200" json:"password"`
}
