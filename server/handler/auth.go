package handler

import (
	"net/http"

	"github.com/manabie-com/togo/internal/services/auth"
)

type authHandler struct {
	authSvc auth.Service
}

func NewAuthHandler(authSvc auth.Service) *authHandler {
	return &authHandler{
		authSvc: authSvc,
	}
}

// LoginResult to return login result for client.
type LoginResult struct {
	Data string `json:"data"`
}

func (s *authHandler) Login(resp http.ResponseWriter, req *http.Request) {
	userID := req.FormValue("user_id")
	pwd := req.FormValue("password")
	if userID == "" || pwd == "" {
		respondWithError(resp, http.StatusBadRequest, "user_id and password are required")
		return
	}

	token, err := s.authSvc.Login(req.Context(), userID, pwd)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		if err == auth.ErrWrongAccount {
			httpStatus = http.StatusUnprocessableEntity
		}
		respondWithError(resp, httpStatus, err.Error())
		return
	}

	response := &LoginResult{
		Data: token,
	}
	respondWithJSON(resp, http.StatusOK, response)

}
