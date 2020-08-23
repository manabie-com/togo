package entities

// Task reflects tasks in DB
type Task struct {
	baseEntity
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}

// User reflects users data from DB
type User struct {
	baseEntity
	Username string
	Password string
}

type baseEntity struct {
	ID          string `json:"id"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}
