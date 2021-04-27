package test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/app/models"
	"github.com/stretchr/testify/suite"
)

type e2eTestSuite struct {
	suite.Suite
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}
func (s *e2eTestSuite) Test_EndToEnd_Login() {
	user := models.User{
		Username: "secondUser",
		Password: "example",
	}
	userStr, err := json.Marshal(user)
	s.NoError(err)
	token, status, err := CreateToken(string(userStr))

	s.NoError(err)
	s.NotEmpty(token)
	s.Equal(http.StatusOK, status)
}

func (s *e2eTestSuite) Test_EndToEnd_Login_Failed_WrongUsernameOrPassword() {
	user := models.User{
		Username: "secondUser",
		Password: "11111",
	}
	userStr, err := json.Marshal(user)
	s.NoError(err)
	_, status, err := CreateToken(string(userStr))

	s.Error(err)
	s.Equal(http.StatusInternalServerError, status)
}

func (s *e2eTestSuite) Test_EndToEnd_GetAllTask() {
	user := models.User{
		Username: "secondUser",
		Password: "example",
	}
	userStr, err := json.Marshal(user)
	s.NoError(err)

	token, status, err := CreateToken(string(userStr))
	s.NoError(err)
	s.NotEmpty(token)
	s.Equal(http.StatusOK, status)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5050/tasks", nil)
	s.NoError(err)

	q := req.URL.Query()
	q.Add("created_date", time.Now().Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)
	defer response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_GetAllTask_Failed_UnAuthorization() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:5050/tasks", nil)
	s.NoError(err)

	q := req.URL.Query()
	q.Add("created_date", time.Now().Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusUnauthorized, response.StatusCode)
	defer response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_GetAllTask_MissingCreateDate() {
	user := models.User{
		Username: "secondUser",
		Password: "example",
	}
	userStr, err := json.Marshal(user)
	s.NoError(err)

	token, status, err := CreateToken(string(userStr))
	s.NoError(err)
	s.NotEmpty(token)
	s.Equal(http.StatusOK, status)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:5050/tasks", nil)
	s.NoError(err)

	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusUnprocessableEntity, response.StatusCode)
	defer response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_AddTask() {
	user := models.User{
		Username: "secondUser",
		Password: "example",
	}
	userStr, err := json.Marshal(user)
	s.NoError(err)

	token, status, err := CreateToken(string(userStr))
	s.NoError(err)
	s.NotEmpty(token)
	s.Equal(http.StatusOK, status)

	client := http.Client{}
	refresh, err := http.NewRequest(http.MethodGet, "http://localhost:5050/refresh-table", nil)
	s.NoError(err)
	_, err = client.Do(refresh)
	s.NoError(err)

	tasks := []models.Task{
		{
			Content: "Statement 1",
		},
		{
			Content: "Statement 2",
		},
		{
			Content: "Statement 3",
		},
		{
			Content: "Statement 4",
		},
		{
			Content: "Statement 5",
		},
	}

	for _, t := range tasks {
		taskStr, err := json.Marshal(t)
		s.NoError(err)
		req, err := http.NewRequest(http.MethodPost, "http://localhost:5050/tasks", strings.NewReader(string(taskStr)))
		s.NoError(err)
		req.Header.Set("Authorization", token)
		response, err := client.Do(req)
		s.NoError(err)
		s.Equal(http.StatusCreated, response.StatusCode)
		defer response.Body.Close()
	}
}

func (s *e2eTestSuite) Test_EndToEnd_AddTask_Limit_Reached() {
	user := models.User{
		Username: "secondUser",
		Password: "example",
	}
	userStr, err := json.Marshal(user)
	s.NoError(err)

	token, status, err := CreateToken(string(userStr))
	s.NoError(err)
	s.NotEmpty(token)
	s.Equal(http.StatusOK, status)

	client := http.Client{}
	refresh, err := http.NewRequest(http.MethodGet, "http://localhost:5050/refresh-table", nil)
	s.NoError(err)
	_, err = client.Do(refresh)
	s.NoError(err)

	tasks := []models.Task{
		{
			Content: "Statement 1",
		},
		{
			Content: "Statement 2",
		},
		{
			Content: "Statement 3",
		},
		{
			Content: "Statement 4",
		},
		{
			Content: "Statement 5",
		},
		{
			Content: "Statement 6",
		},
	}

	for idx, t := range tasks {
		taskStr, err := json.Marshal(t)
		s.NoError(err)
		req, err := http.NewRequest(http.MethodPost, "http://localhost:5050/tasks", strings.NewReader(string(taskStr)))
		s.NoError(err)
		req.Header.Set("Authorization", token)
		response, err := client.Do(req)
		s.NoError(err)
		if idx > 4 {
			s.Equal(http.StatusBadRequest, response.StatusCode)
		} else {
			s.Equal(http.StatusCreated, response.StatusCode)
		}
		defer response.Body.Close()
	}
}
