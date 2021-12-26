package handler

import (
	"context"
	"togo-internal-service/internal/converter"
	"togo-internal-service/pkg/api/v1"
	"github.com/giahuyng98/togo/core-lib/logger"

	"go.uber.org/zap"
)

func (h Handler) GetTask(ctx context.Context, req *togo_internal_v1.GetTaskRequest) (*togo_internal_v1.Task, error) {
	logger.For(ctx).Info("GetTask start", zap.Any("req", req))

	task, err := h.Storage.GetTask(ctx, req.Id)

	if err != nil {
		logger.For(ctx).Error("GetTask get error", zap.Error(err))
		return nil, toGRPCError(err)
	}

	resp := converter.TaskToDTO(task)
	logger.For(ctx).Info("GetTask end", zap.Any("req", req))

	return resp, nil
}
