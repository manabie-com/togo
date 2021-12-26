package handler

import (
	"context"
	"togo-public-api/internal/service/togo_user_session_v1"
	v1 "togo-public-api/pkg/api/v1"

	"github.com/giahuyng98/togo/core-lib/logger"
	"go.uber.org/zap"
)

func (h Handler) RegisterOrLogin(ctx context.Context, req *v1.RegisterOrLoginRequest) (*v1.RegisterOrLoginResponse, error) {
	logger.For(ctx).Info("RegisterOrLogin start", zap.Any("req", req))

	registerOrLoginResp, err := h.TogoUserSessionService.RegisterOrLogin(ctx,
		&togo_user_session_v1.RegisterOrLoginRequest{
			Username: req.Username,
			Password: req.Password,
		},
	)

	if err != nil {
		logger.For(ctx).Error("RegisterOrLogin error", zap.Error(err))
		return nil, toPublicError(err)
	}

	resp := &v1.RegisterOrLoginResponse{
		Token: registerOrLoginResp.Token,
	}
	logger.For(ctx).Info("RegisterOrLogin end", zap.Any("resp", resp))
	return resp, nil
}
