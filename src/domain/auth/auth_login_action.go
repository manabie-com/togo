package auth

import (
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/quochungphp/go-test-assignment/src/domain/users"
	"github.com/quochungphp/go-test-assignment/src/pkgs/security"
	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
	"golang.org/x/crypto/bcrypt"
)

// AuthLoginAction ...
type AuthLoginAction struct {
	Db *pg.DB
}

// Execute ...
func (Auth AuthLoginAction) Execute(Username string, Password string) (tokenDetail token.TokenDetail, err error) {
	user := users.Users{}

	err = Auth.Db.Model(&user).Where("username = ?", Username).Select()
	if err != nil {
		return token.TokenDetail{}, errors.Wrapf(err, "Unauthorized")
	}

	err = security.VerifyPassword(user.Password, Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return token.TokenDetail{}, errors.Wrapf(err, "Unauthorized")
	}

	tokenDetail, err = token.CreateToken(user.ID, user.MaxTodo)
	if err != nil {
		return token.TokenDetail{}, errors.Wrapf(err, "Generate token error")
	}
	return tokenDetail, nil
}
