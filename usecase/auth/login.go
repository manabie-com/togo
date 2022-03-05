package auth

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/util"
	"github.com/khangjig/togo/util/jwt"
	"github.com/khangjig/togo/util/myerror"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UseCase) Login(ctx context.Context, req *LoginRequest) (*ResponseWrapper, error) {
	req.Email = strings.TrimSpace(req.Email)
	if req.Email == "" || req.Password == "" {
		return nil, myerror.ErrInvalidEmailPassword()
	}

	isEmail, err := util.IsEmail(req.Email)
	if err != nil {
		return nil, myerror.ErrRegexp(err)
	}

	if !isEmail {
		return nil, myerror.ErrInvalidEmailPassword()
	}

	myUser, err := u.UserRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrInvalidEmailPassword()
		}

		return nil, myerror.ErrGet(err)
	}

	if !util.ComparePassword(req.Password, myUser.Password) {
		return nil, myerror.ErrInvalidEmailPassword()
	}

	token, err := jwt.EncodeToken(myUser.ID, u.Config.TokenSecretKey)
	if err != nil {
		return nil, myerror.ErrEncodeToken(err)
	}

	return &ResponseWrapper{
		Token: token,
		User:  myUser,
	}, nil
}
