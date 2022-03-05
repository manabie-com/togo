package todo

import (
	"github.com/labstack/echo/v4"

	"github.com/khangjig/togo/usecase/todo"
	"github.com/khangjig/togo/util"
	"github.com/khangjig/togo/util/myerror"
)

func (r *Route) create(c echo.Context) error {
	var (
		ctx = &util.CustomEchoContext{Context: c}
		req = todo.CreateRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return util.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	resp, err := r.useCase.Todo.Create(ctx, &req)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, resp)
}
