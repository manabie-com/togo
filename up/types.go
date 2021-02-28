package up

type RegisterRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	MaxTodo  int    `json:"max_todo"`
}

type RegisterResponse struct {
	ID      string `json:"id"`
	MaxTodo int    `json:"max_todo"`
}

type LoginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type LoginResponse string

type ListTasksRequest struct {
	CreatedDate string `json:"created_date"`
}

type ListTasksResponse []*Task

type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

type AddTaskRequest struct {
	Content string `json:"content"`
}
