package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	d "github.com/manabie-com/togo/internal/todo/domain"
)

type PGUserRepository struct {
	PGRepository
}

func NewPGUserRepository(dbConn *sqlx.DB) *PGUserRepository {
	return &PGUserRepository{PGRepository{dbConn}}
}

func (t *PGUserRepository) GetByCredentials(username string, password string) (*d.User, error) {
	user := d.User{}
	err := t.DBConn.Get(&user,
		"SELECT * FROM users WHERE username = $1 AND password = crypt($2, password)",
		username, password)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, nil
}

func (t *PGUserRepository) GetByID(id int) (*d.User, error) {
	user := d.User{}
	err := t.DBConn.Get(&user,
		"SELECT * FROM users WHERE id = $1",
		id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, nil
}
