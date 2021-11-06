package users

import (
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/quochungphp/go-test-assignment/src/pkgs/security"
)

// UserCreateAction ...
type UserCreateAction struct {
	Db *pg.DB
}

// UserCreateAction ...
func (U UserCreateAction) Execute(Username string, Password string) (user User, err error) {
	count, err := U.Db.Model(new(User)).Where("username = ?", Username).Count()
	if err != nil {
		return User{}, errors.Wrapf(err, "Check existed user error")
	}

	if count == 1 {
		return User{}, errors.Errorf("Current user existed")
	}

	hashPassword, err := security.Hash(Password)
	if err != nil {
		return User{}, errors.Wrapf(err, "While creating hash password error")
	}

	user = User{Username: Username, Password: string(hashPassword), MaxTodo: 5}

	_, err = U.Db.Model(&user).Insert()
	if err != nil {
		return User{}, errors.Wrapf(err, "Create a user error")
	}

	return user, nil
}
