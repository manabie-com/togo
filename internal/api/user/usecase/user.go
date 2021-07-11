package usecase

import (
	"context"
	"fmt"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/api/user/storages"
	"github.com/manabie-com/togo/internal/api/utils"
	"github.com/manabie-com/togo/internal/pkg/logger"
)

type User struct {
	Cfg   *config.Config
	Store storages.Store
}

func (s *User) IsValidate(ctx context.Context, userID, password string) (bool, error) {
	user, err := s.Store.ValidateUser(ctx, userID)
	if err != nil {
		logger.MBError(ctx, err)
		return false, fmt.Errorf("failed to get user")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return false, nil
	}

	return true, nil
}
