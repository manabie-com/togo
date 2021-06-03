package auth

import (
	"database/sql"
	"net/http"

	"github.com/manabie-com/togo/internal/conf"
	"github.com/manabie-com/togo/internal/consts"
	userService "github.com/manabie-com/togo/internal/services/users"
	"github.com/manabie-com/togo/internal/utils/response"

	requestUtils "github.com/manabie-com/togo/internal/utils/request"
	"github.com/manabie-com/togo/internal/utils/token"
)

func NewAuthHandler(database *sql.DB) AuthHandler {
	return AuthHandler{
		NewRequest:  func() ILoginRequest { return &LoginRequest{} },
		UserService: userService.NewService(database),
	}

}

type AuthHandler struct {
	NewRequest  func() ILoginRequest
	UserService userService.IUserService
}

// Handle login operation
func (h AuthHandler) Login(resp http.ResponseWriter, req *http.Request) error {
	request := h.NewRequest()

	// parse request data
	request.Bind(req)

	// validate request, usually used for validating
	// text regex or any data restriction
	if err := request.Validate(); err != nil {
		return err
	}

	// return user id from request, which was binded
	// from request context, updated during the auth middleware,
	userId := request.GetUserID()

	// validate the user existency by querying to database
	if err := h.UserService.Validate(req.Context(), userId, request.GetPassword()); err != nil {
		return err
	}
	jwtSecret := requestUtils.GetJWTSecret(req)
	token, err := token.NewToken(userId,
		jwtSecret,
		conf.DefaultTokenTimeOut, // this value can be configured by env, or pass through server, or conf
		conf.DefaultTokenIssuer,  // this value can be configured by env, or pass through server, or conf
	)
	if err != nil {
		return consts.ErrInvalidAuth
	}

	return response.JSON(resp, token)
}
