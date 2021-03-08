package storages

import _ "gorm.io/gorm"

// Task reflects tasks in DB
type Task struct {
	ID          string `gorm:"primaryKey;type:char(36);" json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string   `gorm:"primaryKey;type:char(36);" json:"id"`
	Password string   `json:"password"`
}
