package handler

import (
	"context"
	"togo-public-api/internal/converter"
	"togo-public-api/internal/service/togo_internal_v1"
	v1 "togo-public-api/pkg/api/v1"

	"github.com/giahuyng98/togo/core-lib/logger"
	"go.uber.org/zap"
)

func (h Handler) GetTask(ctx context.Context, req *v1.GetTaskRequest) (*v1.Task, error) {
	logger.For(ctx).Info("GetTask start", zap.Any("req", req))

	taskResp, err := h.TogoInternalService.GetTask(ctx,
		&togo_internal_v1.GetTaskRequest{
			Id: req.Id,
		})
	if err != nil {
		logger.For(ctx).Error("GetTask error", zap.Error(err))
		return nil, toPublicError(err)
	}
	resp := converter.ToTask(taskResp)
	logger.For(ctx).Info("GetTask end", zap.Any("resp", resp))
	return resp, nil
}
