package usecase

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/api/dictionary"
	"github.com/manabie-com/togo/internal/api/user/storages"
	"github.com/manabie-com/togo/internal/api/utils"
	"github.com/manabie-com/togo/internal/pkg/logger"
)

type User struct {
	Store storages.Store
}

func (s *User) IsValidate(ctx context.Context, userID, password string) (bool, error) {
	user, err := s.Store.Get(ctx, userID)
	if err != nil {
		logger.MBError(ctx, err)
		return false, errors.New(dictionary.FailedToGetUser)
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return false, nil
	}

	return true, nil
}
