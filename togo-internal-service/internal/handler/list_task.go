package handler

import (
	"context"
	"time"
	"togo-internal-service/internal/converter"
	"togo-internal-service/internal/model"
	v1 "togo-internal-service/pkg/api/v1"
	"github.com/giahuyng98/togo/core-lib/logger"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) ListTask(ctx context.Context, req *v1.ListTaskRequest) (*v1.ListTaskResponse, error) {
	logger.For(ctx).Info("ListTask start", zap.Any("req", req))

	var err error
	var offset, pageSize int
	if pageSize, offset, err = h.validateListTaskRequest(ctx, req); err != nil {
		return nil, err
	}

	var date time.Time
	if req.Date != nil {
		date = time.Date(int(req.Date.Year), time.Month(req.Date.Month), int(req.Date.Day), 0, 0, 0, 0, time.UTC)
	}

	tasks, err := h.Storage.ListTask(ctx, req.UserId, date, pageSize, offset)
	if err != nil {
		logger.For(ctx).Error("ListTask list error", zap.Error(err))
		return nil, toGRPCError(err)
	}
	logger.For(ctx).Info("ListTask list ok", zap.Any("tasks", tasks))

	resp := &v1.ListTaskResponse{
		Tasks: converter.TasksToDTO(tasks),
	}

	if len(tasks) < int(req.PageSize) {
		resp.NextPageToken = ""
	} else {
		resp.NextPageToken = model.PagingToToken(&model.Paging{
			Offset: offset + len(tasks),
		})
	}

	logger.For(ctx).Info("ListTask end", zap.Any("resp", resp))

	return resp, nil
}
func (h Handler) validateListTaskRequest(ctx context.Context, req *v1.ListTaskRequest) (pageSize int, offset int, err error) {
	pageSize = int(req.PageSize)
	offset = 0

	if pageSize > h.Config.MaxListTaskPageSize {
		pageSize = h.Config.MaxListTaskPageSize
	}

	if len(req.UserId) == 0 {
		err = status.Error(codes.InvalidArgument, "UserID")
		return
	}

	if req.PageSize <= 0 {
		err = status.Error(codes.InvalidArgument, "PageSize")
		return
	}

	paging, err := model.TokenToPaging(req.PageToken)
	if err != nil {
		err = status.Error(codes.InvalidArgument, "PageToken")
		return
	}
	offset = paging.Offset
	return
}
