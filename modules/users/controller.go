package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manabie-com/togo/common/response"
)

type UsersController struct {
	Srv Service
}

func (r UsersController) Login(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	password := c.Query("password")
	if userId == "" || password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "userId and password is required!")
	}

	token, err := r.Srv.Login(userId, password)
	if err != nil {
		resErr := response.ErrorResponse{
			Status:  fiber.StatusUnauthorized,
			Error:   err.Error(),
			Message: "Error",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(resErr)
	}

	if len(token) < 1 {
		resErr := response.ErrorResponse{
			Status:  fiber.StatusUnauthorized,
			Error:   err.Error(),
			Message: "Login Failed!",
		}
		return c.Status(fiber.StatusUnauthorized).JSON(resErr)
	}

	res := response.SuccessResponse{
		Status:  fiber.StatusOK,
		Data:    token,
		Message: "OK",
	}
	return c.JSON(res)
}
