package api

import (
	"errors"
	"net/http"
	db "togo/db/sqlc"
	"togo/token"

	"github.com/gin-gonic/gin"
)

type setDailyCapParams struct {
	Username string `json:"username" binding:"required,alphanum"`
	DailyCap int64  `json:"daily_cap" binding:"required,min=0"`
}

func (server *Server) setDailyCap(ctx *gin.Context) {
	var req setDailyCapParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var user db.User
	var err error
	if authPayload.Username != "admin" {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("not an admin")))
		return
	} else {
		arg := db.UpdateUserDailyCapParams{
			Username: req.Username,
			DailyCap: req.DailyCap,
		}
		user, err = server.store.UpdateUserDailyCap(ctx, arg)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, newUserResponse(user))
}
