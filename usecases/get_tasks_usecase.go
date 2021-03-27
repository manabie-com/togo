package usecases

import (
	"context"
	"github.com/manabie-com/togo/domains"
	"time"
)

type (
	GetTasksUseCase interface {
		Execute(context.Context, GetTasksInput) ([]*TaskOutput, error)
	}

	GetTasksInput struct {
		UserId      int64     `json:"user_id"`
		CreatedDate time.Time `json:"created_date"`
	}

	// Output data
	TaskOutput struct {
		ID          int64  `json:"id"`
		Content     string `json:"content"`
		CreatedDate string `json:"created_date"`
	}

	getTasksInteractor struct {
		taskRepo domains.TaskRepository
	}
)

func NewGetTasksUseCase(taskRepository domains.TaskRepository) GetTasksUseCase {
	return getTasksInteractor{
		taskRepo: taskRepository,
	}
}

// Execute create Hub with dependencies
func (i getTasksInteractor) Execute(ctx context.Context, req GetTasksInput) ([]*TaskOutput, error) {
	taskReq := &domains.TaskRequest{
		UserId: req.UserId,
	}

	if !req.CreatedDate.IsZero() {
		taskReq.CreatedDate = req.CreatedDate
	}

	tasks, err := i.taskRepo.GetTasks(ctx, taskReq)

	if err != nil {
		return nil, err
	}

	return transformTasksToSliceTaskOutput(tasks), nil
}

func transformTasksToSliceTaskOutput(tasks []*domains.Task) []*TaskOutput {
	result := make([]*TaskOutput, 0)
	tsLayout := time.RFC3339
	for _, u := range tasks {
		result = append(result, &TaskOutput{
			ID:          u.Id,
			Content:     u.Content,
			CreatedDate: u.CreatedDate.Format(tsLayout),
		})
	}
	return result
}
