package service

import (
	"fmt"
	"github.com/ansidev/togo/domain/task"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/ansidev/togo/task/dto"
	"github.com/pkg/errors"
	"time"
)

type ITaskService interface {
	Create(request dto.CreateTaskRequest, userId int64) (dto.CreateTaskResponse, error)
}

func NewTaskService(userRepository user.IUserRepository, taskRepository task.ITaskRepository) ITaskService {
	return &TaskService{userRepository, taskRepository}
}

type TaskService struct {
	userRepository user.IUserRepository
	taskRepository task.ITaskRepository
}

func (s *TaskService) Create(request dto.CreateTaskRequest, userId int64) (dto.CreateTaskResponse, error) {
	userModel, err := s.userRepository.FindFirstByID(userId)

	if err != nil {
		if errors.Cause(err) == errs.ErrRecordNotFound {
			return dto.CreateTaskResponse{},
				errs.New(errs.ErrUsernameNotFound).
					WithCode(errs.ErrCodeUsernameNotFound).
					WithErr(err).
					Build()
		} else {
			return dto.CreateTaskResponse{},
				errs.New(errs.ErrInternalServiceError).
					WithCode(errs.ErrCodeDbError).
					WithErr(err).
					Build()
		}
	}

	now := time.Now()
	totalTodayTasksByUser, err1 := s.taskRepository.GetTotalTasksByUserAndDate(userModel, now)
	if err1 != nil {
		return dto.CreateTaskResponse{}, err1
	}

	if totalTodayTasksByUser >= int64(userModel.MaxDailyTask) {
		return dto.CreateTaskResponse{},
			errs.New(errs.ErrReachedLimitDailyTask).
				WithCode(errs.ErrCodeReachedLimitDailyTask).
				WithErr(fmt.Errorf("total tasks today is gte max daily task (%d >= %d)", totalTodayTasksByUser, userModel.MaxDailyTask)).
				Build()
	}

	taskModel := task.Task{
		Title:     request.Title,
		UserID:    userModel.ID,
		User:      userModel,
		CreatedAt: time.Now(),
	}

	createdTask, err2 := s.taskRepository.Create(taskModel, userModel)
	if err2 != nil {
		return dto.CreateTaskResponse{}, err2
	}

	taskResponse := dto.CreateTaskResponse{
		ID:        createdTask.ID,
		Title:     createdTask.Title,
		OwnerID:   createdTask.UserID,
		CreatedAt: createdTask.CreatedAt.Format(time.RFC3339),
		UpdatedAt: createdTask.UpdatedAt.Format(time.RFC3339),
	}

	return taskResponse, nil
}
