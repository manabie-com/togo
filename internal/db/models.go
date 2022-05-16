package db

import (
	"time"
)
type User struct {
	Id int64 `json:"id"`
	UserName string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	MaximumTaskInDay int `json:"maximum_task_in_day"`
}

type Task struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	UserId int64 `json:"user_id"`
}