package task

import "time"

type Task struct {
	Id          string
	Content     string
	UserId      string
	CreatedDate time.Time
}
