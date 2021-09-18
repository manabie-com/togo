package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type StatusCoder interface {
	StatusCode() int
}

func EncodeJSONResponse(ctx context.Context, resW http.ResponseWriter, res interface{}, err error) error {
	resW.Header().Set("Content-Type", "application/json; charset=utf-8")
	code := http.StatusOK
	if err != nil {
		code = http.StatusInternalServerError
		if sc, ok := err.(StatusCoder); ok {
			code = sc.StatusCode()
		}
		resW.WriteHeader(code)
		return json.NewEncoder(resW).Encode(map[string]interface{}{
			"message": err.Error(),
		})
	}

	if sc, ok := res.(StatusCoder); ok {
		code = sc.StatusCode()
	}

	resW.WriteHeader(code)
	if code == http.StatusNoContent {
		return nil
	}
	return json.NewEncoder(resW).Encode(res)
}
