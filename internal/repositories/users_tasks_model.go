package repositories

import "time"

type Task struct {
	ID        string     `gorm:"id" json:"id"`
	Content   string     `gorm:"content" json:"content"`
	UserID    string     `gorm:"user_id" json:"user_id"`
	CreatedAt *time.Time `gorm:"created_date" json:"created_at"`
}

type User struct {
	ID        string     `gorm:"id" json:"id"`
	Password  string     `gorm:"password" json:"password"`
	MaxTodo   int        `gorm:"max_todo" json:"max_todo"`
	CreatedAt *time.Time `gorm:"created_date" json:"created_at"`
}
