package repos

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
)

type IAuthorizeRepo interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}

type AuthorizeRepo struct {
	db *sql.DB
}

// ValidateUser returns tasks if match userID AND password
func (l *AuthorizeRepo) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.db.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

func NewAuthorizeRepo(db *sql.DB) IAuthorizeRepo {
	return &AuthorizeRepo{
		db: db,
	}
}
