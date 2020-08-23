package entities

// Task reflects tasks in DB
type Task struct {
	BaseEntity
	Content string `json:"content"`
	UserID  string `json:"user_id"`
}

// User reflects users data from DB
type User struct {
	BaseEntity
	Password string
}

type BaseEntity struct {
	ID          string `json:"id"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}
