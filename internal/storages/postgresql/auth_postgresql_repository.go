package postgresql

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/domain/entity"
)

type userPostgresqlRepository struct {
	DB *sql.DB
}

func NewUserPostgresqlRepository(db *sql.DB) domain.UserRepository {
	return userPostgresqlRepository{
		DB: db,
	}
}

func (a userPostgresqlRepository) GetUser(ctx context.Context, username string) (*entity.User, error) {
	query := `SELECT username, hashed_password, max_todo FROM users WHERE username = $1`
	stmt, err := a.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, username)
	var usernameQuery, hashedPasswordQuery string
	var maxTodo int
	err = row.Scan(&usernameQuery, &hashedPasswordQuery, &maxTodo)
	if err == sql.ErrNoRows {
		return nil, domain.UserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Username:       usernameQuery,
		HashedPassword: hashedPasswordQuery,
		MaxTodo:        maxTodo,
	}, nil
}
