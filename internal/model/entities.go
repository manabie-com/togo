package model

// Task reflects tasks in DB
type Task struct {
	ID           string `json:"id"`
	Content      string `json:"content"`
	UserID       string `json:"user_id"`
	CreatedDate  string `json:"created_date"`
	NumberInDate int    `json:"number_in_date"`
}

// User reflects users data from DB
type User struct {
	ID       string `json:"user_id"`
	Password string `json:"password"`
	MaxTodo  int    `json:"max_todo"`
}
