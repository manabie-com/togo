package entities

type Task struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
