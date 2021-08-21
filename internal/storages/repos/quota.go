package repos

import (
	"context"
	"database/sql"
)

type IQuotaRepo interface {
	CountTaskPerDay(ctx context.Context, userID string, dateStr string) int
	GetLimitPerUser(ctx context.Context, userID string) int
}

type QuotaRepo struct {
	db *sql.DB
}

func (l *QuotaRepo) CountTaskPerDay(ctx context.Context, userID string, dateStr string) int {
	stmt := `SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date = ?`
	row := l.db.QueryRowContext(ctx, stmt, userID, dateStr)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0
	}

	return count
}

func (l *QuotaRepo) GetLimitPerUser(ctx context.Context, userID string) int {
	stmt := `SELECT max_todo FROM users WHERE id = ?`
	row := l.db.QueryRowContext(ctx, stmt, userID)
	var maxCount int
	err := row.Scan(&maxCount)
	if err != nil {
		return 0
	}
	return maxCount
}

func NewQuotaRepo(db *sql.DB) IQuotaRepo {
	return &QuotaRepo{
		db: db,
	}
}
