package task

import (
	"context"
	"database/sql"

	"manabie/todo/models"
)

type TaskRespository interface {
	Find(ctx context.Context, tx *sql.Tx) ([]*models.Task, error)
}

type taskRespository struct{}

func NewUserRespository() TaskRespository {
	return &taskRespository{}
}

func (tr *taskRespository) Find(ctx context.Context, tx *sql.Tx) ([]*models.Task, error) {
	return nil, nil
}
