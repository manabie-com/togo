package entities

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content" validate:"required"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
