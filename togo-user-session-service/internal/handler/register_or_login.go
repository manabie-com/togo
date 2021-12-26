package handler

import (
	"context"
	"strings"
	v1 "togo-user-session-service/pkg/api/v1"

	"github.com/giahuyng98/togo/core-lib/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) RegisterOrLogin(ctx context.Context, req *v1.RegisterOrLoginRequest) (*v1.RegisterOrLoginResponse, error) {
	logger.For(ctx).Info("RegisterOrLogin start", zap.Any("req", req))

	if err := h.validateRegisterOrLoginRequest(ctx, req); err != nil {
		logger.For(ctx).Error("RegisterOrLogin validate failed", zap.Error(err))
		return nil, err
	}

	token, err := h.DB.RegisterOrLogin(ctx, req.Username, req.Password)

	if err != nil {
		logger.For(ctx).Error("RegisterOrLogin error", zap.Error(err))
		return nil, err
	}

	resp := &v1.RegisterOrLoginResponse{
		Token: token,
	}

	logger.For(ctx).Info("RegisterOrLogin end", zap.Any("resp", resp))
	return resp, nil
}

func (h Handler) validateRegisterOrLoginRequest(ctx context.Context, req *v1.RegisterOrLoginRequest) error {
	if len(strings.TrimSpace(req.Username)) == 0 {
		return status.Error(codes.InvalidArgument, "UserName")
	}
	if len(strings.TrimSpace(req.Password)) == 0 {
		return status.Error(codes.InvalidArgument, "Password")
	}
	return nil
}
