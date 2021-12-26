package handler

import (
	"context"
	v1 "togo-user-session-service/pkg/api/v1"

	"github.com/giahuyng98/togo/core-lib/logger"
	"go.uber.org/zap"
)

func (h Handler) VerifyToken(ctx context.Context, req *v1.VerifyTokenRequest) (*v1.VerifyTokenResponse, error) {
	logger.For(ctx).Info("VerifyToken start", zap.Any("req", req))

	user, err := h.DB.VerifyToken(ctx, req.Token)

	if err != nil {
		logger.For(ctx).Error("VerifyToken error", zap.Error(err))
		return nil, err
	}

	resp := &v1.VerifyTokenResponse{
		UserId:   user.UserID,
		Username: user.UserName,
	}

	logger.For(ctx).Info("VerifyToken end", zap.Any("resp", resp))
	return resp, nil
}
