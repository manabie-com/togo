package rdbms

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/app/user/model"
	"github.com/manabie-com/togo/internal/util"
)

const SQLValidateUser = `SELECT id FROM users WHERE id = $1 AND password = $2`

type UserStorage struct {
	db         *sql.DB
	driverName string
}

func New(db *sql.DB, driverName string) *UserStorage {
	return &UserStorage{db: db, driverName: driverName}
}

// ValidateUser returns tasks if match userID AND password
func (us UserStorage) ValidateUser(ctx context.Context, userID, pwd string) error {
	stmt := util.PrepareQuery(us.driverName, SQLValidateUser)
	row := us.db.QueryRowContext(ctx, stmt, userID, pwd)
	u := model.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}
