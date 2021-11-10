package domain

import (
	"context"
	"github.com/manabie-com/togo/internal/helper"
	"github.com/manabie-com/togo/pkg/token"
	"time"
)

type authUseCase struct {
	authRepository UserRepository
	tokenMaker     token.Token
	tokenConfig    TokenConfig
}

type TokenConfig struct {
	tokenDuration time.Duration
}

func NewTokenConfig(duration time.Duration) TokenConfig {
	return TokenConfig{
		tokenDuration: duration,
	}
}

func NewAuthUseCase(authRepository UserRepository, tokenMaker token.Token, tokenConfig TokenConfig) AuthUseCase {
	return authUseCase{
		authRepository: authRepository,
		tokenMaker:     tokenMaker,
		tokenConfig:    tokenConfig,
	}
}

func (a authUseCase) SignIn(ctx context.Context, username string, password string) (string, error) {
	user, err := a.authRepository.GetUser(ctx, username)
	if err != nil {
		return "", err
	}
	err = helper.CheckPassword(password, user.HashedPassword)
	if err != nil {
		return "", WrongPassword
	}
	accessToken, err := a.tokenMaker.CreateToken(username, a.tokenConfig.tokenDuration)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
