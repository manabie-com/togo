package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/domain/todo"
)

func NewTodoUserRepo(db *sql.DB) *TodoUserRepo {
	return &TodoUserRepo{
		db: sqlx.NewDb(db, "postgres"),
	}
}

var _ todo.UserRepo = (*TodoUserRepo)(nil)

// TodoUserRepo implements todo.UserRepo
type TodoUserRepo struct {
	db *sqlx.DB
}

func (repo *TodoUserRepo) GetUser(ctx context.Context, id string) (todo.User, error) {
	q := `SELECT task_daily_limit FROM "user" WHERE id=$1`

	var row User
	if err := repo.db.GetContext(ctx, &row, q, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return todo.User{}, fmt.Errorf("%w: user not found", domain.ErrNotFound)
		}
		return todo.User{}, fmt.Errorf("select from user: %w", err)
	}
	return todo.User{
		ID:             id,
		TaskDailyLimit: row.TaskDailyLimit,
	}, nil
}

type User struct {
	PK             int `db:"pk"`
	TaskDailyLimit int `db:"task_daily_limit"`
}
