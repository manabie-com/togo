package app

import (
	"context"
	"fmt"
	"time"

	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/domain/todo"
)

type TodoUsecase struct {
	TaskRepo todo.TaskRepo
	UserRepo todo.UserRepo
}

// AddTask add one specified task
//
// timezone is the client timezone
type AddTask struct {
	Task     Task
	TimeZone string
}

func (uc *TodoUsecase) AddTask(ctx context.Context, req AddTask) (todo.Task, error) {
	user, err := uc.UserRepo.GetUser(ctx, req.Task.UserID)
	if err != nil {
		return todo.Task{}, fmt.Errorf("get user: %w", err)
	}

	// calculate the begin & end of day
	// use timezone to avoid DST related issues
	loc, err := time.LoadLocation(req.TimeZone)
	if err != nil {
		return todo.Task{}, fmt.Errorf("%w: invalid timezone", domain.ErrInvalidArg)
	}
	now := time.Now().In(loc)
	year, month, day := now.Date()
	countSince := time.Date(year, month, day, 0, 0, 0, 0, loc)
	countUntil := time.Date(year, month, day+1, 0, 0, 0, 0, loc)

	nTasks, err := uc.TaskRepo.CountInTimeRangeByUserID(ctx, req.Task.UserID, countSince, countUntil)
	if err != nil {
		return todo.Task{}, fmt.Errorf("count tasks by user id: %w", err)
	}

	// checking daily limit here is unreliable,
	// but in most cases, this will suffice
	// repo should handle set-validation
	if err := user.HitTaskDailyLimit(nTasks); err != nil {
		return todo.Task{}, err
	}

	task, err := todo.NewTask(
		todo.TaskUserID(req.Task.UserID),
		todo.TaskMessage(req.Task.Message),
	)
	if err != nil {
		return todo.Task{}, err
	}

	if err := uc.TaskRepo.AddTask(ctx, task, loc, user.TaskDailyLimit); err != nil {
		return todo.Task{}, fmt.Errorf("add task: %w", err)
	}
	return task, nil
}
