package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/manabie-com/togo/internal/constants"
	"github.com/manabie-com/togo/internal/dtos"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
	"time"
)

type taskService struct {
	configurationService ConfigurationService
	taskRepository       repositories.TaskRepository
}

type TaskService interface {
	CreateTask(ctx context.Context, request *dtos.CreateTaskRequest) (*dtos.CreateTaskResponse, error)
	GetListTask(ctx context.Context, userID string, createdDate string) (*dtos.GetListTaskResponse, error)
}

func NewTaskService(injectedConfigurationService ConfigurationService,
	injectedTaskRepository repositories.TaskRepository) TaskService {
	return &taskService{
		configurationService: injectedConfigurationService,
		taskRepository:       injectedTaskRepository,
	}
}

func (s *taskService) CreateTask(ctx context.Context, request *dtos.CreateTaskRequest) (*dtos.CreateTaskResponse, error) {
	userID, createdDate := request.UserID, time.Now().Format("2006-01-02")
	ok, err := s.ValidateCapacity(ctx, userID, createdDate)
	if err != nil {
		logrus.Errorf("Validate Capacity error: %s", err.Error())
		return nil, err
	}

	if !ok {
		logrus.Errorf("Maximum Capacity")
		return nil, constants.ErrMaximumCreatedTask
	}

	createdTask, err := s.taskRepository.Create(ctx, &models.Task{
		ID:          uuid.New().String(),
		Content:     request.Content,
		UserID:      userID,
		CreatedDate: createdDate,
	})
	if err != nil {
		logrus.Errorf("Create Task error: %s", err.Error())
		return nil, err
	}

	var response = &dtos.TaskDto{}
	if err := copier.Copy(response, createdTask); err != nil {
		logrus.Errorf("Create Task Mapping Dto error: %s", err.Error())
		return nil, err
	}

	return &dtos.CreateTaskResponse{Data: response}, nil
}

func (s *taskService) GetListTask(ctx context.Context, userID string, createdDate string) (*dtos.GetListTaskResponse, error) {
	tasks, err := s.taskRepository.GetTasks(ctx, userID, createdDate)
	if err != nil {
		logrus.Errorf("Get Task List error: %s", err.Error())
		return nil, err
	}

	var response []*dtos.TaskDto
	if err := copier.Copy(&response, tasks); err != nil {
		logrus.Errorf("Get Task List Mapping Dto error: %s", err.Error())
		return nil, err
	}

	return &dtos.GetListTaskResponse{Data: response}, nil
}

func (s *taskService) ValidateCapacity(ctx context.Context, userID string, createdDate string) (bool, error) {
	totalCreated, err := s.taskRepository.CountTotalInDate(ctx, userID, createdDate)
	if err != nil {
		logrus.Errorf("Get Task List error: %s", err.Error())
		return false, err
	}
	logrus.Infof("Total Created Task in current date %d", totalCreated)

	capacity, err := s.configurationService.GetTaskConfiguration(ctx, userID, createdDate)
	if err != nil {
		logrus.Errorf("Get Task Configuration error: %s", err.Error())
		return false, err
	}
	logrus.Infof("Get Task configuration %d", capacity)

	if totalCreated == capacity {
		logrus.Infof("Maximum Created task, prevent continuing")
		return false, constants.ErrMaximumCreatedTask
	}

	return true, err
}
