package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/khoale193/togo/models"
	"github.com/khoale193/togo/pkg/app"
	"github.com/khoale193/togo/pkg/e"
	"github.com/khoale193/togo/pkg/util"
)

func SignIn(c *gin.Context) {
	var appG = app.Gin{C: c}
	form := LoginFormValidator{}
	if httpCode, err := form.BindAndValid(c); err != nil {
		appG.ResponseError(httpCode, app.NewValidatorError(err), nil)
		return
	}
	if valid, err := isCorrectUser(form.Username, strings.TrimSpace(form.Password)); err != nil || !valid {
		appG.Response(c, http.StatusBadRequest, e.Msg[e.ERR], e.ERROR_AUTH, nil)
	} else {
		user, err := (models.Member{}).FindByUsername(strings.TrimSpace(form.Username))
		token, err := util.GenerateToken(form.Username, e.AppRoleMember, int(user.ID), 24*730*time.Hour)
		if err != nil {
			appG.Response(c, http.StatusBadRequest, e.Msg[e.ERROR_AUTH], e.ERROR_AUTH, nil)
			return
		}
		data := make(map[string]interface{})
		data["access_token"] = token
		host := c.Request.Host
		if i := strings.Index(c.Request.Host, ":"); i != -1 {
			host = c.Request.Host[:i]
		}
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "access_token",
			Value:    token,
			Domain:   host,
			Path:     c.Request.URL.Path,
			Expires:  time.Now().Add(time.Second * 1000),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		})
		appG.Response(c, http.StatusOK, e.Msg[e.SUCCESS], e.SUCCESS_LOGIN, data)
		return
	}
}

type LoginForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginFormValidator struct {
	LoginForm
}

func (v *LoginFormValidator) BindAndValid(c *gin.Context) (int, error) {
	err := app.BindAndValid(c, &v.LoginForm)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func isCorrectUser(username, password string) (bool, error) {
	data, err := (models.Member{}).FindByUsername(username)
	if err != nil {
		return false, err
	}
	return util.VerifyPassword(password, data.Password, e.AuthenticationTypeMD5), nil
}
