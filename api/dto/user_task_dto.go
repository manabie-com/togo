package dto

// CreateTaskDTO is the data used for creating a user task
type CreateTaskDTO struct {
	UserName    string `json:"user_name" validate:"not_empty"`
	Title       string `json:"title" validate:"not_empty"`
	Description string `json:"description" validate:"not_empty"`
	InsDay      string `json:"ins_day" validate:"not_empty,yyyymmdd_date"`
}

// CreateUserDTO is the data used for creating a user
type CreateUserDTO struct {
	UserName string `json:"user_name" validate:"not_empty"`
	MaxTasks int    `json:"max_tasks" validate:"required"`
}

// GetTaskOfUserDTO is the data used for fetching the tasks of a user
type GetTaskOfUserDTO struct {
	UserName string `json:"user_name" validate:"not_empty"`
	InsDay   string `json:"ins_day"`
}
