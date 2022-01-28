package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/manabie-com/togo/pkg/errorx"

	"github.com/go-chi/chi"

	"github.com/manabie-com/togo/pkg/httpx"

	"github.com/manabie-com/togo/internal/user/service"
)

type UserHandler struct {
	userService service.UserService
}

func New(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	defer r.Body.Close()

	req := new(CreateUserRequest)

	if err = json.Unmarshal(b, req); err != nil {
		httpx.WriteError(w, err)
		return
	}

	if err = req.Validate(); err != nil {
		httpx.WriteError(w, err)
		return
	}

	err = h.userService.CreateUser(r.Context(), &service.CreateUserArgs{
		Password:  req.Password,
		LimitTask: req.LimitTask,
	})
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusCreated, nil)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	defer r.Body.Close()

	req := new(UpdateUserRequest)
	if err = json.Unmarshal(b, req); err != nil {
		httpx.WriteError(w, err)
		return
	}

	userID := chi.URLParam(r, "user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		httpx.WriteError(w, errorx.ErrInvalidParameter(err))
	}

	err = h.userService.UpdateUser(r.Context(), &service.UpdateUserArgs{
		UserID:    id,
		Password:  req.Password,
		TaskLimit: &req.TaskLimit,
	})
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusCreated, nil)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	defer r.Body.Close()

	req := new(LoginRequest)
	if err = json.Unmarshal(b, req); err != nil {
		httpx.WriteError(w, err)
		return
	}

	loginRes, err := h.userService.Login(r.Context(), &service.LoginUserArgs{
		UserID:   req.ID,
		Password: req.Password,
	})
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusCreated, LoginUserResponse{
		AccessToken: loginRes.AccessToken,
		AtExpires:   loginRes.AtExpires,
	})
}
