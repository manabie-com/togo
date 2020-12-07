package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/usecase/auth"
	"github.com/HoangVyDuong/togo/pkg/define"
	authDTO "github.com/HoangVyDuong/togo/pkg/dtos/auth"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"github.com/HoangVyDuong/togo/pkg/utils"
	"github.com/spf13/viper"
	"strconv"
)

type Handler interface {
	Auth(ctx context.Context, request authDTO.AuthUserRequest) (response authDTO.AuthUserResponse, err error)
}

type authHandler struct {
	authService auth.Service
}

func NewHandler(authService auth.Service) Handler{
	return &authHandler{authService}
}

func (ah *authHandler) Auth(ctx context.Context, request authDTO.AuthUserRequest) (response authDTO.AuthUserResponse, err error) {
	logger.Debugf("Start Auth User: %s", request.Username)
	if request.Username == "" || request.Password == "" {
		return authDTO.AuthUserResponse{}, define.FailedValidation
	}

	userID, err := ah.authService.Auth(ctx, request.Username, request.Password)
	if err != nil {
		if err == define.AccountNotExist || err == define.AccountNotAuthorized {
			return authDTO.AuthUserResponse{}, err
		}
		return authDTO.AuthUserResponse{}, define.Unknown
	}

	token, err := utils.CreateToken(strconv.FormatUint(userID, 10), viper.GetString("jwt.key"))
	if err != nil {
		logger.Errorf("[AuthHandler][Auth] create token error: %s", token)
		return authDTO.AuthUserResponse{}, define.Unknown
	}

	logger.Debugf("Auth user successfully %s", request.Username)
	return authDTO.AuthUserResponse{
		Token: token,
	}, nil
}

