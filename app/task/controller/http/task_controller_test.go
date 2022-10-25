package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ansidev/togo/domain/auth"
	"github.com/ansidev/togo/domain/user"
	"github.com/ansidev/togo/errs"
	"github.com/ansidev/togo/task/dto"
	"github.com/ansidev/togo/test"
	"github.com/ansidev/togo/wire"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const DateFormat = "2006-01-02"

func TestTaskController(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}

type TaskControllerTestSuite struct {
	test.ControllerTestSuite
	tokenTTL time.Duration
}

func (s *TaskControllerTestSuite) SetupSuite() {
	s.ControllerTestSuite.SetupSuite()
	s.tokenTTL = 24 * time.Hour

	authService := wire.InitAuthService(s.GormDb, s.Rdb, 1*time.Minute)
	taskService := wire.InitTaskService(s.GormDb)

	NewTaskController(s.Router, authService, taskService)
}

func (s *TaskControllerTestSuite) TestCreateTask_DoNotGiveBearerToken_ShouldReturn401() {
	w := httptest.NewRecorder()
	body := `{"title":"Sample task"}`
	req, _ := http.NewRequest("POST", "/task/v1/tasks", bytes.NewBufferString(body))
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusUnauthorized, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeTokenIsRequired, actualBody.Code)
	require.Equal(s.T(), errs.ErrTokenIsRequired, actualBody.Message)
	require.Equal(s.T(), "Authorize token is required", actualBody.Error)
}

func (s *TaskControllerTestSuite) TestCreateTask_GiveInvalidAuthorizationHeader_ShouldReturn400() {
	w := httptest.NewRecorder()
	body := `{"title":"Sample task"}`
	req, _ := http.NewRequest("POST", "/task/v1/tasks", bytes.NewBufferString(body))
	req.Header.Add("Authorization", "invalid token")
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusBadRequest, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeInvalidAuthorizationHeader, actualBody.Code)
	require.Equal(s.T(), "Bad Request", actualBody.Message)
	require.True(s.T(), len(actualBody.Error) > 0)
}

func (s *TaskControllerTestSuite) TestCreateTask_GiveInvalidWrongAuthorizeToken_ShouldReturn401() {
	w := httptest.NewRecorder()
	body := `{"title":"Sample task"}`
	req, _ := http.NewRequest("POST", "/task/v1/tasks", bytes.NewBufferString(body))
	req.Header.Add("Authorization", "Bearer wrong_token")
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusUnauthorized, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeUnauthorized, actualBody.Code)
	require.Equal(s.T(), "Unauthorized", actualBody.Message)
	require.True(s.T(), len(actualBody.Error) > 0)
}

func (s *TaskControllerTestSuite) TestCreateTask_GiveTaskWithoutTitle_ShouldReturn400() {
	token := s.initUserData(1, 5)

	w := httptest.NewRecorder()
	body := `{}`
	req, _ := http.NewRequest("POST", "/task/v1/tasks", bytes.NewBufferString(body))
	req.Header.Add("Authorization", "Bearer "+token)
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusBadRequest, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeValidationError, actualBody.Code)
	require.Equal(s.T(), errs.ErrValidationError, actualBody.Message)
	require.Equal(s.T(), 1, len(actualBody.Errors))
	require.Equal(s.T(), "title", actualBody.Errors[0].Field)
	require.Equal(s.T(), "Field 'title' is required", actualBody.Errors[0].Message)
}

func (s *TaskControllerTestSuite) TestCreateTask_GiveTaskWithTooLongTitle_ShouldReturn400() {
	token := s.initUserData(1, 5)

	w := httptest.NewRecorder()
	body := `{"title":"8bbTF9GlkeCfxdUssajQcdM26XsQg37IFELN9z3eXA79DypMKXBTWxffuXGESFOY8hGDpizuMONoBnPNvcOPcNZ2HJcUfuRIBtVoJf8rJR2vQNV7tjlyj9fU8vSeNnRvkEWmqIVh3p7ek0cZEEoSlZMhgfU0Lbi22ZsPm4K34fYul0KuS1I65seSxaSAYxNq9cPomA9bnr58xI7wDmAkgizFSYYp1L49jkerrc4zOE1I4CFsUEHjBKJ2HuwUyMkw"}`
	req, _ := http.NewRequest("POST", "/task/v1/tasks", bytes.NewBufferString(body))
	req.Header.Add("Authorization", "Bearer "+token)
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusBadRequest, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeValidationError, actualBody.Code)
	require.Equal(s.T(), errs.ErrValidationError, actualBody.Message)
	require.Equal(s.T(), 1, len(actualBody.Errors))
	require.Equal(s.T(), "title", actualBody.Errors[0].Field)
	require.Equal(s.T(), "Length of field 'title' may not be greater than 255 characters", actualBody.Errors[0].Message)
}

