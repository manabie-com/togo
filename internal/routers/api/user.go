package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/manabie-com/togo/internal/models"
	endpoints "github.com/manabie-com/togo/internal/routers/endpoints"
)

func GetAuthToken(ctx context.Context, request *http.Request) (interface{}, error) {
	var req = models.User{}
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func UserRouter(
	ctx context.Context,
	httpRequest *http.Request,
	path string,
	endpoints *endpoints.Endpoints,
) (resData interface{}, err error) {
	switch path {
	case "":
		if httpRequest.Method == http.MethodPost {
			resData, err = RouteHandle(
				ctx,
				httpRequest,
				GetAuthToken,
				endpoints.User.GetAuthToken,
			)
		} else {
			err = errors.New("not found")
		}
	default:
		err = errors.New("not found")
	}
	return resData, err
}
