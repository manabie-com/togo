package services

import (
	"context"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages"
	"time"
)

type TaskService interface {
	CreateTask(ctx context.Context, req model.TaskCreationRequest) (*model.Task, error)

	GetTaskByDate(ctx context.Context, queryDate time.Time) (*[]model.Task, error)
}

type taskServiceImpl struct {
	taskRepository storages.TaskRepository
	userRepository storages.UserRepository
}

func NewTaskService(taskRepository storages.TaskRepository, userRepository storages.UserRepository) *taskServiceImpl {
	return &taskServiceImpl{taskRepository: taskRepository, userRepository: userRepository}
}

func (t *taskServiceImpl) CreateTask(ctx context.Context, req model.TaskCreationRequest) (*model.Task, error) {
	userId := ctx.Value("userId").(string)

	owner, err := t.userRepository.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	newTask, err := t.taskRepository.CreateTask(ctx, req.Content, owner)
	if err != nil {
		return nil, err
	}

	return &model.Task{
		TaskID:      newTask.TaskID,
		Content:     newTask.Content,
		CreatedDate: newTask.CreatedDate,
		UserID:      userId,
	}, nil
}

func (t *taskServiceImpl) GetTaskByDate(ctx context.Context, queryDate time.Time) (*[]model.Task, error) {
	userId := ctx.Value("userId").(string)
	nextDate := queryDate.Add(time.Hour * 24)

	tasks, err := t.taskRepository.GetTaskByDate(ctx, userId, queryDate, nextDate)
	if err != nil {
		return nil, err
	}

	allTask := []model.Task{}

	for _, e := range tasks {
		elems := model.Task{
			TaskID:      e.TaskID,
			Content:     e.Content,
			CreatedDate: e.CreatedDate,
			UserID:      userId,
		}
		allTask = append(allTask, elems)

	}

	return &allTask, nil
}
