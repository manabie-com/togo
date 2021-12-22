package user_handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/common/response"
	"github.com/shanenoi/togo/common/validation"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/domain"
	"github.com/shanenoi/togo/internal/storages/models"
	"io/ioutil"
	"net/http"
)

func HttpLogin(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)

	user := &models.User{}
	err := json.Unmarshal(body, user)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Failure(config.RESP_JSON_WRONG_FORMAT))
		return
	}

	errors := validation.Validate(user)
	if len(errors) != 0 {
		ctx.JSON(http.StatusOK, response.Failure(errors))
		return
	}

	token := ""
	userDomain := domain.NewUserDomain(ctx)
	token, err = userDomain.LoginUser(user)

	if err == nil {
		ctx.JSON(http.StatusOK, response.Sucess(token))
	} else {
		ctx.JSON(http.StatusOK, response.Failure(err.Error()))
	}
}
