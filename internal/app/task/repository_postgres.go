package task

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dinhquockhanh/togo/internal/pkg/errors"
	db "github.com/dinhquockhanh/togo/internal/pkg/sql/sqlc"
)

type (
	PostgresRepository struct {
		q *db.Queries
	}
)

func NewPostgresRepository(cnn *sql.DB) Repository {
	return &PostgresRepository{
		q: db.New(cnn),
	}
}

func (r *PostgresRepository) CreateTask(ctx context.Context, req *CreateTaskReq) (*Task, error) {
	task, err := r.q.CreateTask(ctx, &db.CreateTaskParams{
		Name:        req.Name,
		Assignee:    sql.NullString{String: req.Assignee, Valid: req.Assignee != ""},
		AssignDate:  time.Now(),
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Status:      req.Status,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Creator:     req.Creator,
	})

	if err != nil {
		op := "postgresRepository.CreateTask"
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return Convert(task), nil
}

func (r *PostgresRepository) AssignTask(ctx context.Context, req *AssignTaskReq) (*Task, error) {
	task, err := r.q.AssignTask(ctx, &db.AssignTaskParams{
		ID: int32(req.ID),
		Assignee: sql.NullString{
			String: req.Assignee,
			Valid:  req.Assignee != "",
		},
	})
	if err != nil {
		op := "postgresRepository.AssignTask"
		if errors.IsSQLNotFound(err) {
			return nil, errors.NewNotFoundErr(err, op, fmt.Sprintf("task with id = %d", req.ID))
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return Convert(task), nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, req *GetTaskByIdReq) (*Task, error) {
	task, err := r.q.GetTask(ctx, int32(req.ID))
	if err != nil {
		op := "postgresRepository.GetByUserName"
		if errors.IsSQLNotFound(err) {
			return nil, errors.NewNotFoundErr(err, op, fmt.Sprintf("task with id = %d", req.ID))
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return Convert(task), nil

}
func (r *PostgresRepository) CountTasksOfUserToDay(ctx context.Context, username string) (int, error) {
	c, err := r.q.CountTaskByAssigneeToday(ctx, sql.NullString{
		String: username,
		Valid:  username != "",
	})
	if err != nil {
		op := "postgresRepository.CountTasksOfUserToDay"
		if errors.IsSQLNotFound(err) {
			return 0, nil
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return int(c), nil
}

func (r *PostgresRepository) ListTasks(ctx context.Context, req *ListTasksReq) ([]*Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepository) Delete(ctx context.Context, req *DeleteTaskByIdReq) error {
	err := r.q.DeleteTask(ctx, int32(req.ID))
	if err != nil {
		op := "postgresRepository.Delete"
		if errors.IsSQLNotFound(err) {
			return errors.NewNotFoundErr(err, op, fmt.Sprintf("task with id = %d", req.ID))
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
