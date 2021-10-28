package service

import (
	"context"

	"github.com/manabie-com/togo/internal/core/domain"
	"github.com/manabie-com/togo/internal/core/port"
	"github.com/manabie-com/togo/pkg/database"
)

func NewTaskService(db database.Database, taskRepo port.TaskRepository, taskValidator port.TaskValidator) port.TaskService {
	return &taskService{
		db:            db,
		taskRepo:      taskRepo,
		taskValidator: taskValidator,
	}
}

type taskService struct {
	db            database.Database
	taskRepo      port.TaskRepository
	taskValidator port.TaskValidator
}

func (p *taskService) WarmUp(ctx context.Context) error {
	return p.db.Transaction(ctx, func(ctx context.Context, conn database.Connection) error {
		return p.taskRepo.InitTables(ctx, conn)
	})
}

func (p *taskService) RetrieveTasks(ctx context.Context, userId, createdDate string) ([]*domain.Task, error) {
	err := p.taskValidator.ValidateBeforeRetrieveTasks(userId, createdDate)
	if err != nil {
		return nil, err
	}

	var tasks []*domain.Task
	err = p.db.Transaction(ctx, func(ctx context.Context, conn database.Connection) error {
		tasks, err = p.taskRepo.RetrieveTasks(ctx, conn, userId, createdDate)
		return err
	})
	return tasks, err
}

func (p *taskService) AddTask(ctx context.Context, task *domain.Task) error {
	err := p.taskValidator.ValidateBeforeAddTask(task)
	if err != nil {
		return err
	}
	return p.db.Transaction(ctx, func(ctx context.Context, conn database.Connection) error {
		return p.taskRepo.AddTask(ctx, conn, task)
	})
}

func (p *taskService) Login(ctx context.Context, username, password string) (string, error) {
	err := p.taskValidator.ValidateBeforeLogin(username, password)
	if err != nil {
		return "", err
	}
	var userId string
	err = p.db.Transaction(ctx, func(ctx context.Context, conn database.Connection) error {
		userId, err = p.taskRepo.Login(ctx, conn, username, password)
		return err
	})
	return userId, err
}
