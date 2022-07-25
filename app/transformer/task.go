package transformer

import (
	"github.com/huuthuan-nguyen/manabie/app/model"
	"time"
)

type TaskTransformer struct {
	ID            int       `json:"id"`
	Content       string    `json:"content"`
	PublishedDate string    `json:"published_date"`
	Status        int       `json:"status"`
	CreatedBy     int       `json:"created_by"`
	CreateAt      time.Time `json:"create_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Transform /**
func (task *TaskTransformer) Transform(e any) any {
	taskModel, ok := e.(model.Task)
	if !ok {
		return e
	}

	task.ID = taskModel.ID
	task.Content = taskModel.Content
	task.PublishedDate = taskModel.PublishedDate.Format("2006-01-02")
	task.Status = taskModel.Status
	task.CreatedBy = taskModel.CreatedBy
	task.CreateAt = taskModel.CreatedAt
	task.UpdatedAt = taskModel.UpdatedAt
	return *task
}
