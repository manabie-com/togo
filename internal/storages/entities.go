package storages

import "time"

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}

// ------------
type PgUser struct {
	Id       int
	Username string
	PwdHash  string
	MaxTodo  int
}

type PgTask struct {
	Id       int       `json:"id"`
	UsrId    int       `json:"usr_id"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
}

