package http_test

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/domain"
	taskhttp "github.com/manabie-com/togo/internal/transport/http"
	"github.com/manabie-com/togo/internal/transport/http/middleware"
	"github.com/manabie-com/togo/internal/transport/http/request"
	"github.com/manabie-com/togo/mocks"
	"github.com/manabie-com/togo/pkg/token"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func addAuthorization(
	t *testing.T,
	req *http.Request,
	jwtMaker token.Token,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	accessToken, err := jwtMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, accessToken)
	req.Header.Set(middleware.AuthorizationHeaderKey, authorizationHeader)
}

func TestCreateTaskSuccess(t *testing.T) {
	content := "sample"
	username := "firstUser"
	mockCreateTaskRequest := request.CreateTaskRequest{
		Content: content,
	}
	jwtMaker := token.NewJWTMaker("123456")
	mockCreateTaskRequestBody, err := json.Marshal(mockCreateTaskRequest)
	require.NoError(t, err)
	mockTaskUseCase := new(mocks.TaskUseCase)
	mockTaskUseCase.On("CreateTask", mock.Anything, content, username).
		Return(nil)
	e := echo.New()
	authMiddleware := middleware.NewAuthMiddleware(jwtMaker)
	taskhttp.NewTaskHandler(e, mockTaskUseCase, authMiddleware)
	req := httptest.NewRequest(echo.POST, "/tasks", strings.NewReader(string(mockCreateTaskRequestBody)))
	addAuthorization(t, req, jwtMaker, middleware.AuthorizationTypeBearer, username, time.Minute*15)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestCreateTaskMissingContent(t *testing.T) {
	mockCreateTaskRequest := request.CreateTaskRequest{
		Content: "",
	}
	jwtMaker := token.NewJWTMaker("123456")
	mockCreateTaskRequestBody, err := json.Marshal(mockCreateTaskRequest)
	require.NoError(t, err)
	mockTaskUseCase := new(mocks.TaskUseCase)
	mockTaskUseCase.On("CreateTask", mock.Anything, "sample", "firstUser").
		Return(nil)
	e := echo.New()
	authMiddleware := middleware.NewAuthMiddleware(jwtMaker)
	taskhttp.NewTaskHandler(e, mockTaskUseCase, authMiddleware)
	req := httptest.NewRequest(echo.POST, "/tasks", strings.NewReader(string(mockCreateTaskRequestBody)))
	addAuthorization(t, req, jwtMaker, middleware.AuthorizationTypeBearer, "f", time.Minute*15)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateTaskUseCaseErrorMaximumTaskPerDay(t *testing.T) {
	content := "sample"
	username := "firstUser"
	mockCreateTaskRequest := request.CreateTaskRequest{
		Content: content,
	}
	jwtMaker := token.NewJWTMaker("123456")
	mockCreateTaskRequestBody, err := json.Marshal(mockCreateTaskRequest)
	require.NoError(t, err)
	mockTaskUseCase := new(mocks.TaskUseCase)
	mockTaskUseCase.On("CreateTask", mock.Anything, content, username).
		Return(domain.ErrorMaximumTaskPerDay)
	e := echo.New()
	authMiddleware := middleware.NewAuthMiddleware(jwtMaker)
	taskhttp.NewTaskHandler(e, mockTaskUseCase, authMiddleware)
	req := httptest.NewRequest(echo.POST, "/tasks", strings.NewReader(string(mockCreateTaskRequestBody)))
	addAuthorization(t, req, jwtMaker, middleware.AuthorizationTypeBearer, username, time.Minute*15)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateTaskBindingError(t *testing.T) {
	bodyJson := `{"content": true}`
	jwtMaker := token.NewJWTMaker("123456")
	mockTaskUseCase := new(mocks.TaskUseCase)
	mockTaskUseCase.On("CreateTask", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(nil)
	e := echo.New()
	authMiddleware := middleware.NewAuthMiddleware(jwtMaker)
	taskhttp.NewTaskHandler(e, mockTaskUseCase, authMiddleware)
	req := httptest.NewRequest(echo.POST, "/tasks", strings.NewReader(bodyJson))
	addAuthorization(t, req, jwtMaker, middleware.AuthorizationTypeBearer, "firstUser", time.Minute*15)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}
