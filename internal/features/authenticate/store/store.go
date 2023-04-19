// Package store contains create todo related functionality.
package store

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"github.com/jmoiron/sqlx"
	"github.com/manabie-com/togo/internal/features/authenticate"
	"github.com/manabie-com/togo/platform/database"
	"go.uber.org/zap"
)

// Store manages the set of APIs for user database access.
type Store struct {
	log    *zap.SugaredLogger
	db     sqlx.ExtContext
	inTran bool
}

// NewStore constructs the api for data access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// GetUserByEmail gets the specified user from the database by email.
func (s *Store) GetUserByEmail(ctx context.Context, email mail.Address) (authenticate.User, error) {
	data := struct {
		Email string `db:"email"`
	}{
		Email: email.Address,
	}

	const q = `
	SELECT
		id, email, password_hash
	FROM
		users
	WHERE
		email = :email`

	var dbUsr dbUser
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbUsr); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return authenticate.User{}, fmt.Errorf("namedquerystruct: %w", authenticate.ErrNotFound)
		}
		return authenticate.User{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	return toFeatureUser(dbUsr), nil
}
