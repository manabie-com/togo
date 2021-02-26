package storages

import "time"

// PgUser reflects tasks in DB
type PgUser struct {
	Id       int
	Username string
	PwdHash  string
	MaxTodo  int
}

// PgTask reflects tasks in DB
type PgTask struct {
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