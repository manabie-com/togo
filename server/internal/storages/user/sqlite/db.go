package sqlite

import (
	"context"
	"database/sql"
	"github.com/HoangVyDuong/togo/internal/storages/user"
)

type userRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *userRepository {
	return &userRepository{DB: db}
}

func (ur *userRepository) Close() error {
	return ur.DB.Close()
}

// GetUser returns User if match ID
func (ur *userRepository) GetUserByName(ctx context.Context, name string) (user.User, error) {
	return user.User{}, nil
}
