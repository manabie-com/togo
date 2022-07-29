package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"togo/internal/response"
	"togo/internal/user/dto"
	"togo/internal/user/mocks"
	"togo/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/test-go/testify/assert"
)

func mockCreateUserDtoAndResponse() (*dto.CreateUserDto, *response.UserResponse) {
	createUserDto := &dto.CreateUserDto{
		Name:       "name",
		LimitCount: 1,
	}
	response := &response.UserResponse{
		ID:         1,
		Name:       createUserDto.Name,
		LimitCount: createUserDto.LimitCount,
	}
	return createUserDto, response
}

func mockRequestCreateUser() (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = validator.NewValidator()
	json := `{
		"name": "name",
		"limit_count": 1
	}`
	request := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(json))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)

	return e, context, recorder
}

func TestUserHandler_CreateUserSuccess(t *testing.T) {
	createUserDto, response := mockCreateUserDtoAndResponse()
	e, context, recorder := mockRequestCreateUser()

	userGroup := e.Group("/users")
	service := mocks.NewUserService(t)
	service.On("Create", createUserDto).Return(response, nil)

	handler := NewUserHandler(userGroup, service)

	err := handler.Create(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	response1 := `{"data":{"id":1,"limit_count":1,"name":"name"},"status":200}`
	response2 := `{"status":200,"data":{"id":1,"limit_count":1,"name":"name"}}`
	dataResponse := strings.Trim(recorder.Body.String(), "\n")
	assert.True(t, dataResponse == response1 || dataResponse == response2)
}

func TestUserHandler_HandleCreateUser_ServiceReturnError(t *testing.T) {
	createUserDto, _ := mockCreateUserDtoAndResponse()
	e, context, recorder := mockRequestCreateUser()

	userGroup := e.Group("/users")
	service := mocks.NewUserService(t)
	mockError := errors.New("mock_err")
	service.On("Create", createUserDto).Return(nil, mockError)

	handler := NewUserHandler(userGroup, service)

	err := handler.Create(context)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, recorder.Code)

	response1 := `{"data":"mock_err","status":400}`
	response2 := `{"status":400,"data":"mock_err"}`
	dataResponse := strings.Trim(recorder.Body.String(), "\n")
	assert.True(t, dataResponse == response1 || dataResponse == response2)
}
