package tasks

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/manabie-com/togo/common/response"
	"github.com/manabie-com/togo/modules/auth"
)

type TasksController struct {
	Srv Service
}

func (r TasksController) Create(c *fiber.Ctx) error {
	var req Tasks
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return err
	}

	result, err := r.Srv.Create(req, c.Locals("userAuth").(auth.UserAuth))
	if err != nil {
		resErr := response.ErrorResponse{
			Status:  err.StatusCode,
			Error:   err.Error(),
			Message: "Error",
		}
		return c.Status(err.StatusCode).JSON(resErr)
	}

	res := response.SuccessResponse{
		Status:  fiber.StatusOK,
		Data:    result,
		Message: "OK",
	}
	return c.JSON(res)
}

func (r TasksController) GetList(c *fiber.Ctx) error {
	createDate := c.Query("created_date")
	if createDate == "" {
		return fiber.NewError(fiber.StatusBadRequest, "error_created_date_not_valid")
	}

	result, err := r.Srv.GetList(createDate)
	if err != nil {
		return fiber.NewError(err.StatusCode, err.Error())
	}

	res := response.SuccessResponse{
		Status:  fiber.StatusOK,
		Data:    result,
		Message: "OK",
	}
	return c.JSON(res)
}
