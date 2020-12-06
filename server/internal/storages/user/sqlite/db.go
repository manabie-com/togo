package sqlite

import (
	"context"
	"database/sql"
	"github.com/HoangVyDuong/togo/internal/storages/user"
	userUsecase "github.com/HoangVyDuong/togo/internal/usecase/user"
)

type userRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) userUsecase.Repository {
	return &userRepository{DB: db}
}


// GetUser returns User if match ID
func (l *userRepository) GetUser(ctx context.Context, id int64) (user.User, error) {
	stmt := `SELECT id FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, id)
	u := user.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}
