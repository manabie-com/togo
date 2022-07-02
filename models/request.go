package models

// CreateUserRequest model for create user request
type CreateUserRequest struct {
	Username       string `json:"username"`
	TaskDailyLimit int32  `json:"taskDailyLimit"`
}

// UpdateUserRequest model for update user request
type UpdateUserRequest struct {
	Username       string `json:"username"`
	TaskDailyLimit int32  `json:"taskDailyLimit"`
}

// CreateTaskRequest model for create task request
type CreateTaskRequest struct {
	Username    string  `json:"username"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
}

// DeleteUserRequest model for delete user request
type DeleteUserRequest struct {
	Username string `json:"username"`
}