func (s *TaskControllerTestSuite) TestCreateTask_GiveValidRequest_ShouldReturn201() {
	token := s.initUserData(1, 5)
	w := httptest.NewRecorder()
	body := `{"title":"Sample task"}`
	req, _ := http.NewRequest("POST", "/task/v1/tasks", bytes.NewBufferString(body))
	req.Header.Add("Authorization", "Bearer "+token)
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusCreated, w.Code)

	var actualBody dto.CreateTaskResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Greater(s.T(), actualBody.ID, int64(0))
	require.Equal(s.T(), "Sample task", actualBody.Title)
	require.Equal(s.T(), int64(1), actualBody.OwnerID)
	_, err1 := time.Parse(time.RFC3339, actualBody.CreatedAt)
	require.NoError(s.T(), err1)
	_, err2 := time.Parse(time.RFC3339, actualBody.UpdatedAt)
	require.NoError(s.T(), err2)
}

func (s *TaskControllerTestSuite) TestCreateTask_ReachedDailyLimitTask_ShouldReturn422() {
	userId := int64(1)
	token := s.initUserData(userId, 1)
	s.insertTaskIntoDb(userId, time.Now())

	w := httptest.NewRecorder()
	body := `{"title":"Sample task"}`
	req, _ := http.NewRequest("POST", "/task/v1/tasks", bytes.NewBufferString(body))
	req.Header.Add("Authorization", "Bearer "+token)
	s.Router.ServeHTTP(w, req)

	resp := w.Body.Bytes()

	require.NotNil(s.T(), resp)
	require.Equal(s.T(), http.StatusUnprocessableEntity, w.Code)

	var actualBody errs.ErrorResponse
	err := json.Unmarshal(resp, &actualBody)

	require.NoError(s.T(), err)
	require.Equal(s.T(), errs.ErrCodeReachedLimitDailyTask, actualBody.Code)
	require.Equal(s.T(), errs.ErrReachedLimitDailyTask, actualBody.Error)
}

func (s *TaskControllerTestSuite) initUserData(userId int64, maxDailyTask int) string {
	_, err := s.SqlDb.Exec(`INSERT INTO "user" (id, username, password, max_daily_task, created_at) VALUES ($1, $2, $3, $4, $5)`,
		userId,
		"test_user",
		"$2a$12$IsAJrIc1yhMtlcXC1KfhLOqJSon.NAUMo3KG8NHA9myPm05F85Id2", // test_password
		maxDailyTask,
		"2022-02-22 01:23:45")

	require.NoError(s.T(), err)

	var userModel user.User

	result := s.GormDb.First(&userModel, 1)
	require.NoError(s.T(), result.Error)

	expectedAuthModel := auth.AuthenticationCredential{
		ID:           userId,
		MaxDailyTask: maxDailyTask,
	}

	b, err1 := json.Marshal(expectedAuthModel)
	require.NoError(s.T(), err1)

	token := uuid.NewString()
	cmd := s.RedisClient.Set(context.Background(), token, b, s.tokenTTL)
	_, err2 := cmd.Result()
	require.NoError(s.T(), err2)

	return token
}

func (s *TaskControllerTestSuite) insertTaskIntoDb(userId int64, date time.Time) {
	_, err := s.SqlDb.Exec(`INSERT INTO "task" (title, user_id, created_at) VALUES ($1, $2, $3)`,
		fmt.Sprintf("Task %d", rand.Int()),
		userId,
		fmt.Sprintf("%s %02d:%02d:%02d", date.Format(DateFormat), rand.Intn(23), rand.Intn(59), rand.Intn(59)))

	require.NoError(s.T(), err)
}
