package rest

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/laghodessa/togo/app"
	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/infra/postgres"
)

func NewFiber(db *sql.DB) *fiber.App {
	server := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	api := server.Group("/api/v1")

	userRepo := postgres.NewTodoUserRepo(db)
	taskRepo := postgres.NewTodoTaskRepo(db)

	todoUC := &app.TodoUsecase{
		TaskRepo: taskRepo,
		UserRepo: userRepo,
	}

	RegisterTasks(api, todoUC)
	RegisterUsers(api, todoUC)

	return server
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if errors.Is(err, ErrMalformedRequestBody) {
		return c.
			Status(ErrMalformedRequestBody.Code).
			JSON(ErrorResponse{
				Code:    "malformed_request_body",
				Message: ErrMalformedRequestBody.Message,
			})
	}

	status := httpStatusFromError(err)
	resp := toErrorResponse(err)
	return c.Status(status).JSON(resp)
}

func toErrorResponse(err error) ErrorResponse {
	var domainErr domain.Error
	if errors.As(err, &domainErr) {
		return ErrorResponse{
			Code:    domainErr.Code,
			Message: domainErr.Message,
		}
	}

	return ErrorResponse{
		Code:    domain.CodeInternal,
		Message: "a server error has occured",
	}
}

// httpStatusFromError translate domain error to http status.
// It returns 500 on unknown error.
func httpStatusFromError(err error) int {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound

	case errors.As(err, &domain.Error{}):
		return http.StatusUnprocessableEntity

	default:
		return http.StatusInternalServerError
	}
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	ErrMalformedRequestBody = fiber.NewError(http.StatusBadRequest, "malformed request body")
)
