package formatter

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/infrastructure/auth"
	"github.com/valonekowd/togo/usecase/interfaces"
	"github.com/valonekowd/togo/usecase/response"
)

type userFormatter struct {
	authCfg auth.Config
}

var _ interfaces.UserPresenter = userFormatter{}

func NewUserFormatter(authCfg auth.Config) interfaces.UserPresenter {
	return userFormatter{
		authCfg: authCfg,
	}
}

func (f userFormatter) generateAccessToken(ctx context.Context, u *entity.User) (string, error) {
	token := jwt.NewWithClaims(f.authCfg.JWT.SigningMethod, jwt.MapClaims{
		"iss":     f.authCfg.JWT.Issuer,
		"exp":     time.Now().UTC().Add(time.Hour * 8).Unix(),
		"user_id": u.ID,
	})

	key, err := f.authCfg.JWT.Keyfunc(token)
	if err != nil {
		return "", fmt.Errorf("identifying secret key using jwt.Keyfunc: %w", err)
	}

	t, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("signing token: %w", err)
	}

	return t, nil
}

func (f userFormatter) SignUp(ctx context.Context, u *entity.User) (*response.SignUp, error) {
	token, err := f.generateAccessToken(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("generating access token: %w", err)
	}

	return &response.SignUp{
		Data: &response.SignUpPayload{
			ID:          u.ID,
			Email:       u.Email,
			MaxTodo:     u.MaxTodo,
			CreatedAt:   u.CreatedAt.Format("2020-02-10 20:00:00 +0700"),
			AccessToken: token,
		},
	}, nil
}

func (f userFormatter) SignIn(ctx context.Context, u *entity.User) (*response.SignIn, error) {
	token, err := f.generateAccessToken(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("generating access token: %w", err)
	}

	return &response.SignIn{
		Data: &response.SignInPayload{
			AccessToken: token,
		},
	}, nil
}
