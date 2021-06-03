package auth

import (
	"net/http"

	"github.com/manabie-com/togo/internal/consts"
	"github.com/manabie-com/togo/internal/utils/request"
)

// To be removed and moved into centralized place

const (
	UserIDParam       = "user_id"
	PasswordParam     = "password"
	MinUserIDLength   = 0
	MinPasswordLength = 0
)

func NewLoginRequest() ILoginRequest {
	return &LoginRequest{}
}

type ILoginRequest interface {
	Bind(req *http.Request)
	Validate() error
	GetUserID() string
	GetPassword() string
}

type LoginRequest struct {
	UserID   string
	Password string
}

// Used for binding requests params, body, etc.
func (r *LoginRequest) Bind(req *http.Request) {
	r.UserID = request.QueryParam(req, UserIDParam)
	r.Password = request.QueryParam(req, PasswordParam)
}

// Currently, I only validate min length of both params
// however, in actual cases, more validation should be done
func (r LoginRequest) Validate() error {
	if len(r.UserID) <= MinUserIDLength {
		return consts.ErrInvalidAuth
	}

	if len(r.Password) <= MinPasswordLength {
		return consts.ErrInvalidAuth
	}
	return nil
}

func (r LoginRequest) GetUserID() string {
	return r.UserID
}

func (r LoginRequest) GetPassword() string {
	return r.Password
}
