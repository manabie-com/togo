package sql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/trangmaiq/togo/internal/model"
)

type Persister struct {
	db *sqlx.DB
}

func NewPersister(db *sqlx.DB) *Persister {
	return &Persister{db: db}
}

func (p *Persister) CreateTask(ctx context.Context, task *model.Task) error {
	_, err := p.db.NamedExec(`
			INSERT INTO tasks 
			    (id, user_id, title, note, status, created_at, updated_at)
			VALUES (:id, :user_id, :title, :note, :status, :created_at, :updated_at)`, task,
	)
	if err != nil {
		return fmt.Errorf("create task failed: %w", err)
	}

	return nil
}

func (p *Persister) GetTask(ctx context.Context, id string) (*model.Task, error) {
	var task model.Task
	err := p.db.Get(&task, `SELECT * FROM tasks WHERE id=$1`, id)
	if err != nil {
		return nil, fmt.Errorf("get task by id failed: %w", err)
	}

	return &task, nil
}
