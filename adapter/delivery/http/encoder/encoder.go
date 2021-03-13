package encoder

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	if f, ok := resp.(endpoint.Failer); ok && f.Failed() != nil {
		EncodeError(ctx, f.Failed(), w)
		return nil
	}

	return kithttp.EncodeJSONResponse(ctx, w, resp)
}

type errorWrapper struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if headerer, ok := err.(kithttp.Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(kithttp.StatusCoder); ok {
		code = sc.StatusCode()
	}

	w.WriteHeader(code)

	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
