package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/app/models"
	"github.com/manabie-com/togo/internal/services/auth"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
)

type UsecaseImpl struct {
	Store storages.Storages
	DB    *sql.DB
}

func NewUsecase(db *sql.DB) *UsecaseImpl {
	return &UsecaseImpl{
		Store: postgres.NewPosgresql(db),
		DB:    db,
	}
}

func (uc *UsecaseImpl) GetAuthToken(ctx context.Context, username string, pwd string) (string, error) {

	user, err := uc.Store.ValidateUser(ctx, username, pwd)
	if err != nil {
		return "", errors.New("incorrect username or password")
	}

	token, err := auth.CreateToken(int(user.ID))
	if err != nil {
		return "", errors.New("there was an error generating the API token")
	}

	return token, nil
}

func (uc *UsecaseImpl) RetrieveTasks(ctx context.Context, userID uint64, createdDate string) ([]*models.Task, error) {
	tasks, err := uc.Store.RetrieveTasks(ctx, userID, createdDate)
	if err != nil {
		return nil, errors.New("failure to retrieve tasks")
	}
	return tasks, nil
}

func (uc *UsecaseImpl) AddTask(ctx context.Context, userID uint64, task *models.Task) (*models.Task, error) {
	if task.ID == 0 {
		taskId := uuid.New().ID()
		task.ID = uint64(taskId)
	}
	task.UserID = userID
	task.CreatedDate = time.Now().Format("2006-01-02")

	err := uc.Store.AddTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
