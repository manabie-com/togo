package decoder

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/valonekowd/togo/usecase/request"
)

func SignIn(_ context.Context, req *http.Request) (interface{}, error) {
	r := request.SignIn{}

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func SignUp(_ context.Context, req *http.Request) (interface{}, error) {
	r := request.SignUp{}

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
