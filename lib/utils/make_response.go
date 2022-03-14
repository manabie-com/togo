package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"Message,omitempty"`
	Data    interface{} `json:"Data,omitempty"`
	Status  int         `json:"Status,omitempty"`
}

func WriteSuccess(c *fiber.Ctx, v interface{}) error {
	res := Response{
		Message: "Success",
		Data:    v,
		Status:  200,
	}

	return c.JSON(res)
}

func WriteSuccessEmptyContent(c *fiber.Ctx) error {
	res := Response{
		Message: "Success",
		Status:  fiber.StatusOK,
	}

	return c.JSON(res)
}

func WriteError(c *fiber.Ctx, status int, err error) error {
	res := Response{
		Message: "Error",
		Data:    err.Error(),
		Status:  status,
	}

	return c.Status(status).JSON(res)
}
