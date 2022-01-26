package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/trinhdaiphuc/togo/database/ent"
	taskent "github.com/trinhdaiphuc/togo/database/ent/task"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
	"time"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entities.Task) (*entities.Task, error)
}

type taskRepositoryImpl struct {
	db *infrastructure.DB
}

func NewTaskRepository(db *infrastructure.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: db,
	}
}

func (t *taskRepositoryImpl) Create(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	tx, err := t.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	user, err := tx.User.Get(ctx, task.UserID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	timeNow := time.Now().Local()
	currentDate := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
	count, err := tx.Task.Query().
		Where(taskent.UserIDEQ(user.ID), taskent.CreatedAtGTE(currentDate)).
		Count(ctx)
	if err != nil {
		return nil, err
	}
	if count >= user.TaskLimit {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Task limit exceeded")
	}
	resp, err := tx.Task.Create().
		SetName(task.Name).
		SetContent(task.Content).
		SetUserID(task.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &entities.Task{
		ID:        resp.ID,
		Name:      resp.Name,
		Content:   resp.Content,
		UserID:    resp.UserID,
		CreatedAt: resp.CreatedAt.Format(time.RFC3339),
		UpdatedAt: resp.UpdatedAt.Format(time.RFC3339),
	}, nil
}
