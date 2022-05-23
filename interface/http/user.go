package httpInterface

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"togo/domain/model"
	"togo/domain/service"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService: userService}
}

type RegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (this *userController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"rc": 1,
			"rd": err.Error(),
		})
		return
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	u, err := this.userService.Login(ctx, req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"rc": 1,
			"rd": err.Error(),
		})
		return
	}

	c.Header("bearer", u.Token)
	c.JSON(http.StatusOK, gin.H{
		"rc": 0,
		"rd": "Successfully",
	})
}

func (this *userController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"rc": 1,
			"rd": err.Error(),
		})
		return
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	err := this.userService.Register(ctx, model.User{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"rc": 1,
			"rd": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"rc": 0,
		"rd": "Successfully registered",
	})

}
