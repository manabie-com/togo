package transport

import (
	"encoding/json"
	"net/http"

	"manabie/togo/internal/domain"
	"manabie/togo/internal/model"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authDomain domain.AuthDomain
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	token, err := h.authDomain.Login(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseWithJson(w, http.StatusOK, token)
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := h.authDomain.Register(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseWithJson(w, http.StatusOK, true)
}
func NewAuthHandler(authDomain domain.AuthDomain) AuthHandler {
	return &authHandler{
		authDomain: authDomain,
	}
}
