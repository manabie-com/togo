package postgres

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/pkg/errors"
)

type PGUserRepository struct {
	PGRepository
}

func NewPGUserRepository(dbConn *sqlx.DB) *PGUserRepository {
	return &PGUserRepository{PGRepository{dbConn}}
}

func (r *PGUserRepository) GetByCredentials(ctx context.Context, username, password string) (*d.User, error) {
	user := d.User{}
	err := r.DBConn.GetContext(
		ctx, &user,
		"SELECT * FROM users WHERE username = $1 AND password = crypt($2, password)",
		username, password)

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "db error")
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, nil
}
