package handler

import (
	"encoding/json"
	"net/http"

	mm "github.com/manabie-com/togo/internal/pkg/middleware"
	d "github.com/manabie-com/togo/internal/todo/domain"
	s "github.com/manabie-com/togo/internal/todo/service"
)

type AuthHandler struct {
	AppHandler
	authService *s.AuthService
}

func NewAuthHandler(authService *s.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	rLog := mm.GetLogEntry(r)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	authParam := d.UserAuthParam{}
	if err := decoder.Decode(&authParam); err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}
	if authParam.Password == "" || authParam.Username == "" {
		h.responseError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	token, err := h.authService.ValidateUser(tokenAuth, authParam)
	if err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if token == "" {
		h.responseError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	data := map[string]string{
		"data": token,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		rLog.WithField("err", err).Errorln()
		h.responseError(w, http.StatusInternalServerError, "Error parsing response")
	}
}
