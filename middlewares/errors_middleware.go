package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/SVincentTran/togo/errors"
)

// Err handler middleware
type ErrHandler func(http.ResponseWriter, *http.Request) error

func (fn ErrHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			xerr := errors.GetError(errors.InternalErrorContext, errors.InteralErrorMessage, errors.UnexpectedError)
			w.WriteHeader(xerr.Code)
			_ = json.NewEncoder(w).Encode(xerr)
			return
		}
	}()
	if err := fn(w, r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		xerr, ok := err.(*errors.CustomError)
		if !ok {
			xerr = errors.GetError(errors.InternalErrorContext, errors.InteralErrorMessage, errors.UnexpectedError)
		}

		w.WriteHeader(xerr.Code)
		_ = json.NewEncoder(w).Encode(xerr)
		return
	}
}
