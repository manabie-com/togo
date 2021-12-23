package entity

type Task struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Content     string `json:"content"`
	Status      string `json:"status"`
	CreatedDate string `json:"created_date"`
}
