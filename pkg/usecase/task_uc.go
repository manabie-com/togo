package usecase

import (
	"context"
	"github.com/google/uuid"
	"time"
	"togo.com/pkg/model"
	"togo.com/pkg/repository"
)

type TaskUseCase interface {
	AddTask(ctx context.Context, userId string, reqAddTask model.AddTaskRequest) (model.AddTaskResponse, error)
}
type taskUseCase struct {
	repo repository.Repository
}

func NewTaskUseCase(repo repository.Repository) TaskUseCase {
	return taskUseCase{repo: repo}
}

func (t taskUseCase) AddTask(ctx context.Context, userId string, reqAddTask model.AddTaskRequest) (model.AddTaskResponse, error) {
	//TODO check limit task for user in day.
	now := time.Now()
	createDate := now.Format("2006-01-02")
	addParams := model.AddTaskParams{
		Id:         uuid.New().String(),
		UserId:     userId,
		CreateDate: createDate,
		Content:    reqAddTask.Content,
	}
	err := t.repo.AddTask(ctx, addParams)
	return model.AddTaskResponse{UserId: userId, Content: reqAddTask.Content, CreateDate: createDate}, err
}
