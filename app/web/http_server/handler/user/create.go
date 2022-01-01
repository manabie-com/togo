package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manabie-com/togo/app/utils"

	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"

	gErrcode "github.com/manabie-com/togo/app/common/gconstant/errcode"
	gHandler "github.com/manabie-com/togo/app/common/gstuff/handler"
)

// Create created new user
func (s *service) Create(c echo.Context) error {
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
	// validate
	// 1. check user existed by username
	existedUser, err := s.userRepo.GetOneByUsername(httpCtx, request.Username)
	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("get user [%v]: %s", request.Username, err), gErrcode.ServerErrorCommon)
	}
	if existedUser.IsExists() {
		return gHandler.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user [%v] is existed", request.Username), gErrcode.UserErrCommon)
	}

	// get hashedPassword
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, err.Error(), gErrcode.ServerErrorCommon)
	}

	// storerage
	createUserReq := userRepo.CreateReq{
		Username:       request.Username,
		HashedPassword: hashedPassword,
		MaxTasks:       request.MaxTasks,
		// tracing
		CreatedIP: c.RealIP(),
	}

	result, err := s.userRepo.Create(httpCtx, createUserReq)
	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("create user [%v]: %s", request.Username, err), gErrcode.ServerErrorCommon)
	}

	return c.JSON(gHandler.Success(result))
}
