package module

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/module/auth"
	"github.com/manabie-com/togo/internal/module/task"
	"github.com/manabie-com/togo/internal/module/user"
	"github.com/manabie-com/togo/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	e              *echo.Echo
	authController auth.Controller
	userController user.Controller
	taskController task.Controller
	db             *gorm.DB
	token          string
}

func (s *Suite) SetupSuite() {

	testConfig := config.Config{
		Env:            "local",
		Port:           "5050",
		DbHost:         "localhost",
		DbUser:         "postgres",
		DbPassword:     "password",
		DbName:         "postgres",
		DbPort:         "5432",
		JwtKey:         "ThisIsJwtKey",
		JwtExp:         1440,
		MaxTodoDefault: 5,
	}
	s.db, _ = util.CreateConnectionDB(testConfig)

	s.e = echo.New()
	s.e.Validator = &util.CustomValidator{Validator: validator.New()}

	LoadModules(s.e, s.db)

}

func (s *Suite) TearDownSuite() {

}

func TestModule(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_01_Register() {

	samples := []struct {
		url        string
		statusCode int
		data       string
	}{
		{
			url:        `/user/register`,
			statusCode: http.StatusBadRequest,
			data:       `{"email":"test@mail.com", "password":"12345678"}`,
		},
		{
			url:        `/user/register`,
			statusCode: http.StatusBadRequest,
			data:       `{"email":"test@mail.com", "password":"1234"}`,
		},
		{
			url:        `/user/register`,
			statusCode: http.StatusBadRequest,
			data:       `{"email":"test@mail.com"`,
		},
		{
			url:        `/user/register`,
			statusCode: http.StatusBadRequest,
			data:       `{"email":"testmail.com", "password":"12345678"}`,
		},
	}

	for i, sample := range samples {
		fmt.Printf("TestRegister: %v/%v\n", i+1, len(samples))

		req := httptest.NewRequest(http.MethodPost, sample.url, strings.NewReader(sample.data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		s.e.ServeHTTP(rec, req)

		if assert.Equal(s.T(), sample.statusCode, rec.Code) == false {
			fmt.Println("rec", rec)
		}
	}
}

func (s *Suite) Test_02_Login() {

	samples := []struct {
		url        string
		statusCode int
		data       string
	}{
		{
			url:        `/login`,
			statusCode: http.StatusOK,
			data:       `{"email":"test@mail.com", "password":"12345678"}`,
		},
		{
			url:        `/login`,
			statusCode: http.StatusBadRequest,
			data:       `{"email":"test@mail.com", "password":"1234"}`,
		},
		{
			url:        `/login`,
			statusCode: http.StatusBadRequest,
			data:       `{"email":"test@mail.com"`,
		},
		{
			url:        `/login`,
			statusCode: http.StatusBadRequest,
			data:       `{"email":"testmail.com", "password":"12345678"}`,
		},
	}

	for i, sample := range samples {
		fmt.Printf("TestLogin: %v/%v\n", i+1, len(samples))

		req := httptest.NewRequest(http.MethodPost, sample.url, strings.NewReader(sample.data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		s.e.ServeHTTP(rec, req)

		if assert.Equal(s.T(), sample.statusCode, rec.Code) == false {
			fmt.Println("rec", rec)
		}
	}
}

func (s *Suite) Test_03_UserLogin() {

	sample := struct {
		url        string
		statusCode int
		data       string
	}{

		url:        `/login`,
		statusCode: http.StatusOK,
		data:       `{"email":"test@mail.com", "password":"12345678"}`,
	}

	req := httptest.NewRequest(http.MethodPost, sample.url, strings.NewReader(sample.data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	s.e.ServeHTTP(rec, req)

	if assert.Equal(s.T(), sample.statusCode, rec.Code) {

		dataJSON := make(map[string]interface{})
		err := json.Unmarshal(rec.Body.Bytes(), &dataJSON)
		if err != nil {

		}
		token, ok := dataJSON["token"].(string)
		if !ok {
		}
		s.token = token
	}
}

func (s *Suite) Test_04_AddTask() {
	samples := []struct {
		url        string
		statusCode int
		data       string
	}{
		{
			url:        `/tasks`,
			statusCode: http.StatusCreated,
			data:       `{"content":"content 01"}`,
		},
		{
			url:        `/tasks`,
			statusCode: http.StatusBadRequest,
			data:       ``,
		},
	}

	for i, sample := range samples {
		fmt.Printf("TestAddTask: %v/%v\n", i+1, len(samples))

		req := httptest.NewRequest(http.MethodPost, sample.url, strings.NewReader(sample.data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Add("Authorization", s.token)

		rec := httptest.NewRecorder()
		s.e.ServeHTTP(rec, req)

		if assert.Equal(s.T(), sample.statusCode, rec.Code) == false {
			fmt.Println("rec", rec)
		}
	}
}

func (s *Suite) Test_04_AddManyTasks() {
	samples := []struct {
		url        string
		statusCode int
		data       string
	}{
		{
			url:        `/tasks/many`,
			statusCode: http.StatusTooManyRequests,
			data: `{
				"contents": ["content 01","content 02","content 03","content 04","content 05"]
			}`,
		},
	}

	for i, sample := range samples {
		fmt.Printf("AddManyTasks: %v/%v\n", i+1, len(samples))

		req := httptest.NewRequest(http.MethodPost, sample.url, strings.NewReader(sample.data))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Add("Authorization", s.token)

		rec := httptest.NewRecorder()
		s.e.ServeHTTP(rec, req)

		if assert.Equal(s.T(), sample.statusCode, rec.Code) == false {
			fmt.Println("rec", rec)
		}
	}
}
