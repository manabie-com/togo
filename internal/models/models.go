package models

// User deinfes `users` table in the database
type User struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	MaxTaskPerDay uint   `json:"max_task_per_day"`
}

// Task defines `tasks` table in the database
type Task struct {
	ID          string `json:"id"`
	Detail      string `json:"detail"`
	UserID      uint   `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
