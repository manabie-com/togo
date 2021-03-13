package request

import "time"

type GetTasks struct {
	CreatedDate time.Time `validate:"required"`
}

type CreateTask struct {
	Content string `validate:"required"`
}
