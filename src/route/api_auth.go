package route

import (
	"errors"

	"github.com/HoangMV/todo/lib/utils"
	"github.com/HoangMV/todo/src/models/request"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func (r *Route) CheckAuth(ctx *fiber.Ctx) error {
	authToken := string(ctx.Request().Header.Peek("Authorization"))
	if len(authToken) <= 0 {
		return utils.WriteError(ctx, fiber.StatusBadRequest, errors.New("empty auth token"))
	}

	if authToken == viper.GetString("Test.Token") { // for test
		ctx.Locals("user_id", 1)
		return ctx.Next()
	}

	userID := r.biz.CheckAuth(authToken)
	if userID < 0 {
		return utils.WriteError(ctx, fiber.StatusUnauthorized, errors.New("wrong token"))
	}

	ctx.Locals("user_id", userID)
	return ctx.Next()
}

func (r *Route) Register(ctx *fiber.Ctx) error {
	req := &request.LoginReq{}
	if err := ctx.BodyParser(&req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	if err := request.Validate(req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	if err := r.biz.Register(req); err != nil {
		return utils.WriteError(ctx, fiber.StatusInternalServerError, err)
	}

	return utils.WriteSuccessEmptyContent(ctx)
}

func (r *Route) Login(ctx *fiber.Ctx) error {
	req := &request.LoginReq{}
	if err := ctx.BodyParser(&req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	if err := request.Validate(req); err != nil {
		return utils.WriteError(ctx, fiber.StatusBadRequest, err)
	}

	resp, err := r.biz.Login(req)
	if err != nil {
		return utils.WriteError(ctx, fiber.StatusInternalServerError, err)
	}

	return utils.WriteSuccess(ctx, resp)
}
