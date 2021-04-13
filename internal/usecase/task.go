package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/manabie-com/togo/internal/entities"
)

var MaxTask = 5

func (t *uc) List(ctx context.Context, id, createdAt string) ([]*entities.Task, error) {
	fmt.Println("id:", id)
	if id == "" {
		return nil, nil
	}
	return t.task.RetrieveTasks(ctx, sql.NullString{String: id, Valid: true}, sql.NullString{String: createdAt, Valid: true})
}

func (t *uc) Add(ctx context.Context, id, date string, task *entities.Task) error {
	if task == nil {
		return errors.New("data empty")
	}

	list, err := t.List(ctx, id, date)
	if err != nil {
		return err
	}
	fmt.Println("Len Task:", len(list), err)
	if len(list) >= MaxTask {
		return errors.New("quantity exceeded allowed")
	}

	return t.task.AddTask(ctx, task)
}
