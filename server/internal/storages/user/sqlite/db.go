package sqlite

import (
	"context"
	"database/sql"
	"github.com/HoangVyDuong/togo/internal/storages/user"
)

type LiteDB struct {
	DB *sql.DB
}

// GetUser returns User if match ID
func (l *LiteDB) GetUser(ctx context.Context, id int64) (user.User, error) {
	stmt := `SELECT id FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, id)
	u := user.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}
