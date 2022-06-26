package common

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/m/v2/constants"
	"example.com/m/v2/internal/api/handlers"
	"example.com/m/v2/internal/pkg/responses"
	"example.com/m/v2/utils"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(service handlers.MainUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var inputUser Input
		if err := ctx.ShouldBindJSON(&inputUser); err != nil {
			responses.ResponseForError(ctx, err, http.StatusBadRequest, "Fail BindJSON user")
			return
		}

		user, err := service.User.Login(inputUser.Username, inputUser.Password)
		if err != nil {
			responses.ResponseForError(ctx, err, http.StatusBadRequest, "Fail Login")
			return
		}

		if user == nil {
			responses.ResponseForError(ctx, nil, http.StatusBadRequest, "Fail Login")
			return
		}

		maxTaskPerDay := strconv.Itoa(user.MaxTaskPerDay)

		token, err := service.Auth.GenerateToken(user.ID, maxTaskPerDay)
		if err != nil {
			responses.ResponseForError(ctx, err, http.StatusInternalServerError, "Fail GenerateToken")
			return
		}

		if utils.SafeString(token) == "" {
			responses.ResponseForError(ctx, nil, http.StatusInternalServerError, "Fail GenerateToken")
			return
		}

		//Set Cookie
		ctx.SetCookie(constants.CookieTokenKey, utils.SafeString(token), 60*60*24, "/", fmt.Sprintf("%s:%s", utils.Env.DBHost, utils.Env.DBPort), true, true)

		responses.ResponseForOK(ctx, http.StatusOK, nil, "Success")
	}
}
