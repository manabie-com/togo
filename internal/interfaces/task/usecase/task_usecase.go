package usecase

import (
	"context"
	"database/sql"

	"github.com/datshiro/togo-manabie/internal/infras/errors"
	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/datshiro/togo-manabie/internal/interfaces/service/cache"
	postgres_task "github.com/datshiro/togo-manabie/internal/interfaces/task/repository/postgres"
)

type taskUseCase struct {
	taskRepo     domain.TaskRepository
	CacheService cache.CacheService
	DB           *sql.DB
}

func NewTaskUseCase(dbc *sql.DB, cacheService cache.CacheService) domain.TaskUseCase {
	return &taskUseCase{
		DB:           dbc,
		CacheService: cacheService,
		taskRepo:     postgres_task.NewTaskRepository(),
	}
}

func (t *taskUseCase) CreateTask(ctx context.Context, m *models.Task, user *models.User) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	valid, err := t.CacheService.ValidateQuota(ctx, user)
	if err != nil {
		return err
	}
	if !valid {
		return errors.CustomError("Err quota exceeded")
	}

	if err := t.CacheService.IncreaseQuota(ctx, user); err != nil {
		return err
	}
	if err := t.taskRepo.CreateOne(ctx, tx, m); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
