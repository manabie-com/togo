package response

import (
	"github.com/manabie-com/togo/internal/domain"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetStatusCode(err error) int {
	switch err {
	case domain.ErrorMaximumTaskPerDay:
		return http.StatusBadRequest
	case domain.UserNotFound:
		return http.StatusNotFound
	case domain.Unauthorized:
		return http.StatusUnauthorized
	case domain.WrongPassword:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
