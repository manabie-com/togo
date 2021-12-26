package handler

import (
	"context"
	"togo-public-api/internal/auth"
	"togo-public-api/internal/converter"
	"togo-public-api/internal/service/togo_internal_v1"
	v1 "togo-public-api/pkg/api/v1"

	"github.com/giahuyng98/togo/core-lib/logger"
	"go.uber.org/zap"
)

func (h Handler) ListTask(ctx context.Context, req *v1.ListTaskRequest) (*v1.ListTaskResponse, error) {
	logger.For(ctx).Info("ListTask start", zap.Any("req", req))

	listTaskResp, err := h.TogoInternalService.ListTask(ctx,
		&togo_internal_v1.ListTaskRequest{
			UserId:    auth.GetUserID(ctx),
			Date:      req.Date,
			PageSize:  req.PageSize,
			PageToken: req.PageToken,
		})

	if err != nil {
		logger.For(ctx).Error("ListTask error", zap.Error(err))
		return nil, toPublicError(err)
	}

	tasksResp := converter.ToTasks(listTaskResp.Tasks)
	resp := &v1.ListTaskResponse{
		Tasks:         tasksResp,
		NextPageToken: listTaskResp.NextPageToken,
	}

	logger.For(ctx).Info("ListTask end", zap.Any("resp", resp))
	return resp, nil
}
