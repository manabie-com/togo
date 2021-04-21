package respository

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/model"
)

type UserLiteDB struct {
	DB *sql.DB
}

func NewUserLiteDBRespository(db *sql.DB) model.UserRespository {
	return &UserLiteDB{
		DB: db,
	}
}

// ValidateUser returns tasks if match userID AND password
func (l *UserLiteDB) ValidateUser(ctx context.Context, userID, pwd string) (bool, error) {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`

	row := l.DB.QueryRowContext(ctx, stmt, convertToNullString(userID), convertToNullString(pwd))
	u := &model.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false, err
	}

	return true, nil
}
