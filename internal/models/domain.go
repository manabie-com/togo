package models

type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

func (Task) TableName() string {
	return "tasks"
}

type User struct {
	ID       string
	Password string
}

func (User) TableName() string {
	return "users"
}

type Configuration struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Capacity int64  `json:"capacity"`
	Date     string `json:"date"`
}

func (Configuration) TableName() string {
	return "configurations"
}
