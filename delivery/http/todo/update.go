package todo

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/khangjig/togo/usecase/todo"
	"github.com/khangjig/togo/util"
	"github.com/khangjig/togo/util/myerror"
)

func (r *Route) update(c echo.Context) error {
	var (
		ctx   = &util.CustomEchoContext{Context: c}
		req   = todo.UpdateRequest{}
		idStr = c.Param("id")
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return util.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	if err = c.Bind(&req); err != nil {
		return util.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	req.ID = id

	resp, err := r.useCase.Todo.Update(ctx, &req)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, resp)
}
