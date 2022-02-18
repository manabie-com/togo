package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/api/dto"
	"github.com/kier1021/togo/api/services"
)

// ApiError is the response for validation errors
type ApiError struct {
	Param        string `json:"param"`
	ErrorMessage string `json:"error_message"`
}

// UserController handles the handler for users
type UserController struct {
	userSrv *services.UserService
}

// NewUserController is the constructor for UserController struct
func NewUserController(userSrv *services.UserService) *UserController {
	return &UserController{
		userSrv: userSrv,
	}
}

// CreateUser is the http handler for creation of user
func (ctrl *UserController) CreateUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Bind the body param to CreateUserDTO
		var createUserDto dto.CreateUserDTO
		if err := c.Bind(&createUserDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Bad Request",
				"error":   err.Error(),
			})
			return
		}

		// Request the user service
		results, err := ctrl.userSrv.CreateUser(createUserDto)
		if err != nil {
			makeErrResponse(err, c)
			return
		}

		c.JSON(201, map[string]interface{}{
			"message": "User created successfully.",
			"data":    results,
		})
	}
}

// GetUsers is the http handler for fetching users
func (ctrl *UserController) GetUsers() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Request the user service
		results, err := ctrl.userSrv.GetUsers()
		if err != nil {
			makeErrResponse(err, c)
			return
		}

		c.JSON(201, map[string]interface{}{
			"message": "Users fetched successfully.",
			"data":    results,
		})
	}
}
