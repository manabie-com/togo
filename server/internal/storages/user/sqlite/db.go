package sqlite

import (
	"context"
	"database/sql"
	"github.com/HoangVyDuong/togo/internal/storages"
)

type LiteDB struct {
	DB *sql.DB
}

// GetUser returns User if match ID
func (l *LiteDB) GetUser(ctx context.Context, userID int64) bool {
	stmt := `SELECT id FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}
	return true
}
