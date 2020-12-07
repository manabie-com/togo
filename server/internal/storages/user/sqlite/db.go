package sqlite

import (
	"context"
	"database/sql"
	userEntity "github.com/HoangVyDuong/togo/internal/storages/user"
	"github.com/HoangVyDuong/togo/pkg/logger"
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

const getUser = `
	SELECT id, name, password FROM users WHERE name = ?
`

func (ur *userRepository) GetUserByName(ctx context.Context, name string) (userEntity.User, error) {
	row := ur.DB.QueryRowContext(ctx, getUser, name)
	var i userEntity.User
	err := row.Scan(&i.ID, &i.Name, &i.Password)
	if err != nil {
		logger.Errorf("[UserRepository][GetUserByName] error: %s", err.Error())
	}
	return i, err
}
