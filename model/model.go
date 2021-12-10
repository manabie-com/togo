package model

import "time"

type User struct {
	Id       int       `json:"id" gorm:"primaryKey"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	MaxTodo  int       `json:"max_todo"`
	Task     []Task    `gorm:"foreignKey:UserId"`
}

type Task struct {
	Id 	        	int  	   `json:"id" gorm:"primaryKey"`
	Content 		string  	`json:"content"`
	UserId 			int     	`json:"user_id"`
	CreatedDate     time.Time   `json:"created_date"`
}
