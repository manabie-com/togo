package task

import (
	"strings"

	"github.com/manabie-com/backend/entity"
	"github.com/manabie-com/backend/repository"
	"github.com/manabie-com/backend/utils"
	uuid "github.com/satori/go.uuid"
)

type I_TaskService interface {
	CreateTask(*entity.Task) (*entity.Task, *utils.ErrorRest)
	UpdateTask(*entity.Task) (*entity.Task, *utils.ErrorRest)
	GetTaskAll() ([]entity.Task, *utils.ErrorRest)
	DeleteTask(id string) *utils.ErrorRest
}

type Service struct {
	Repo repository.I_Repository
}

func NewTaskService(repo repository.I_Repository) I_TaskService {

	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateTask(task *entity.Task) (*entity.Task, *utils.ErrorRest) {
	task.ID = uuid.NewV4().String()
	task.CreatedDate = utils.GetNowFormat()
	if len(strings.TrimSpace(task.Status)) == 0 {
		task.Status = "ACTIVE"
	}

	if err := s.Repo.CreateTask(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Service) UpdateTask(task *entity.Task) (*entity.Task, *utils.ErrorRest) {
	existTask, err := s.Repo.GetTask(task.ID)
	if err != nil {
		return nil, err
	}

	task.Content = existTask.Content
	task.CreatedDate = existTask.CreatedDate
	task.UserID = existTask.UserID

	if err := s.Repo.UpdateTask(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *Service) DeleteTask(id string) *utils.ErrorRest {
	_, err := s.Repo.GetTask(id)
	if err != nil {
		return err
	}

	if err := s.Repo.DeleteTask(id); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetTaskAll() ([]entity.Task, *utils.ErrorRest) {
	tasks, err := s.Repo.GetTaskAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
