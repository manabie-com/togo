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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverImpl) AddToDoList(ctx context.Context, req *pb.AddToDoListRequest) (*pb.AddToDoListResponse, error) {
	var (
		userID = req.GetUserId()
		today  = utils.RoundDate(
			time.Now().Add(7 * time.Hour), // TODO: allow config time zone later
		)
	)
	userCfg, err := s.toDoStore.GetConfig(ctx, userID)
	if errors.Is(err, store.ErrNotFound) {
		userCfg = s.getDefaultUserConfig()
	} else if err != nil {
		ll.Error("Get user config failed", l.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	userUsedCount, err := s.toDoStore.GetUsedCount(ctx, userID, today)
	if err != nil {
		ll.Error("Get user used count failed", l.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	if userUsedCount >= userCfg.Limited {
		return nil, status.Error(codes.InvalidArgument, "you have reached daily limit")
	}
	var addingAmount = int64(len(req.GetEntry()))
	if userUsedCount+addingAmount > userCfg.Limited {
		return nil, status.Error(codes.InvalidArgument, "you will exceed daily limit adding this list")
	}
	userUsedCount, err = s.toDoStore.IncreaseUsedCount(ctx, userID, today, addingAmount)
	if err != nil {
		ll.Error("Increase user used count failed", l.Error(err))
		return nil, status.Error(codes.Internal, "internal server error")
	}
	var todoModel = mapper.Proto2ModelToDoEntryList(req.GetEntry())
	for _, todo := range todoModel {
		todo.UserID = userID
	}
	if err := s.toDoStore.Record(ctx, todoModel); err != nil {
		ll.Error("Record to do list failed", l.Error(err))
		userUsedCount, err = s.toDoStore.DecreaseUsedCount(ctx, userID, today, addingAmount)
		if err != nil {
			ll.Error("Decrease user used count failed", l.Error(err))
		}
		return nil, status.Error(codes.Internal, "internal server error")
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
