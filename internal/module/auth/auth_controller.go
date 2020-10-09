package auth

import (
	"net/http"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/module/user"
	"github.com/manabie-com/togo/internal/util"

	"github.com/labstack/echo/v4"
)

// Controller interface
type Controller interface {
	Login(c echo.Context) error
}

type controller struct {
	userService user.Service
}

// NewAuthController func
func NewAuthController(userService user.Service) (Controller, error) {
	return &controller{
		userService: userService,
	}, nil
}

func (controller *controller) Login(c echo.Context) error {
	loginDTO := new(dto.LoginDTO)
	if err := c.Bind(loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := c.Validate(loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	user, err := controller.userService.Login(loginDTO)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}
	token, err := util.CreateAuthTokenPair(user.Email, user.ID, config.Cfg.JwtKey, config.Cfg.JwtExp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  user.ToObject(),
		"token": token,
	})
}
