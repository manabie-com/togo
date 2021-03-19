package utils

type LoginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type AddTaskResponse struct {
	TaskId  string `json:"task_id"`
	Content string `json:"content"`
}
