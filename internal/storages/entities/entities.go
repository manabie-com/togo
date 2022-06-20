package entities

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type RateLimit struct {
	Count  uint
	UserID string
	Date   string
}
