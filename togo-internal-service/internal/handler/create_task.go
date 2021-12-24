package handler

import (
	"context"
	"togo-internal-service/internal/converter"
	"togo-internal-service/internal/model"
	v1 "togo-internal-service/pkg/api/v1"
	"github.com/giahuyng98/togo/core-lib/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (*v1.Task, error) {
	logger.For(ctx).Info("CreateTask start", zap.Any("req", req))

	if err := h.validateCreateTaskRequest(ctx, req); err != nil {
		logger.For(ctx).Error("CreateTask validate error", zap.Error(err))
		return nil, err
	}

	task, err := h.Storage.CreateTask(ctx, &model.Task{
		Title:       req.Title,
		Content:     req.Content,
		UserID:      req.UserId,
		CreatedTime: req.CreatedTime.AsTime(),
	})

	if err != nil {
		logger.For(ctx).Error("CreateTask save error", zap.Error(err))
		return nil, toGRPCError(err)
	}

	resp := converter.TaskToDTO(task)

	logger.For(ctx).Info("CreateTask end", zap.Any("resp", resp))
	return resp, nil
}

func (h Handler) validateCreateTaskRequest(ctx context.Context, req *v1.CreateTaskRequest) error {
	if len(req.Title) == 0 {
		return status.Error(codes.InvalidArgument, "Title")
	}
	if len(req.UserId) == 0 {
		return status.Error(codes.InvalidArgument, "UserID")
	}

	if req.CreatedTime == nil {
		return status.Error(codes.InvalidArgument, "CreatedTime")
	}
	return nil
}
