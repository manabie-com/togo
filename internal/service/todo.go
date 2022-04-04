package service

import (
	"context"
	"errors"
	"time"

	"github.com/vchitai/l"
	"github.com/vchitai/togo/internal/mapper"
	"github.com/vchitai/togo/internal/models"
	"github.com/vchitai/togo/internal/store"
	"github.com/vchitai/togo/internal/utils"
	"github.com/vchitai/togo/pb"
)

var now = time.Now

func (s *serverImpl) AddToDoList(ctx context.Context, req *pb.AddToDoListRequest) (resp *pb.AddToDoListResponse, respErr error) {
	var (
		userID = req.GetUserId()
		today  = utils.RoundDate(
			now().Add(7 * time.Hour), // TODO: allow config time zone later
		)
	)
	// get user config if exist
	userCfg, err := s.toDoStore.GetConfig(ctx, userID)
	if errors.Is(err, store.ErrNotFound) {
		// if not exist, use default config
		userCfg = s.getDefaultUserConfig()
	} else if err != nil {
		// if other error occurred, return
		ll.Error("Get user config failed", l.Error(err))
		return nil, errInternal
	}

	// get user used quota
	userUsedCount, err := s.toDoStore.GetUsedCount(ctx, userID, today)
	if err != nil {
		// if error occurred, return
		ll.Error("Get user used count failed", l.Error(err))
		return nil, errInternal
	}
	// if user use over the limited
	if userUsedCount >= userCfg.Limited {
		return nil, errDailyQuotaReached
	}
	// if after added the list, user use over the limited
	var addingAmount = int64(len(req.GetEntry()))
	if userUsedCount+addingAmount > userCfg.Limited {
		return nil, errDailyQuotaExceed
	}

	// acquire the lock by commit the adding amount
	userUsedCount, err = s.toDoStore.IncreaseUsedCount(ctx, userID, today, addingAmount)
	if err != nil {
		ll.Error("Increase user used count failed", l.Error(err))
		return nil, errInternal
	}
	defer func() {
		if respErr != nil {
			userUsedCount, err = s.toDoStore.DecreaseUsedCount(ctx, userID, today, addingAmount)
			if err != nil {
				ll.Error("Decrease user used count failed", l.Error(err))
			}
		}
	}()
	if userUsedCount > userCfg.Limited {
		return nil, errDailyQuotaExceed
	}

	// record the list
	var todoModel = mapper.Proto2ModelToDoEntryList(req.GetEntry())
	for _, todo := range todoModel {
		todo.UserID = userID
	}
	if err := s.toDoStore.Record(ctx, todoModel); err != nil {
		// if record failed, release the lock
		ll.Error("Record to do list failed", l.Error(err))
		return nil, errInternal
	}

	return &pb.AddToDoListResponse{
		Message: "ok",
	}, nil
}

func (s *serverImpl) getDefaultUserConfig() *models.ToDoConfig {
	return &models.ToDoConfig{
		Limited: s.cfg.ToDoListAddLimitedPerDay,
	}
}
