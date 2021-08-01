package transport

import (
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/storages"
)

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	authDomain domain.AuthDomain
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := &storages.User{}
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

	responseWithJson(w, http.StatusOK, map[string]string{
		"data": token,
	})
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := &storages.User{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := h.authDomain.Register(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseWithJson(w, http.StatusOK, map[string]string{
		"status": "success",
	})
}
func NewAuthHandler(authDomain domain.AuthDomain) AuthHandler {
	return &authHandler{
		authDomain: authDomain,
	}
}
