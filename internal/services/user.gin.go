package services

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


type loginUserRequest struct {
	Username string `form:"user_id" binding:"required,alphanum"`
	Password string `form:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
}

var (
	accessDuration = 10 * time.Minute
)

func (server *Server) login(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.Store.RetrieveUser(ctx, req.Username) 
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if req.Password != user.Password {
		err := errors.New("username or password does not match")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.TokenMaker.CreateToken(user.Username, accessDuration) 
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := loginUserResponse{
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, resp)
}
 