package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const db = "data.test.db"

type IntegrationTestSuite struct {
	suite.Suite
	cfg *config.Config
	api *API
	token string
}

func (s *IntegrationTestSuite) TearDownSuite() {
	// remove file test db
	os.Remove(db)
}

func (s *IntegrationTestSuite) SetupSuite() {
	// remove db anyway to make sure no database exist in dir
	os.Remove(db)
	// init miniredis
	m, err := miniredis.Run()
	s.NoError(err)

	// new config
	s.cfg = &config.Config{
		MaxTodo:     5,
		Environment: "D",
		Port:        5050,
		JWTKey:      "testKey",
		SQLite:      "data.test.db",
		Redis:       &config.Redis{
			Address:                m.Addr(),
			Password:               "",
			DatabaseNum:            0,
			MaxIdle:                3,
			MaxActive:              0,
			MaxIdleTimeout:         300,
			Wait:                   false,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
		},
	}
	s.api, err = NewAPI(s.cfg)
	s.NoError(err)
}

func (s *IntegrationTestSuite) Test_1_Signup() {
	user := &storages.User{
		ID: "firstUser",
		Password: "example",
	}
	bData, err := json.Marshal(user)
	s.NoError(err)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(bData))
	s.NoError(err)

	resp := httptest.NewRecorder()
	s.api.Router.ServeHTTP(resp, req)

	result := resp.Result()
	s.Equal(http.StatusOK, result.StatusCode)

	key := &ResponseData{}
	err = json.NewDecoder(result.Body).Decode(key)
	s.NoError(err)
	s.NotEqual("", key.Data.(string))
	println(fmt.Sprintf("return key from signup is %v", key.Data))
}

func (s *IntegrationTestSuite) Test_2_Login() {
	user := &storages.User{
		ID: "firstUser",
		Password: "example",
	}
	bData, err := json.Marshal(user)
	s.NoError(err)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bData))
	s.NoError(err)

	resp := httptest.NewRecorder()
	s.api.Router.ServeHTTP(resp, req)

	result := resp.Result()
	s.Equal(http.StatusOK, result.StatusCode)

	responseData := &ResponseData{}
	err = json.NewDecoder(result.Body).Decode(responseData)
	s.NoError(err)
	s.NotEqual("", responseData.Data.(string))
	s.token = responseData.Data.(string)
	println(fmt.Sprintf("return key from login is %v", responseData.Data))
}

func addTask(s *IntegrationTestSuite, count int, hasToken, success bool, failedMsg string) {
	content := fmt.Sprintf("task %d", count)
	task := &storages.Task{
		Content:     content,
		UserID:      "firstUser",
	}
	bData, err := json.Marshal(task)
	s.NoError(err)

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(bData))
	s.NoError(err)

	// set authorization
	if hasToken {
		req.Header.Set("Authorization", s.token)
	}

	resp := httptest.NewRecorder()
	s.api.Router.ServeHTTP(resp, req)
	result := resp.Result()
	responseData := &ResponseData{}
	err = json.NewDecoder(result.Body).Decode(responseData)
	s.NoError(err)

	if success {
		s.Equal(http.StatusOK, result.StatusCode)

		responseTask := &storages.Task{}
		err := mapstructure.Decode(responseData.Data, responseTask)
		s.NoError(err)

		s.Equal(task.Content, responseTask.Content)
		s.Equal(task.UserID, responseTask.UserID)
	} else {
		s.Equal(http.StatusUnauthorized, result.StatusCode)
		if failedMsg == "" {
			s.Equal(MaxLimitReach.Error(), responseData.Data.(string))
		} else {
			s.Equal(failedMsg, responseData.Data.(string))
		}
	}
}

func (s *IntegrationTestSuite) Test_3_AccessWithoutAuth() {
	addTask(s, 1, false, false, "Unauthorized")
}

func (s *IntegrationTestSuite) Test_4_AddTask() {
	addTask(s, 1, true, true, "")
}

func (s *IntegrationTestSuite) Test_5_AddTask_ReachLimit() {
	// because we did an addTask above, therefore expected count must be 1
	count := 1
	for int32(count) < s.cfg.MaxTodo {
		count += 1
		addTask(s, count, true, true, "")
	}
	// add one more and expect fail
	addTask(s, count+1, true, false, "")
}

func (s *IntegrationTestSuite) Test_6_ListTask() {
	req, err := http.NewRequest("GET", "/tasks", nil)
	s.NoError(err)

	// set authorization
	req.Header.Set("Authorization", s.token)

	resp := httptest.NewRecorder()
	s.api.Router.ServeHTTP(resp, req)
	result := resp.Result()
	responseData := &ResponseData{}
	err = json.NewDecoder(result.Body).Decode(responseData)
	s.NoError(err)

	s.Equal(http.StatusOK, result.StatusCode)
	var tasks []*storages.Task
	err = mapstructure.Decode(responseData.Data, &tasks)
	s.NoError(err)
	s.Len(tasks, int(s.cfg.MaxTodo))

	for i, task := range tasks {
		expectedContent := fmt.Sprintf("task %d", i+1)
		expectedUser := "firstUser"
		s.Equal(expectedContent, task.Content)
		s.Equal(expectedUser, task.UserID)
	}
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
