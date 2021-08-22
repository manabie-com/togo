package dto

type TaskDTO struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
type ListTasksRequestDTO struct {
	UserID      string
	CreatedDate string
}

type ListTaskResponseDTO struct {
	Data []TaskDTO `json:"data"`
}

type AddTaskRequestDTO struct {
	Content string `json:"content"`
	UserID  string
}

type AddTaskResponseDTO struct {
	Data TaskDTO `json:"data"`
}
