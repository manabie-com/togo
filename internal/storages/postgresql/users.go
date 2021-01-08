package postgresql

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
	"upper.io/db.v3/lib/sqlbuilder"
)

const usersTable string = "users"

type UserStore struct {
	db sqlbuilder.Database
}

func NewUserStore(db sqlbuilder.Database) UserStore {
	return UserStore{
		db: db,
	}
}

func (s UserStore) Get(ctx context.Context, userID sql.NullString) (*model.User, error) {
	var u model.User
	err := s.db.WithContext(ctx).SelectFrom(usersTable).Where("id = ?", userID).One(&u)
	if err != nil {
		return nil, model.NewError(model.ErrGetUser, err.Error())
	}

	return &u, nil
}

func (s UserStore) Create(ctx context.Context, u *model.User) (*model.User, error) {
	if err := u.IsValid(); err != nil{
		return nil, err
	}

	_, err := s.db.WithContext(ctx).InsertInto(usersTable).Values(u).Exec()
	if err != nil {
		return nil, model.NewError(model.ErrCreateUser, err.Error())
	}

	return u, nil
}