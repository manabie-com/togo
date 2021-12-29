package tests

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
	"net/url"
	"testing"
	"time"
	"togo/db"
	"togo/form"
	api2 "togo/handler"
	"togo/middleware"
)

func TestTodo(t *testing.T) {
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

	req, w, err := setRegister(bytes.NewBuffer(reqBody), database)
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	type myResponse struct {
		Error string `json:"error"`
		Id    int
	}

	actual := new(myResponse)
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	if actual.Error != "" {
		a.Error(fmt.Errorf(actual.Error))
	}

	requestLogin := form.Login{
		Username: requestRegister.Username,
		Password: requestRegister.Password,
	}

	reqBody, err = json.Marshal(requestLogin)
	if err != nil {
		a.Error(err)
	}

	req, w, err = setLogin(bytes.NewBuffer(reqBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err = ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	type myResponseLogin struct {
		Error string `json:"error"`
		Token string `json:"token"`
	}

	actualLogin := new(myResponseLogin)
	if err := json.Unmarshal(body, &actualLogin); err != nil {
		a.Error(err)
	}

	if actualLogin.Error != "" {
		a.Error(fmt.Errorf(actualLogin.Error))
	}

	t.Run("testLoginWrongPassword", func(t *testing.T) {
		requestLogin := form.Login{
			Username: requestRegister.Username,
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

	t.Run("testListTask", func(t *testing.T) {
		requestAddTask := form.Task{
			Content:     "task_test",
		}

		reqBody, err := json.Marshal(requestAddTask)
		if err != nil {
			a.Error(err)
		}

		req, w, err := addTask(bytes.NewBuffer(reqBody),database,actualLogin.Token)
		if err != nil {
			a.Error(err)
		}

		a.Equal(http.MethodPost, req.Method, "HTTP request method error")
		a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

		type taskResponse struct {
			Data struct {
				ID int `json:"ID"`
				CreatedDate string `json:"CreatedDate"`
			}
		}

		body, err = ioutil.ReadAll(w.Body)
		if err != nil {
			a.Error(err)
		}

		actualCreated := new(taskResponse)
		if err := json.Unmarshal(body, &actualCreated); err != nil {
			a.Error(err)
		}


		req, w, err = listTask(actualCreated.Data.CreatedDate,database,actualLogin.Token)
		if err != nil {
			a.Error(err)
		}

		a.Equal(http.MethodGet, req.Method, "HTTP request method error")
		a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

		body, err = ioutil.ReadAll(w.Body)
		if err != nil {
			a.Error(err)
		}

		actualList := new(taskResponse)
		if err := json.Unmarshal(body, &actualList); err != nil {
			a.Error(err)
		}

		if actualList.Data.ID != actualCreated.Data.ID {
			err = fmt.Errorf("error when get task")
			a.Error(err)
		}


		database.Exec("DELETE FROM tasks Where id = ?", actualCreated.Data.ID)
	})

	database.Exec("DELETE FROM users Where username = ?", requestRegister.Username)
}

func listTask(createdDate string, db *gorm.DB, token string) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	api := &api2.APIEnv{DB: db}
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
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.GET("/task", api.GetTask)
	req, err := http.NewRequest(http.MethodGet, "/auth/task", nil)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	urlEnd := url.Values{}
	urlEnd.Add("created_date", createdDate)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.URL.RawQuery = urlEnd.Encode()

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}

func addTask(body *bytes.Buffer, db *gorm.DB, token string) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	api := &api2.APIEnv{DB: db}
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
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.POST("/task", api.CreateTask)
	req, err := http.NewRequest(http.MethodPost, "/auth/task", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
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

func setRegister(body *bytes.Buffer, db *gorm.DB) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	api := &api2.APIEnv{DB: db}
	r.POST("/register", api.Register)
	req, err := http.NewRequest(http.MethodPost, "/register", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w, nil
}
