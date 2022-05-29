package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dinhquockhanh/togo/internal/app/user"
	"github.com/dinhquockhanh/togo/internal/pkg/config"
	"github.com/dinhquockhanh/togo/internal/pkg/http/response"
	"github.com/dinhquockhanh/togo/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

type (
	Service interface {
		Auth(ctx context.Context, username, password string) (*user.User, error)
	}
	Handler struct {
		srv       Service
		tokenizer token.Tokenizer
	}
)

func NewHandler(tokenizer token.Tokenizer, srv Service) *Handler {
	return &Handler{
		srv:       srv,
		tokenizer: tokenizer,
	}
}

func (h *Handler) Tokenizer() token.Tokenizer {
	return h.tokenizer
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err)
		return
	}

	usr, err := h.srv.Auth(c.Request.Context(), req.UserName, req.Password)
	if err != nil {
		response.Error(c, err)
		return
	}

	tokenKey, err := h.tokenizer.Create(usr.ToPayload(*config.All.Token.Duration))
	if err != nil {
		response.Error(c, fmt.Errorf("create token: %w", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tokenKey,
	})
}
