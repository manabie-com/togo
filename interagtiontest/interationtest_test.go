package interagtiontest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/manabie-com/togo/constants"
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/usecases/auth"
	"github.com/manabie-com/togo/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite
	host   string
	port   int
	dbConn *gorm.DB
}

func (s *IntegrationTestSuite) SetupSuite() {
	utils.LoadEnv("../.env")

	port, err := strconv.Atoi(utils.Env.Port)
	if err != nil {
		log.Fatal("Can't getenv", err)
	}

	s.port = port
	s.host = utils.Env.Host
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, &IntegrationTestSuite{})
}

func (s *IntegrationTestSuite) TestIntegration_Login_Success() {
	reqBodyStr := `{"username": "manabie", "password": "example"}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%d/login", s.host, s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()
	s.Require().Equal(http.StatusOK, response.StatusCode)
}

func (s *IntegrationTestSuite) TestIntegration_Login_Error() {
	reqBodyStr := `{"username": "username1", "password": "123456"}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%d/login", s.host, s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()

	s.Require().Equal(http.StatusBadRequest, response.StatusCode)
}

func (s *IntegrationTestSuite) TestIntegration_AddTaskSuccess() {
	// get token
	repositories := handlers.NewRepositories(s.dbConn)
	authUsecase := auth.NewAuthUseCase(repositories.Auth)
	token, err := authUsecase.GenerateToken("firstUser", "5")
	s.Require().NoError(err)

	reqBodyStr := `{
		"content": "task_success",
		"create_date": "2022-06-22",
		"userID": "firstUser"
		}`

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%d/tasks", s.host, s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")
	cookie := &http.Cookie{
		Name:   constants.CookieTokenKey,
		Value:  utils.SafeString(token),
		MaxAge: 300,
	}
	req.AddCookie(cookie)

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()
	s.Require().Equal(http.StatusOK, response.StatusCode)
}

func (s *IntegrationTestSuite) TestIntegration_AddTaskFail() {
	// get token
	repositories := handlers.NewRepositories(s.dbConn)
	authUsecase := auth.NewAuthUseCase(repositories.Auth)
	token, err := authUsecase.GenerateToken("firstUser1", "5")
	s.Require().NoError(err)

	reqBodyStr := `{
		"content": "task_success",
		"create_date": "2022-06-22",
		"userID": "firstUser1"
		}`

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%d/tasks", s.host, s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")
	cookie := &http.Cookie{
		Name:   constants.CookieTokenKey,
		Value:  utils.SafeString(token),
		MaxAge: 300,
	}
	req.AddCookie(cookie)

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()
	s.Require().Equal(response.StatusCode, http.StatusBadRequest)

}

func (s *IntegrationTestSuite) TestIntegration_CreateUser_Success() {
	reqBodyStr := `{
		"username": "manabie2",
		"password": "123456",
		"max_task_per_day": 2
		}`

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%d/users", s.host, s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()
	s.Require().Equal(http.StatusOK, response.StatusCode)

	byteResBody, err := ioutil.ReadAll(response.Body)
	s.Require().NoError(err)

	data := map[string]interface{}{}
	err = json.Unmarshal(byteResBody, &data)
	s.Require().NoError(err)

	byteUser, err := json.Marshal(data["data"])
	s.Require().NoError(err)

	user := models.User{}
	err = json.Unmarshal(byteUser, &user)
	s.Require().NoError(err)

	s.Require().Equal(user.Username, "manabie2")
}

func (s *IntegrationTestSuite) TestIntegration_CreateUser_Fail() {
	// get token
	repositories := handlers.NewRepositories(s.dbConn)
	authUsecase := auth.NewAuthUseCase(repositories.Auth)
	token, err := authUsecase.GenerateToken("firstUser1", "5")
	s.Require().NoError(err)

	reqBodyStr := `{
		"username": "manabie",
		"password": "123456",
		"max_task_per_day": 2
		}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%d/users", s.host, s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)
	cookie := &http.Cookie{
		Name:   constants.CookieTokenKey,
		Value:  utils.SafeString(token),
		MaxAge: 300,
	}
	req.AddCookie(cookie)

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()
	s.Require().Equal(http.StatusConflict, response.StatusCode)
}
