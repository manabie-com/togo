package repos

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type IQuotaRepo interface {
	CountTaskPerDay(ctx context.Context, userID string, dateStr string) (int, *tools.TodoError)
	GetLimitPerUser(ctx context.Context, userID string) (int, *tools.TodoError)
}

type QuotaRepo struct {
	db *sql.DB
}

const QueryCountTaskPerDay = `SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date = ?`

func (l *QuotaRepo) CountTaskPerDay(ctx context.Context, userID string, dateStr string) (int, *tools.TodoError) {
	stmt := QueryCountTaskPerDay
	row := l.db.QueryRowContext(ctx, stmt, userID, dateStr)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}

	return count, nil
}

const QueryGetLimitPerUser = `SELECT max_todo FROM users WHERE id = ?`

func (l *QuotaRepo) GetLimitPerUser(ctx context.Context, userID string) (int, *tools.TodoError) {
	stmt := QueryGetLimitPerUser
	row := l.db.QueryRowContext(ctx, stmt, userID)
	var maxCount int
	err := row.Scan(&maxCount)
	if err != nil {
		return 0, tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}
	return maxCount, nil
}

func NewQuotaRepo(db *sql.DB) IQuotaRepo {
	return &QuotaRepo{
		db: db,
	}
}
