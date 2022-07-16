package users

import (
	"net/http"

	"manabie/todo/models"
	"manabie/todo/pkg/utils"
	"manabie/todo/service/user"

	"github.com/labstack/echo/v4"
)

type handler struct {
	User user.UserService
}

func NewUserHandler(e *echo.Echo, us user.UserService) {
	h := &handler{
		User: us,
	}
	e.GET("/users", h.Index)
}

func (h *handler) Index(c echo.Context) error {
	users, err := h.User.Index(c.Request().Context())
	if err != nil {
		return utils.ResponseFailure(c, http.StatusInternalServerError, err)
	}

	return utils.ResponseSuccess(c, models.UserIndexResponse{
		Users: users,
	})
}
