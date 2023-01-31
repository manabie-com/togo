package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/trinhdaiphuc/togo/database/ent"
	"github.com/trinhdaiphuc/togo/database/ent/predicate"
	taskent "github.com/trinhdaiphuc/togo/database/ent/task"
	"github.com/trinhdaiphuc/togo/internal/dto"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entities.Task) (*entities.Task, error)
	GetTask(ctx context.Context, id int) (*entities.Task, error)
	GetTasks(ctx context.Context, filter *entities.TaskFilter) (*entities.Tasks, error)
	UpdateTask(ctx context.Context, task *entities.Task) (*entities.Task, error)
	DeleteTask(ctx context.Context, id int) error
}

type taskRepositoryImpl struct {
	db *infrastructure.DB
}

var (
	ErrTaskNotFound = fiber.NewError(fiber.StatusNotFound, "Task not found")
	ErrTaskLimit    = fiber.NewError(fiber.StatusBadRequest, "Task limit exceeded")
)

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
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				logrus.Errorf("Error rolling back %v", err)
			}
		} else {
			err := tx.Commit()
			if err != nil {
				logrus.Errorf("Error commit %v", err)
			}
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
		return nil, ErrTaskLimit
	}
	resp, err := tx.Task.Create().
		SetName(task.Name).
		SetContent(task.Content).
		SetUserID(task.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return dto.Task2TaskEntity(resp), nil
}

func (t *taskRepositoryImpl) GetTask(ctx context.Context, id int) (*entities.Task, error) {
	task, err := t.db.Task.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return dto.Task2TaskEntity(task), nil
}

func (t *taskRepositoryImpl) GetTasks(ctx context.Context, filter *entities.TaskFilter) (*entities.Tasks, error) {
	var (
		condition []predicate.Task
	)
	if filter.UserID != 0 {
		condition = append(condition, taskent.UserIDEQ(filter.UserID))
	}
	if filter.Page < 0 {
		filter.Page = 1
	}
	if filter.Limit < 10 {
		filter.Limit = 10
	}
	offset := (filter.Page - 1) * filter.Limit
	resp, err := t.db.Task.Query().
		Where(condition...).
		Offset(offset).
		Limit(filter.Limit).All(ctx)
	if err != nil {
		return nil, err
	}

	total, err := t.db.Task.Query().
		Where(condition...).
		Offset(offset).
		Limit(filter.Limit).Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entities.Tasks{
		Tasks: dto.Tasks2TasksEntity(resp),
		Total: total,
		Page:  filter.Page,
	}, nil
}

func (t *taskRepositoryImpl) UpdateTask(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	resp, err := t.db.Task.UpdateOneID(task.ID).
		SetName(task.Name).
		SetContent(task.Content).
		SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	return dto.Task2TaskEntity(resp), nil
}

func (t *taskRepositoryImpl) DeleteTask(ctx context.Context, id int) error {
	err := t.db.Task.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrTaskNotFound
		}
		return err
	}
	return nil
}
