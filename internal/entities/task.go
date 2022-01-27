package entities

type Task struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Tasks struct {
	Tasks []*Task `json:"tasks"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
}

type TaskFilter struct {
	UserID int `json:"userId"`
	Page   int `json:"page"`
	Limit  int `json:"limit"`
}
