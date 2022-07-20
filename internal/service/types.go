package service

type CreateTodoTaskRequest struct {
	Name   string `json:"name"`
	UserID uint64 `json:"user_id"`
}

type CreateTodoTaskResponse struct {
	Message string `json:"message"`
}

type GetTodoTaskResponse struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
