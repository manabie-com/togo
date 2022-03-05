package user

import (
	"github.com/labstack/echo/v4"

	"github.com/khangjig/togo/util"
	"github.com/khangjig/togo/util/myerror"
)

func (r *Route) getMe(c echo.Context) error {
	ctx := &util.CustomEchoContext{Context: c}

	resp, err := r.useCase.User.GetMe(ctx)
	if err != nil {
		return util.Response.Error(c, err.(myerror.MyError))
	}

	return util.Response.Success(c, resp)
}
