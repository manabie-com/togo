package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/chi07/todo/internal/model"
)

type Task struct {
	db *sqlx.DB
}

func NewTask(db *sqlx.DB) *Task {
	return &Task{db: db}
}

func (repo *Task) CountUserTasks(ctx context.Context, userID uuid.UUID) (int64, error) {
	var total int64
	err := repo.db.GetContext(ctx, &total, "SELECT count(*) as total FROM tasks WHERE DATE(created_at)=DATE(NOW()) AND user_id=$1", userID)
	return total, errors.Wrap(err, "cannot count user task from DB")
}

func (repo *Task) Create(ctx context.Context, t *model.Task) (uuid.UUID, error) {
	query, err := repo.db.PrepareNamedContext(ctx, "INSERT INTO tasks (id, title, status, priority, user_id, created_at, updated_at) VALUES(:id, :title, :status, :priority, :user_id, :created_at, :updated_at) RETURNING id")
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "cannot save new task to DB")
	}

	var taskID uuid.UUID
	err = query.Get(&taskID, t)
	if err != nil {
		return uuid.Nil, err
	}
	return taskID, nil
}
