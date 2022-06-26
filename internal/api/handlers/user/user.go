package user

import (
	"net/http"

	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/pkg/responses"
	"github.com/manabie-com/togo/internal/repositories/user"

	"github.com/gin-gonic/gin"
)

type Input struct {
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	MaxTaskPerDay int    `json:"max_task_per_day,omitempty"`
}

func CreateUser(service handlers.MainUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var inputUser Input
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

		input := user.New(&models.User{
			Username:      inputUser.Username,
			Password:      inputUser.Password,
			MaxTaskPerDay: inputUser.MaxTaskPerDay,
		})

		if err := service.User.Create(input); err != nil {
			responses.ResponseForError(ctx, nil, http.StatusInternalServerError, "Fail CreateUser")
			return
		}

		responses.ResponseForOK(ctx, http.StatusCreated, inputUser, "Success")
	}
}
