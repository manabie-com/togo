package repository

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/domains"
)

func NewUserRepositoryImpl(db *sql.DB) domains.UserRepository {
	return UserRepositoryPostgresImpl{
		db: db,
	}
}

type UserRepositoryPostgresImpl struct {
	db *sql.DB
}

func (u UserRepositoryPostgresImpl) VerifyUser(ctx context.Context, request *domains.LoginRequest) (*domains.User, error) {
	var query = `
		SELECT "id", "username", "max_todo"
		FROM 
			users
		WHERE 
			username = $1 AND password = crypt($2, password)
		LIMIT 1 FOR NO KEY UPDATE
	`

	user := &domains.User{}
	err := u.db.QueryRowContext(ctx, query, request.Username, request.Password).Scan(&user.Id, &user.Username, &user.MaxTodo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domains.ErrorNotFound
		}
		return nil, err
	}

	return user, err
}

func (u UserRepositoryPostgresImpl) GetUserById(ctx context.Context, userId int64) (*domains.User, error) {
	var query = `
		SELECT "id", "username", "max_todo"
		FROM 
			users
		WHERE 
			id = $1
		LIMIT 1 FOR NO KEY UPDATE
	`

	user := &domains.User{}
	err := u.db.QueryRowContext(ctx, query, userId).Scan(&user.Id, &user.Username, &user.MaxTodo)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domains.ErrorNotFound
		}
		return nil, err
	}

	return user, err
}
