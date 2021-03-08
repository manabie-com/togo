package schema

type LoginResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type AddTaskResponse struct {
	TaskId string `json:"task_id"`
}
type RegisterResponse struct {
}

type CreateTaskByOwnerResponse struct {
}

type DeleteTaskByOwnerResponse struct {
}

type GetOwnerTaskResponse struct {
}
