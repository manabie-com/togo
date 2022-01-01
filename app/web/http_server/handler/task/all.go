package task

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	gErrcode "github.com/manabie-com/togo/app/common/gconstant/errcode"
	gHandler "github.com/manabie-com/togo/app/common/gstuff/handler"
	"github.com/manabie-com/togo/app/utils/token"
)

// All get all task by user
func (s *service) All(c echo.Context) error {
	httpCtx := c.Request().Context()

	type myRequest struct {
		Username string `json:"username" query:"username" validate:"required,max=50"`
		Password string `json:"password" query:"password" validate:"required,max=50"`
		MaxTasks int    `json:"max_tasks" query:"max_tasks" validate:"required,min=1"`
	}
	request := new(myRequest)
	if err := c.Bind(request); err != nil {
		return gHandler.NewHTTPError(http.StatusBadRequest, err.Error(), gErrcode.UserErrCommon)
	}
	if err := c.Validate(request); err != nil {
		return gHandler.NewHTTPError(http.StatusBadRequest, err.Error(), gErrcode.UserErrCommon)
	}

	authPayload := c.Get("authorization_payload").(*token.Payload)
	fmt.Println(authPayload)
	_ = httpCtx

	return c.JSON(gHandler.Success(nil))
}
