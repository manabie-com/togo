package controller_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io"
	"log"
	jwt_service "manabie-com/togo/config/jwt-service"
	"manabie-com/togo/controller/user"
	"manabie-com/togo/entity"
	"manabie-com/togo/global"
	"manabie-com/togo/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performRequestWithJWTToken(r http.Handler, method, path string, body io.Reader, tokenGiven string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Authorization", tokenGiven)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestCreateUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON  string
		id         string
		maxTodo    int
		password   string
		statusCode int
		errorCode  string
	}{
		{
			inputJSON:  `{"id":"test1", "max_todo": 1, "password": "example"}`,
			id:         "test1",
			maxTodo:    1,
			password:   "example",
			statusCode: http.StatusOK,
			errorCode:  "",
		},
		{
			inputJSON:  `{"id":"test1", "max_todo": 2, "password": "example"}`,
			id:         "test1",
			maxTodo:    1,
			password:   "example",
			statusCode: http.StatusConflict,
			errorCode:  util.ERR_CODE_USER_EXISTED,
		},
	}

	for _, v := range samples {
		router := gin.New()
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
		router.MaxMultipartMemory = 8 << 20 // 8 MiB
		user.RegisterRoutes(router)

		w := performRequest(router, "POST", global.Config.Prefix+"/v1/users", bytes.NewBufferString(v.inputJSON))
		responseMap := make(map[string]interface{})

		err = json.Unmarshal([]byte(w.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, w.Code, v.statusCode)
		if v.statusCode == http.StatusOK {
			assert.Equal(t, responseMap["max_todo"], float64(v.maxTodo))
			assert.Equal(t, responseMap["id"], v.id)

		}
		if v.statusCode == http.StatusConflict {
			assert.Equal(t, responseMap["error_code"], v.errorCode)
		}
	}
}

func TestGetTaskByUserID(t *testing.T) {
	err := refreshUserAndTaskTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser, _, err := seedOneUserAndOneTask()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	user.RegisterRoutes(router)

	claim := jwt_service.Claims{
		Id:      newUser.ID,
		MaxTodo: newUser.MaxTodo,
	}
	var token = jwt_service.CreateJwt(claim)

	w := performRequestWithJWTToken(router, "GET", global.Config.Prefix+"/v1/users/tasks", nil, token)

	var tasks []entity.Task
	err = json.Unmarshal([]byte(w.Body.String()), &tasks)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(tasks), 1)
}

func TestAddTask(t *testing.T) {
	err := refreshUserAndTaskTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser, _, err := seedOneUserAndOneTask()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	user.RegisterRoutes(router)

	claim := jwt_service.Claims{
		Id:      newUser.ID,
		MaxTodo: newUser.MaxTodo,
	}
	var token = jwt_service.CreateJwt(claim)









	w := performRequestWithJWTToken(router, "GET", global.Config.Prefix+"/v1/users/tasks", nil, token)

	var tasks []entity.Task
	err = json.Unmarshal([]byte(w.Body.String()), &tasks)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(tasks), 1)
}
