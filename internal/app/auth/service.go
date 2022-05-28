package auth

import (
	"context"
	"fmt"

	"github.com/dinhquockhanh/togo/internal/app/user"
	"github.com/dinhquockhanh/togo/internal/pkg/errors"
	"github.com/dinhquockhanh/togo/internal/pkg/util"
)

type (
	service struct {
		userSvc user.Service
	}
)

func NewService(us user.Service) Service {
	return &service{
		userSvc: us,
	}
}

func (s *service) Auth(ctx context.Context, username, password string) (*user.User, error) {
	errInvalid := &errors.Error{
		Code:    400,
		Message: "wrong user name or password",
	}

	usr, err := s.userSvc.GetByUserName(ctx, &user.GetUserByUserNameReq{UserName: username})
	if err != nil {
		if errors.IsSQLNotFound(err) {
			return nil, errInvalid
		}
		return nil, fmt.Errorf("auth: %w", err)
	}

	if err := util.CheckPassword(password, usr.HashedPassword); err != nil {

		return nil, errInvalid
	}

	return usr, nil
}
