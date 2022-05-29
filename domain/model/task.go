package model

import "time"

type Task struct {
	Id          string
	Title       string
	Description string
	CreatedBy   int
	CreatedDate string
	CreatedTime time.Time
}
