package storages

import "time"

// User reflects tasks in DB
type User struct {
	Id       int
	Username string
	PwdHash  string
	MaxTodo  int
}

// Task reflects tasks in DB
type Task struct {
	Id       int       `json:"id"`
	UsrId    int       `json:"usr_id"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
}

// LEGACY CODE----------------------------

// SqliteTask reflects tasks in DB
type SqliteTask struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// SqliteUser reflects users data from DB
type SqliteUser struct {
	ID       string
	Password string
}
