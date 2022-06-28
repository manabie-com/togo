package entities

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Completed   bool   `json:"completed"`
	UserID      int    `json:"user_id"`
}
