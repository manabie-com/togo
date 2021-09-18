package endpoints

import (
	"context"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/services"
)

type UserEndpoint struct {
	GetAuthToken Endpoint
	ValidToken   Endpoint
}

func MakeGetAuthTokenEndpoint(s *services.Service) Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		user, _ := request.(models.User)
		res, err := s.UserService.GetAuthToken(ctx, user)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
