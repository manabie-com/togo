package user

import (
	"net/http"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/util"

	"github.com/labstack/echo/v4"
)

// Controller interface
type Controller interface {
	Register(c echo.Context) error
}

// NewUserController func
func NewUserController(userService Service) (Controller, error) {
	return &controller{
		userService: userService,
	}, nil
}

type controller struct {
	userService Service
}

func (controller *controller) Register(c echo.Context) error {
	registerDTO := new(dto.RegisterDTO)
	if err := c.Bind(registerDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if err := c.Validate(registerDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	isEmailExisted, err := controller.userService.CheckEmailExist(registerDTO.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if isEmailExisted {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Email is existed",
		})
	}

	password := util.HashPassword(registerDTO.Password)
	println("password:", password)

	user, err := controller.userService.Create(registerDTO.Email, password, config.Cfg.MaxTodoDefault)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user.ToObject())
}
