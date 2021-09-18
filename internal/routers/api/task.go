package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/manabie-com/togo/internal/models"
	endpoints "github.com/manabie-com/togo/internal/routers/endpoints"
)

func AddTask(ctx context.Context, request *http.Request) (interface{}, error) {
	req := models.Task{}
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func ListTasks(ctx context.Context, request *http.Request) (interface{}, error) {
	createdDate := request.URL.Query().Get("created_date")
	layout := "2006-01-02"
	_createdDate, err := time.Parse(layout, createdDate)
	if err != nil {
		return nil, err
	}
	return _createdDate, nil
}

func TaskRouter(
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
				AddTask,
				endpoints.Task.AddTask,
			)
		} else if httpRequest.Method == http.MethodGet {
			resData, err = RouteHandle(
				ctx,
				httpRequest,
				ListTasks,
				endpoints.Task.ListTasks,
			)
		} else {
			err = errors.New("not found")
		}
	default:
		err = errors.New("not found")
	}
	return resData, err
}
