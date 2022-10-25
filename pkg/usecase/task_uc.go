package usecase

import (
	"context"
	"errors"
	"time"
	"togo.com/pkg/model"
	"togo.com/pkg/repository"
)

type TaskUseCase interface {
	AddTask(ctx context.Context, userId string, reqAddTask model.AddTaskRequest) (model.AddTaskResponse, error)
	GetListTaskByDate(ctx context.Context, userId string, createDate string) ([]model.Task, error)
}
type taskUseCase struct {
	repo repository.Repository
}

func NewTaskUseCase(repo repository.Repository) TaskUseCase {
	return taskUseCase{repo: repo}
}

func (t taskUseCase) AddTask(ctx context.Context, userId string, reqAddTask model.AddTaskRequest) (model.AddTaskResponse, error) {
	now := time.Now()
	createDate := now.Format("2006-01-02")
	//get limit task by user
	limit, err := t.repo.GetLimitPerUser(ctx, userId)
	if err != nil {
		return model.AddTaskResponse{}, err
	}
	//get count task for user
	count, err := t.repo.CountTaskPerDay(ctx, userId, createDate)
	if err != nil {
		return model.AddTaskResponse{}, err
	}
	if count >= limit {
		return model.AddTaskResponse{}, errors.New("Limit reached for the day ")
	}
	addParams := model.AddTaskParams{
		UserId:     userId,
		CreateDate: createDate,
		Content:    reqAddTask.Content,
	}
	err = t.repo.AddTask(ctx, addParams)
	if err != nil {
		return model.AddTaskResponse{}, err
	}
	return model.AddTaskResponse{UserId: userId, Content: reqAddTask.Content, CreateDate: createDate}, err
}

func (t taskUseCase) GetListTaskByDate(ctx context.Context, userId string, createDate string) ([]model.Task, error) {
	if createDate == "" {
		now := time.Now()
		createDate = now.Format("2006-01-02")
	}
	task, err := t.repo.RetrieveTasks(ctx, userId, createDate)
	return task, err
}
