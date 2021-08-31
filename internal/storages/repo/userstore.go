package repo

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

// UserStore for working with users table
type UserStore struct {
	DB *sql.DB
}

// RetrieveUser return user information by given user ID
func (this *UserStore) RetrieveUser(ctx context.Context, userID string) (*storages.User, error) {
	stmt := `SELECT id
	              , password
				  , max_todo
			   FROM users
			  WHERE id = $1`
	row := this.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.Password, &u.MaxTodo)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// AddUser adds a new user to DB
func (this *UserStore) AddUser(ctx context.Context, userId, password string) error {
	stmt := `INSERT INTO users (id, password) VALUES ($1, $2)`
	_, err := this.DB.ExecContext(ctx, stmt, &userId, &password)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
/*func (this *UserStore) ValidateUser(ctx context.Context, userID, pwd string) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := this.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}*/
