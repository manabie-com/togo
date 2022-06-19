package api

import (
	"github.com/gin-gonic/gin"
	"togo/internal/pkg/http"
	"togo/internal/services/task/application"
)

type UserAPI struct {
	userService *application.UserService
}

func NewUserAPI(
	userService *application.UserService,
) *UserAPI {
	return &UserAPI{userService}
}

func (api *UserAPI) AddRoutes(g *gin.Engine) {
	g.POST("/users", api.create)
}

func (api *UserAPI) create(c *gin.Context) {
	var req CreateUserRequest
	if http.HandleError(c, c.BindJSON(&req)) {
		return
	}
	newUser, err := api.userService.CreateUser(application.AddUserCommand{DailyTaskLimit: req.DailyTaskLimit})
	if http.HandleError(c, err) {
		return
	}
	http.Success(c, newUser, "create user successfully")
}

type CreateUserRequest struct {
	DailyTaskLimit int `json:"dailyTaskLimit" binding:"required,gt=0"`
}
