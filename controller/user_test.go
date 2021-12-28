package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"togo/db"
	"togo/form"
	"togo/middleware"
	"togo/model"
)

func TestLogin(t *testing.T) {
	a := assert.New(t)
	database := db.SetupDatabaseConnection(true)

	user, err := insertUser(database)
	if err != nil {
		a.Error(err)
	}

	t.Run("testLoginSuccess", func(t *testing.T) {
		requestLogin := form.Login{
			Username: user.Username,
			Password: user.Password,
		}

		reqBody, err := json.Marshal(requestLogin)
		if err != nil {
			a.Error(err)
		}

		req, w, err := setLogin(bytes.NewBuffer(reqBody))
		if err != nil {
			a.Error(err)
		}

		a.Equal(http.MethodPost, req.Method, "HTTP request method error")
		a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
	})

	t.Run("testLoginWrongPassword", func(t *testing.T) {
		requestLogin := form.Login{
			Username: user.Username,
			Password: "1208312780",
		}

		reqBody, err := json.Marshal(requestLogin)
		if err != nil {
			a.Error(err)
		}

		req, w, err := setLogin(bytes.NewBuffer(reqBody))
		if err != nil {
			a.Error(err)
		}

		a.Equal(http.MethodPost, req.Method, "HTTP request method error")
		a.Equal(http.StatusUnauthorized, w.Code, "HTTP request status code error")
	})

	database.Exec("DELETE FROM users Where username = ?", user.Username)
}

func TestRegister(t *testing.T) {
	a := assert.New(t)
	database := db.SetupDatabaseConnection(true)
	requestRegister := form.User{
		Username: "test123",
		Password: "1234",
	}

	reqBody, err := json.Marshal(requestRegister)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setRegister(bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	t.Run("testRegisterUserExisted", func(t *testing.T) {
		requestRegister := form.User{
			Username: "test123",
			Password: "1234",
		}

		reqBody, err := json.Marshal(requestRegister)
		if err != nil {
			a.Error(err)
		}

		req, w, err := setRegister(bytes.NewBuffer(reqBody))
		if err != nil {
			a.Error(err)
		}

		a.Equal(http.MethodPost, req.Method, "HTTP request method error")
		body, err := ioutil.ReadAll(w.Body)
		if err != nil {
			a.Error(err)
		}

		type myResponse struct {
			error string
		}

		actual := new(myResponse)
		if err := json.Unmarshal(body, &actual); err != nil {
			a.Error(err)
		}

		if actual.error != "" {
			a.Error(fmt.Errorf(actual.error))
		}
	})
	database.Exec("DELETE FROM users Where username = ?", requestRegister.Username)
}

func setLogin(body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "togo-api",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     middleware.IdentityKey,
		PayloadFunc:     middleware.PayloadFunc,
		IdentityHandler: middleware.IdentityHandler,
		Authenticator:   middleware.Authenticator,
		Authorizator:    middleware.Authorizator,
		Unauthorized:    middleware.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
	})
	r.POST("/login", authMiddleware.LoginHandler)
	req, err := http.NewRequest(http.MethodPost, "/login", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}

func setRegister(body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	r.POST("/register", Register)
	req, err := http.NewRequest(http.MethodPost, "/register", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}

func insertUser(db *gorm.DB) (model.User, error) {
	user := model.User{
		Username: "congvantesting",
		Password: "1234",
		MaxTodo:  100,
	}

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
