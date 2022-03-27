package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/laghodessa/togo/domain/todo"
)

func NewTodoTaskRepo(db *sql.DB) *TodoTaskRepo {
	return &TodoTaskRepo{
		db: sqlx.NewDb(db, "postgres"),
	}
}

var _ todo.TaskRepo = (*TodoTaskRepo)(nil)

// TodoTaskRepo implements todo.TaskRepo
type TodoTaskRepo struct {
	db *sqlx.DB
}

func (repo *TodoTaskRepo) AddTask(ctx context.Context, task todo.Task, loc *time.Location, dailyLimit int) error {
	q := `
WITH user_pk AS (SELECT pk FROM "user" WHERE id=$3)
INSERT INTO task (id, message, user_pk)
SELECT $1, $2, (SELECT * FROM user_pk)
WHERE (SELECT COUNT(*) FROM task
	WHERE user_pk=(SELECT * FROM user_pk)
	AND created_at >= $4 AND created_at < $5) < $6
`

	now := time.Now().In(loc)
	year, month, day := now.Date()
	dayStart := time.Date(year, month, day, 0, 0, 0, 0, loc)
	dayEnd := time.Date(year, month, day+1, 0, 0, 0, 0, loc)

	result, err := repo.db.ExecContext(ctx, q, task.ID, task.Message, task.UserID, dayStart, dayEnd, dailyLimit)
	if err != nil {
		return fmt.Errorf("insert task: %w", err)
	}
	if n, _ := result.RowsAffected(); n == 0 {
		return todo.ErrUserHitTaskDailyLimit
	}
	return nil
}

func (repo *TodoTaskRepo) CountInTimeRangeByUserID(ctx context.Context, userID string, since time.Time, until time.Time) (count int, _ error) {
	q := `SELECT COUNT(*) FROM task
WHERE user_pk=(SELECT pk FROM "user" WHERE id=$1)
	AND created_at >= $2 AND created_at < $3
`
	if err := repo.db.GetContext(ctx, &count, q, userID, since, until); err != nil {
		return 0, err
	}
	return
}
