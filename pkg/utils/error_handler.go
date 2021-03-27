package utils

import (
	"github.com/manabie-com/togo/domains"
	"github.com/manabie-com/togo/pkg/core/servehttp"
	"github.com/manabie-com/togo/usecases"
	"net/http"
)

func ResponseErrorHandler(w http.ResponseWriter, err error) {
	switch true {
	case err == domains.ErrorUnAuthorized:
		servehttp.ResponseErrorJSON(w, http.StatusUnauthorized, err.Error())
	case err == usecases.ErrorUserNotFound:
		servehttp.ResponseErrorJSON(w, http.StatusUnauthorized, err.Error())
	case err == usecases.ErrorReachedLimitCreateTaskPerDay:
		servehttp.ResponseErrorJSON(w, http.StatusTooManyRequests, err.Error())
	case err == domains.ErrorNotFound:
		servehttp.ResponseErrorJSON(w, http.StatusNotFound, err.Error())
	default:
		servehttp.ResponseErrorJSON(w, http.StatusInternalServerError, err.Error())
	}
}
