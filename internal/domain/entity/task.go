package entity

import "time"

type Task struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}
