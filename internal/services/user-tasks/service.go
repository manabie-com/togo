package user_tasks

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/manabie-com/togo/internal/pkg/util"
	"github.com/manabie-com/togo/internal/services/users"

	"github.com/google/uuid"
	"github.com/looplab/eventhorizon"
)

type Service interface {
	CreateTask(ctx context.Context, userID uuid.UUID, content string) error
}

type service struct {
	commandBus     eventhorizon.CommandHandler
	userConfigRepo users.UserConfigRepo
}

func NewService(
	commandBus eventhorizon.CommandHandler,
	userConfigRepo users.UserConfigRepo,
) *service {
	return &service{
		commandBus:     commandBus,
		userConfigRepo: userConfigRepo,
	}
}

func (s *service) CreateTask(ctx context.Context, userID uuid.UUID, content string) error {
	usrCfg, err := s.userConfigRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if usrCfg == nil {
		return errors.New("User is not configed")
	}

	err = s.commandBus.HandleCommand(ctx, &CreateTask{
		UserID:    userID,
		Content:   content,
		TaskLimit: usrCfg.TaskLimit,
	})
	if err != nil {
		logrus.WithError(err).Errorf("CreateTask, user_id (%s), content (%s)", userID.String(), content)
		return util.UnwrapAggError(err)
	}

	return nil
}
