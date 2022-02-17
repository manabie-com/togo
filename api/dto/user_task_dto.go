package dto

type CreateTaskDTO struct {
	UserName    string `json:"user_name" validate:"not_empty"`
	Title       string `json:"title" validate:"not_empty"`
	Description string `json:"description" validate:"not_empty"`
	InsDay      string `json:"ins_day" validate:"not_empty,yyyymmdd_date"`
}

type CreateUserDTO struct {
	UserName string `json:"user_name" validate:"not_empty"`
	MaxTasks int    `json:"max_tasks" validate:"required"`
}

type GetTaskOfUserDTO struct {
	UserName string `json:"user_name" validate:"not_empty"`
	InsDay   string `json:"ins_day"`
}
