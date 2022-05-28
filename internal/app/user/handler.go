package user

import (
	"context"
	"net/http"

	"github.com/dinhquockhanh/togo/internal/pkg/http/response"
	"github.com/gin-gonic/gin"
)

type (
	//go:generate mockgen -package user -destination ./internal/app/user/service_mock.go -source internal/app/user/handler.go
	Service interface {
		Create(ctx context.Context, req *CreateUserReq) (*UserSafe, error)
		UpdateUserTier(ctx context.Context, req *UpdateUserTierReq) (*UserSafe, error)
		GetByUserName(ctx context.Context, req *GetUserByUserNameReq) (*User, error)
		List(ctx context.Context, req *ListUsersReq) ([]*UserSafe, error)
		Delete(ctx context.Context, req *DeleteUserByNameReq) error
	}

	Handler struct {
		svc Service
	}
)

func NewHandler(s Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) GetByUserName(gc *gin.Context) {
	var req GetUserByUserNameReq
	if err := gc.ShouldBindUri(&req); err != nil {
		response.Error(gc, err)
		return
	}
	user, err := h.svc.GetByUserName(gc.Request.Context(), &req)
	if err != nil {
		response.Error(gc, err)
		return
	}
	res := user.Safe()
	gc.JSON(http.StatusOK, res)

}

func (h *Handler) CreateUser(gc *gin.Context) {
	var req CreateUserReq
	if err := gc.ShouldBindJSON(&req); err != nil {
		response.Error(gc, err)
		return
	}

	task, err := h.svc.Create(gc.Request.Context(), &req)
	if err != nil {
		response.Error(gc, err)
		return
	}
	gc.JSON(http.StatusCreated, task)

}

func (h *Handler) UpdateUser(gc *gin.Context) {
	var req CreateUserReq
	if err := gc.ShouldBindJSON(&req); err != nil {
		response.Error(gc, err)
		return
	}

	task, err := h.svc.Create(gc.Request.Context(), &req)
	if err != nil {
		response.Error(gc, err)
		return
	}
	gc.JSON(http.StatusCreated, task)

}
