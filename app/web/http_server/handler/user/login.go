package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	gErrcode "github.com/manabie-com/togo/app/common/gconstant/errcode"
	gHandler "github.com/manabie-com/togo/app/common/gstuff/handler"

	"github.com/manabie-com/togo/app/utils"
)

// Login login and return JWT token
func (s *service) Login(c echo.Context) error {
	httpCtx := c.Request().Context()

	type myRequest struct {
		Username string `json:"username" query:"username" validate:"required,max=50"`
		Password string `json:"password" query:"password" validate:"required,max=50"`
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
	if !existedUser.IsExists() {
		return gHandler.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user [%v] is not existed", request.Username), gErrcode.UserErrCommon)
	}
	if !existedUser.IsActive() {
		return gHandler.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user [%v] is not active", request.Username), gErrcode.UserErrCommon)
	}

	// 2. check password
	if err = utils.CheckPassword(request.Password, existedUser.HashedPassword); err != nil {
		return gHandler.NewHTTPError(http.StatusUnauthorized, "user or password wrong", gErrcode.UserErrCommon)
	}

	// 3. gen access token
	accessToken, err := s.tokenMaker.CreateToken(
		existedUser.ID,
		existedUser.Username,
		600,
	)
	if err != nil {
		return gHandler.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("gen token: %s", err), gErrcode.ServerErrorCommon)
	}

	return c.JSON(gHandler.Success(map[string]interface{}{
		"access_token": accessToken,
		"user":         existedUser,
	}))
}
