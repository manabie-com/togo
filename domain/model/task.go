package model

import "time"

type Task struct {
	Id          string
	Title       string
	Description string
	CreatedTime time.Time
	UpdatedTime time.Time
}
