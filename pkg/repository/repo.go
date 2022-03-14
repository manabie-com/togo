package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"togo.com/pkg/model"
)

type Repository interface {
	AddTask(ctx context.Context, addParams model.AddTaskParams) error
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
		addParams.Id,
		addParams.Content,
		addParams.UserId,
		addParams.CreateDate)

	return err
}
func (r repository) GetUser(ctx context.Context, req model.LoginRequest) (id string, err error) {
	var user = model.User{}
	err = r.db.GetContext(ctx, &user, "SELECT id FROM users WHERE id = $1 AND password = $2", req.UserName, req.Password)
	if err != nil {
		return "", err
	}
	return user.Id, err
}
