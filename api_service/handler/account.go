package handler

import (
	"log"
	"net/http"

	"api_service/connection"
	"api_service/proto"

	"github.com/gin-gonic/gin"
)

//Create ...
func Create(ctx *gin.Context) {

	var createRequest proto.CreateRequest
	if err := ctx.ShouldBindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Create(ctx, &createRequest)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

//Login ...
func Login(ctx *gin.Context) {

	var loginRequest proto.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Login(ctx, &loginRequest)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Value:    response.Token,
		HttpOnly: true,
	})

	ctx.JSON(http.StatusOK, response)
}

//Logout ...
func Logout(ctx *gin.Context) {

	var logoutRequest proto.LogoutRequest
	if err := ctx.ShouldBindJSON(&logoutRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToAccountServiceServer()
	response, err := sc.ClientAccountService.Logout(ctx, &logoutRequest)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		HttpOnly: true,
	})

	ctx.JSON(http.StatusOK, response)
}

func Update(ctx *gin.Context) {

}

func Show(ctx *gin.Context) {

}
