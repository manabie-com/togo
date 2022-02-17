package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/api/apierrors.go"
	"github.com/kier1021/togo/api/dto"
	"github.com/kier1021/togo/api/services"
)

type UserController struct {
	userSrv *services.UserService
}

func NewUserController(userSrv *services.UserService) *UserController {
	return &UserController{
		userSrv: userSrv,
	}
}

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

			// Check if the error is UserAlreadyExists
			if errors.Is(err, apierrors.UserAlreadyExists) {
				c.AbortWithStatusJSON(422, map[string]interface{}{
					"message": "Error in data input",
					"error":   err.Error(),
				})
				return
			}

			c.AbortWithStatusJSON(500, map[string]interface{}{
				"message": "Internal server error occurred.",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(201, map[string]interface{}{
			"message": "User created successfully.",
			"data":    results,
		})
	}
}

func (ctrl *UserController) GetUsers() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Request the user service
		results, err := ctrl.userSrv.GetUsers()
		if err != nil {
			c.AbortWithStatusJSON(500, map[string]interface{}{
				"message": "Internal server error occurred.",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(201, map[string]interface{}{
			"message": "Users fetched successfully.",
			"data":    results,
		})
	}
}
