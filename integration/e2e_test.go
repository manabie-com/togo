package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"
	"togo/configs"
	"togo/internal/models"
	"togo/internal/response"
	"togo/migrations/migrate"
	"togo/pkg/databases"
	"togo/pkg/logger"
	"togo/server"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/test-go/testify/suite"
)

type e2eTestSuite struct {
	suite.Suite
	server *server.Server
}

type ResponseAPI struct {
	Status int                    `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(e2eTestSuite))
}

func (s *e2eTestSuite) SetupSuite() {
	logger := logger.NewLogger()
	config := configs.DefaultConfig()
	db := databases.NewPostgres(config.PostgreSQL)

	server := server.Server{
		Config: config,
		Logger: logger,
		Db:     db,
	}

	s.server = &server
	go server.Start()
}

func (s *e2eTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *e2eTestSuite) SetupTest() {
	s.Require().NoError(migrate.Migrate(s.server.Db))
}

func (s *e2eTestSuite) TearDownTest() {
	s.server.Db.Migrator().DropTable(
		&models.User{},
		&models.Task{},
		"public.migrations",
	)
}

func (s *e2eTestSuite) CreateUser() *response.UserResponse {
	request, err := http.NewRequest(
		echo.POST,
		fmt.Sprintf("http://localhost:%s/users", s.server.Config.Server.Port),
		strings.NewReader(`{
			"name": "name day nhe",
			"limit_count": 1
		}`),
	)
	s.NoError(err)

	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	responseRequest, err := client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusOK, responseRequest.StatusCode)

	byteBody, err := ioutil.ReadAll(responseRequest.Body)
	s.NoError(err)

	var body ResponseAPI
	err = json.Unmarshal(byteBody, &body)
	s.NoError(err)

	userResponse := response.UserResponse{
		ID:         int(body.Data["id"].(float64)),
		LimitCount: int(body.Data["limit_count"].(float64)),
		Name:       body.Data["name"].(string),
	}

	fmt.Println(userResponse)

	s.Equal(userResponse.Name, "name day nhe")
	s.Equal(userResponse.LimitCount, 1)
	responseRequest.Body.Close()
	return &userResponse
}

func (s *e2eTestSuite) CreateTask(user *response.UserResponse) *response.TaskResponse {
	request, err := http.NewRequest(
		echo.POST,
		fmt.Sprintf("http://localhost:%s/users/%d/tasks", s.server.Config.Server.Port, user.ID),
		strings.NewReader(`{
			"description": "description",
			"ended_at": "2022-07-27T11:55:37+07:00"
		}`),
	)
	s.NoError(err)

	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	responseRequest, err := client.Do(request)
	s.NoError(err)
	s.Equal(http.StatusOK, responseRequest.StatusCode)

	byteBody, err := ioutil.ReadAll(responseRequest.Body)
	s.NoError(err)

	var body ResponseAPI
	err = json.Unmarshal(byteBody, &body)
	s.NoError(err)

	var taskResponse response.TaskResponse
	err = mapstructure.Decode(body.Data, &taskResponse)
	s.NoError(err)

	s.NoError(err)
	s.Equal(user.ID, taskResponse.ID)
	s.Equal(taskResponse.Description, "description")

	responseRequest.Body.Close()
	return &taskResponse
}

func (s *e2eTestSuite) Test_EndToEnd_FlowCreateUserAndCreateTask() {
	user := s.CreateUser()
	s.CreateTask(user)
}
