package kit

import (
	"context"
	"encoding/json"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/dtos"
	"net/http"
)

func defaultEncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func defaultEncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	errResponse := &dtos.ErrorResponse{
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(errResponse)
}

func codeFrom(err error) int {
	switch err.Error() {
	case define.AccountNotAuthorized:
	case define.Unauthenticated:
		return http.StatusUnauthorized
	case define.AccountNotExist:
		return http.StatusNotFound
	case define.FailedValidation:
		return http.StatusBadRequest
	case define.Unknown:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
