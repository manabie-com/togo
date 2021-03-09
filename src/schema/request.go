package schema

type LoginRequest struct {
	UserId   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AddTaskRequest struct {
	Content string `json:"content" validate:"required"`
}

type RegisterRequest struct {
}

type CreateTaskByOwnerRequest struct {
}

type DeleteTaskByOwnerRequest struct {
}

type GetOwnerTaskRequest struct {
}
