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
func (U UserCreateAction) Execute(Username string, Password string) (user Users, err error) {
	count, err := U.Db.Model(new(Users)).Where("username = ?", Username).Count()
	if err != nil {
		return Users{}, errors.Wrapf(err, "Check existed user error")
	}

	if count == 1 {
		return Users{}, errors.Errorf("Current user existed")
	}

	hashPassword, err := security.Hash(Password)
	if err != nil {
		return Users{}, errors.Wrapf(err, "While creating hash password error")
	}

	user = Users{Username: Username, Password: string(hashPassword), MaxTodo: 5}

	_, err = U.Db.Model(&user).Insert()
	if err != nil {
		return Users{}, errors.Wrapf(err, "Create a user error")
	}

	return user, nil
}
