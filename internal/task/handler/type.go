package handler

import (
	"time"

	"github.com/manabie-com/togo/internal/task/service"

	"github.com/go-playground/validator/v10"
	"github.com/manabie-com/togo/pkg/errorx"
)

type CreateTaskRequest struct {
	Content string `json:"content"`
}

func (p *CreateTaskRequest) Validate() error {
	if err := validator.New().Struct(p); err != nil {
		return errorx.ErrInvalidParameter(err)
	}

	return nil
}

type UpdateTaskRequest struct {
	Content string `json:"content"`
}

type Task struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convertServiceTaskToHandlerTask(args *service.Task) *Task {
	if args == nil {
		return nil
	}
	return &Task{
		ID:        args.ID,
		Content:   args.Content,
		UserID:    args.UserID,
		CreatedAt: args.CreatedAt,
		UpdatedAt: args.UpdatedAt,
	}
}

func convertServiceTasksToHandlerTasks(args []*service.Task) []*Task {
	var res []*Task
	for _, v := range args {
		res = append(res, convertServiceTaskToHandlerTask(v))
	}
	return res
}
