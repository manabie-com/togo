package dto

type CreateTaskDTO struct {
	UserName    string `json:"user_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	InsDay      string `json:"ins_day"`
}

type CreateUserDTO struct {
	UserName string `json:"user_name"`
	MaxTasks int    `json:"max_tasks"`
}

type GetTaskOfUserDTO struct {
	UserName string `json:"user_name"`
	InsDay   string `json:"ins_day"`
}
