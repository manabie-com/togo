package http

import (
	"errors"
	"net/http"
	"togo/internal/domain"

	"github.com/labstack/echo/v4"
)

var errorStatusMap = map[error]int{
	// Auth error statuses
	domain.ErrCredentialInvalid:  http.StatusBadRequest,
	domain.ErrDuplicatedUsername: http.StatusBadRequest,
	domain.ErrLoginFailed:        http.StatusServiceUnavailable,
	// User error statuses
	domain.ErrUserNotFound: http.StatusBadRequest,
	// Task error satuses
	domain.ErrTaskLimitExceed: http.StatusTooManyRequests,
	domain.ErrTaskNotFound:    http.StatusBadRequest,
}

func unwraps(err error) error {
	if inner := errors.Unwrap(err); inner != nil {
		return unwraps(inner)
	}
	return err
}

func httpErrorResolver(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		coreError := unwraps(err)
		if expectedStatus := errorStatusMap[coreError]; expectedStatus != 0 {
			return echo.NewHTTPError(expectedStatus, coreError.Error())
		}
		return err
	}
}
