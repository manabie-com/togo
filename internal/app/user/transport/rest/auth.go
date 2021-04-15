package rest

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/app/user/usecase"
	"github.com/manabie-com/togo/internal/util"
	"net/http"
)

var ignorePaths = [1]string{"/login"}

type Delivery struct {
	authService AuthService
	restUtil    util.RestUtil
}

func NewDelivery(authService AuthService) Delivery {
	return Delivery{
		authService: authService,
		restUtil:    util.NewRestUtil(),
	}
}

//go:generate mockgen -package mock -destination mock/auth_mock.go github.com/manabie-com/togo/internal/app/user/transport/rest AuthService
type AuthService interface {
	GetAuthToken(ctx context.Context, userID, pwd string) (string, error)
	Authorize(token string) (string, error)
}

func (d Delivery) Login(resp http.ResponseWriter, req *http.Request) {
	userID := req.FormValue("user_id")
	if userID == "" {
		d.restUtil.WriteFailedResponse(resp, http.StatusBadRequest, errors.New("user_id is missing"))
		return
	}
	pwd := req.FormValue("password")
	if pwd == "" {
		d.restUtil.WriteFailedResponse(resp, http.StatusBadRequest, errors.New("user_id is missing"))
		return
	}
	token, err := d.authService.GetAuthToken(req.Context(), userID, pwd)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrInvalidUser) {
			httpStatus = http.StatusUnauthorized
		}
		d.restUtil.WriteFailedResponse(resp, httpStatus, err)
		return
	}
	d.restUtil.WriteSuccessfulResponse(resp, token)
}

func (d Delivery) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if !d.isInIgnoreList(req.URL.Path) {
			token := req.Header.Get("Authorization")
			userID, err := d.authService.Authorize(token)
			if err != nil {
				d.restUtil.WriteFailedResponse(resp, http.StatusUnauthorized, err)
				return
			}
			req = req.WithContext(util.SetUserIDToContext(req.Context(), userID))
		}
		next.ServeHTTP(resp, req)
	})
}

func (d Delivery) isInIgnoreList(path string) bool {
	for _, ignorePath := range ignorePaths {
		if path == ignorePath {
			return true
		}
	}
	return false
}
