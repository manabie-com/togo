package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/pkg/responses"
	"github.com/manabie-com/togo/internal/repositories/user"
)

func CreateUser(service handlers.MainUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var inputUser models.User
		if err := ctx.ShouldBindJSON(&inputUser); err != nil {
			responses.ResponseForError(ctx, err, http.StatusInternalServerError, "Fail BindJSON user")
			return
		}

		isExistUser, err := service.Auth.ValidateUser(inputUser.Username)
		if err != nil {
			responses.ResponseForError(ctx, err, http.StatusInternalServerError, "Fail ValidateUser")
			return
		}

		if isExistUser {
			responses.ResponseForError(ctx, nil, http.StatusConflict, "User is exists")
			return
		}

		input := user.New(&inputUser)

		if err := service.User.Create(input); err != nil {
			responses.ResponseForError(ctx, nil, http.StatusInternalServerError, "Fail CreateUser")
			return
		}

		responses.ResponseForOK(ctx, http.StatusCreated, inputUser, "Success")
	}
}
