package todo

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/khangjig/togo/util"
	"github.com/khangjig/togo/util/myerror"
)

func (r *Route) getByID(c echo.Context) error {
	var (
		ctx   = &util.CustomEchoContext{Context: c}
		idStr = c.Param("id")
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return util.Response.Error(ctx, myerror.ErrInvalidParams(err))
	}

	resp, err := r.useCase.Todo.GetByID(ctx, id)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, resp)
}
