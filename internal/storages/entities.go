package storages

import "gorm.io/gorm"

// Task reflects tasks in DB
type Task struct {
	gorm.Model
	ID          string `json:"id" gorm:"primary_key:true;column:id;not null"`
	Content     string `json:"content" gorm:"column:content;not null"`
	UserID      string `json:"user_id" gorm:"column:user_id;not null"`
	CreatedDate string `json:"created_date" gorm:"column:created_date;not null"`
	User		*User  `json:"-" gorm:"foreignKey:UserID"`
}

func(Task) TableName() string {
	return "tasks"
}

func(Task) BeforeCreate(tx *gorm.DB) error {
	return nil
}

// User reflects users data from DB
type User struct {
	gorm.Model
	ID       string	`json:"id" gorm:"primary_key;column:id;not null"`
	Password string	`json:"password" gorm:"column:password;not null"`
	MaxTodo  int32 	`json:"max_todo" gorm:"column:max_todo;not null"`
}

func(User) TableName() string{
	return "users"
}

func(User) BeforeCreate(tx *gorm.DB) error {
	return nil
}
