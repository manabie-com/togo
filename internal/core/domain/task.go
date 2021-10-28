package domain

// Task reflects tasks in DB
type Task struct {
	Id          string `json:"id"`
	Content     string `json:"content"`
	UserId      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
