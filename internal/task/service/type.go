package service

import (
	"time"

	"github.com/manabie-com/togo/model"
)

type GetTaskArgs struct {
	ID int
}

type Task struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateTaskArgs struct {
	Content string
}

type UpdateTaskArgs struct {
	TaskID  int
	Content string
}

type DeleteTaskArgs struct {
	ID int
}

type GetTasksArgs struct {
	Limit  int
	Offset int
}

func convertModelTaskToServiceTask(args *model.Task) *Task {
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

func convertModelTasksToServiceTasks(args []*model.Task) []*Task {
	var res []*Task
	for _, v := range args {
		res = append(res, convertModelTaskToServiceTask(v))
	}
	return res
}
