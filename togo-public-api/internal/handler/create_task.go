package handler

import (
	"context"
	"togo-public-api/internal/auth"
	"togo-public-api/internal/converter"
	"togo-public-api/internal/service/togo_internal_v1"
	v1 "togo-public-api/pkg/api/v1"

	"github.com/giahuyng98/togo/core-lib/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h Handler) CreateTask(ctx context.Context, req *v1.CreateTaskRequest) (*v1.Task, error) {
	logger.For(ctx).Info("CreateTask start", zap.Any("req", req))

	taskResp, err := h.TogoInternalService.CreateTask(ctx,
		&togo_internal_v1.CreateTaskRequest{
			UserId:      auth.GetUserID(ctx),
			Title:       req.Title,
			Content:     req.Content,
			CreatedTime: timestamppb.Now(),
		})

	if err != nil {
		logger.For(ctx).Error("CreateTask error", zap.Error(err))
		return nil, toPublicError(err)
	}
	resp := converter.ToTask(taskResp)
	logger.For(ctx).Info("CreateTask end", zap.Any("resp", resp))
	return resp, nil
}
