package http

import (
	"net/http"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/usecases"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// AuthHandler
type AuthHandler struct {
	usecase usecases.AuthUsecase
}

// Login func
func (h *AuthHandler) Login(c *gin.Context) {
	req := dtos.LoginRequest{}
	err := c.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		data := dtos.BaseResponse{
			Status: http.StatusBadRequest,
			Error: &dtos.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	token, err := h.usecase.Login(c, req)
	if err != nil {
		data := dtos.BaseResponse{
			Status: http.StatusNotFound,
			Error: &dtos.ErrorResponse{
				ErrorMessage: err.Error(),
			},
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	data := dtos.BaseResponse{
		Status: http.StatusOK,
		Data: dtos.LoginResponse{
			Token: token,
		},
	}
	c.JSON(http.StatusOK, data)
}

// NewAuthHandler
func NewAuthHandler(usecase usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}
