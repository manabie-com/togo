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
			return todo.User{}, domain.NotFound("user not found")
		}
		return todo.User{}, fmt.Errorf("select from user: %w", err)
	}
	return todo.User{
		ID:             id,
		TaskDailyLimit: row.TaskDailyLimit,
	}, nil
}

func (repo *TodoUserRepo) AddUser(ctx context.Context, u todo.User) error {
	q := `INSERT INTO "user" (id, task_daily_limit) VALUES ($1, $2)`
	if _, err := repo.db.ExecContext(ctx, q, u.ID, u.TaskDailyLimit); err != nil {
		return fmt.Errorf("insert user: %w", err)
	}
	return nil
}

type User struct {
	PK             int `db:"pk"`
	TaskDailyLimit int `db:"task_daily_limit"`
}
