package postgres

import (
	"context"

	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type taskRepository struct{}

func NewTaskRepository() domain.TaskRepository {
	return &taskRepository{}
}

func (t *taskRepository) CreateOne(ctx context.Context, exec boil.ContextExecutor, m *models.Task) error {
	return m.Insert(ctx, exec, boil.Infer())
}
