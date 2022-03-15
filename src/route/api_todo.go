package route

import (
	"errors"

	"github.com/HoangMV/todo/lib/utils"
	"github.com/HoangMV/todo/src/models/request"
	"github.com/gofiber/fiber/v2"
)

func (r *Route) CreateUserTodo(ctx *fiber.Ctx) error {

	req := &request.CreateTodoReq{}
	if err := ctx.BodyParser(&req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	if err := request.Validate(req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	var ok bool
	req.UserID, ok = ctx.Locals("user_id").(int)
	if !ok {
		return utils.WriteError(ctx, fiber.StatusBadRequest, errors.New("not found user id from token"))
	}

	// biz
	if err := r.biz.CreateTodo(req); err != nil {
		return utils.WriteError(ctx, fiber.StatusInternalServerError, err)
	}

	return utils.WriteSuccessEmptyContent(ctx)
}

func (r *Route) UpdateUserTodo(ctx *fiber.Ctx) error {

	req := &request.UpdateTodoReq{}
	if err := ctx.BodyParser(&req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	if err := request.Validate(req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	// biz
	err := r.biz.UpdateTodo(req)

	if err != nil {
		return utils.WriteError(ctx, fiber.StatusInternalServerError, err)
	}

	return utils.WriteSuccessEmptyContent(ctx)
}

func (r *Route) GetUserListTodo(ctx *fiber.Ctx) error {

	req := &request.GetTodosReq{}
	if err := ctx.QueryParser(req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	if req.Index < 0 {
		req.Index = 0
	}
	if req.Size <= 0 {
		req.Size = 20
	}
	var ok bool
	req.UserID, ok = ctx.Locals("user_id").(int)
	if !ok {
		return utils.WriteError(ctx, fiber.StatusBadRequest, errors.New("not found user id from token"))
	}

	// biz
	resp, err := r.biz.GetListUserTodo(req)

	if err != nil {
		return utils.WriteError(ctx, fiber.StatusInternalServerError, err)
	}

	return utils.WriteSuccess(ctx, resp)
}
