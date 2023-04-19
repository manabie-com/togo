package services

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strings"
	"togo/pkg/db_client"
	"togo/pkg/models"
	"togo/pkg/utils"
)

type Server struct {
	H     db_client.Handler
	Jwt   utils.JwtWrapper
	Redis *redis.Client
}

func (s *Server) RegisterAccount(ctx *gin.Context) {
	payload := models.AuthRequest{}
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if result := s.H.DB.Where(&models.User{UserName: payload.UserName}).First(&user); result.Error == nil {
		ctx.JSON(http.StatusConflict, "UserName already exists")
		return
	}

	user.UserName = payload.UserName
	user.Password = utils.HashPassword(payload.PassWord)
	user.LimitTask = 5
	s.H.DB.Create(&user)

	var response models.Response[any]
	response.Status = http.StatusOK
	response.Message = "Create User Success"
	ctx.JSON(http.StatusOK, response)
}

func (s *Server) LoginAccount(ctx *gin.Context) {
	payload := models.AuthRequest{}
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if result := s.H.DB.Where(&models.User{UserName: payload.UserName}).First(&user); result.Error != nil {
		ctx.JSON(http.StatusNotFound, "Username not found.")
		return
	}

	checkPass := utils.CheckPasswordHash(payload.PassWord, user.Password)
	if !checkPass {
		ctx.JSON(http.StatusNotFound, "Password invalid.")
		return
	}

	token, _ := s.Jwt.GenerateToken(user)
	var tokenResponse = models.LoginResponse{
		Status:  http.StatusOK,
		Message: "Login Success",
		Token:   token,
	}
	ctx.JSON(http.StatusOK, tokenResponse)
}

func (s *Server) Validate(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, err := s.Jwt.ValidateToken(token[1])
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	if result := s.H.DB.Where(&models.User{Id: claims.Id}).First(&user); result.Error != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("userId", user.Id)
	ctx.Next()
}
