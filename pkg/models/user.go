package models

type User struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	UserName  string `json:"user_name" gorm:"column:user_name"`
	Password  string `json:"pass_word" gorm:"column:pass_word"`
	LimitTask int64  `json:"limit_task" gorm:"column:limit_task"`
}
