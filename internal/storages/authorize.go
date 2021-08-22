package storages

import (
	"context"
	"database/sql"
)

type IAuthorizeRepo interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}

type AuthorizeRepo struct {
	db DBTX
}

const QueryValidateUser = `SELECT id FROM users WHERE id = ? AND password = ?`

// ValidateUser returns tasks if match userID AND password
func (l *AuthorizeRepo) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := QueryValidateUser
	row := l.db.QueryRowContext(ctx, stmt, userID, pwd)
	u := &User{}
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
