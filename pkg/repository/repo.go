package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"togo.com/pkg/model"
)

type Repository interface {
	AddTask(ctx context.Context, addParams model.AddTaskParams) error
	CountTaskPerDay(ctx context.Context, userId string, createDate string) (int64, error)
	RetrieveTasks(ctx context.Context, userId string, createDate string) ([]model.Task, error)
	GetLimitPerUser(ctx context.Context, userId string) (int64, error)
	GetUser(ctx context.Context, req model.LoginRequest) (id string, err error)
}
type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) AddTask(ctx context.Context, addParams model.AddTaskParams) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)",
		uuid.New().String(),
		addParams.Content,
		addParams.UserId,
		addParams.CreateDate)

	return err
}

type RetrieveTasksParams struct {
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

func (r repository) RetrieveTasks(ctx context.Context, userId string, createDate string) ([]model.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2", userId, createDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.Task
	for rows.Next() {
		var i model.Task
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.UserID,
			&i.CreatedDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r repository) GetLimitPerUser(ctx context.Context, userId string) (int64, error) {
	row := r.db.QueryRowContext(ctx, "SELECT max_todo FROM users WHERE id = $1", userId)
	var maxTodo int64
	err := row.Scan(&maxTodo)
	return maxTodo, err
}

func (r repository) CountTaskPerDay(ctx context.Context, userId string, createDate string) (int64, error) {
	row := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_date = $2", userId, createDate)
	var count int64
	err := row.Scan(&count)
	return count, err
}

func (r repository) GetUser(ctx context.Context, req model.LoginRequest) (id string, err error) {
	var user = model.User{}
	err = r.db.GetContext(ctx, &user, "SELECT id FROM users WHERE id = $1 AND password = $2", req.UserName, req.Password)
	if err != sql.ErrNoRows && err != nil {
		return "", err
	}
	return user.Id, nil
}
