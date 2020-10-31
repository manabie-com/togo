package usecase

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
)

type togoUsecase struct {
	togoStorages storages.Storages
}

// NewTogoUsecase will create new a togoUsecase object representation of togo.Usecase interface
func NewTogoUsecase(d storages.Storages) togoUsecase {
	return togoUsecase{togoStorages: d}
}
func (t *togoUsecase) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*entities.Task, error) {
	return t.togoStorages.RetrieveTasks(context.Background(), userID, createdDate)
}
func (t *togoUsecase) AddTask(ctx context.Context, task *entities.Task) error {
	return t.togoStorages.AddTask(ctx, task)
}
func (t *togoUsecase) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	return t.togoStorages.ValidateUser(ctx, userID, pwd)
}
