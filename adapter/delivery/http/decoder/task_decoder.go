package decoder

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/valonekowd/togo/usecase/request"
	"github.com/valonekowd/togo/util/helper"
)

func GetTasks(_ context.Context, req *http.Request) (interface{}, error) {
	createdDateStr := helper.QueryParam(req, "created_date")

	createdDate := time.Now()
	var err error
	if createdDateStr != "" {
		createdDate, err = helper.RequestParam(createdDateStr).Time()
	}

	if err != nil {
		return nil, err
	}

	return request.GetTasks{CreatedDate: createdDate.UTC()}, nil
}

func CreateTask(_ context.Context, req *http.Request) (interface{}, error) {
	r := request.CreateTask{}

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
