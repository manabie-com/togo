package model

import "time"

type User struct {
	/* ID will get from keystone*/
	ID string `sql:"size:64;unique;index" json:"id" gorm:"primary_key"`
	/* Equal with "name" of keystone auth request*/
	Name       string `sql:"size:256;unique;index" json:"name"`
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	NumberTask int    `json:"number_task"`
	Enabled    bool   `json:"enabled"`
	Deleted    bool   `json:"deleted" gorm:"type:boolean;default:false"`
}

func (User) TableName() string {
	return "user"
}

type Task struct {
	UserID    string    `json:"user_id"`
	TaskName  string    `json:"task_name"`
	CreatedAt time.Time `gorm:"type:datetime(6)"`
}

func (Task) TableName() string {
	return "task"
}
