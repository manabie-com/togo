package http

import (
	"net/http"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/usecases"
	"togo/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// TodoHandler
type TodoHandler struct {
	usecase usecases.TodoUsecase
}

// Login func
func (h *TodoHandler) Create(c *gin.Context) {
	req := dtos.CreateTodoRequest{}
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

	user := middleware.GetUserFromContext(c)
	err = h.usecase.Create(c, req, user)
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

	data := dtos.BaseResponse{
		Status: http.StatusOK,
	}
	c.JSON(http.StatusOK, data)
}

// NewTodoHandler
func NewTodoHandler(usecase usecases.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		usecase: usecase,
	}
}
