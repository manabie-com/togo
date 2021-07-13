package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"

	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/pkg/jwtprovider"
	"github.com/manabie-com/togo/pkg/utils/generator"
	"github.com/manabie-com/togo/server/handler"
)

// e2eTestSuite defines a instance suite test to implement test integration for server.
type e2eTestSuite struct {
	suite.Suite
	dbMigration *migrate.Migrate
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	dbMigration, err := migrate.New("file://..//migrations", config.PostgreSQL.String())
	fmt.Println("MigrationFolder", config.MigrationFolder)
	s.Require().NoError(err)
	if err := dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}
	s.dbMigration = dbMigration
	go Serve()
}

func (s *e2eTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *e2eTestSuite) SetupTest() {
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}
}

func (s *e2eTestSuite) TearDownTest() {
	s.NoError(s.dbMigration.Down())
}

func (s *e2eTestSuite) Test_EndToEnd_LoginHappyCase() {
	userIDTest, passwordTest := generator.NewUUID(), "example"
	tearDown := createUser(s.T(), userIDTest, passwordTest, 5)
	defer tearDown()
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/login?user_id=%s&password=%s", config.HTTPPort, userIDTest, passwordTest), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	resp := &handler.LoginResult{}
	err = json.Unmarshal(byteBody, resp)
	if err != nil {
		s.T().Fatal(err)
	}

	jwtP := jwtprovider.NewJWTProvider(config.JWT.Key, config.JWT.ExpiresIn)
	payload, ok := jwtP.Parse(resp.Data)
	s.Equal(true, ok)
	userID, _ := payload["user_id"].(string)
	s.Equal(userIDTest, userID)
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_LoginFailCases() {

	type input struct {
		userID   string
		password string
	}

	tests := []struct {
		name     string
		input    input
		httpCode int
	}{
		{
			name: "should return error bad request when not send user_id.",
			input: input{
				userID:   "",
				password: "example",
			},
			httpCode: http.StatusBadRequest,
		},
		{
			name: "should return error bad request when not send password.",
			input: input{
				password: "",
				userID:   "test",
			},
			httpCode: http.StatusBadRequest,
		},
		{
			name: "should return error with StatusUnprocessableEntity server when sending wrong user_id or password",
			input: input{
				password: "r",
				userID:   "firstUser",
			},
			httpCode: http.StatusUnprocessableEntity,
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/login?user_id=%s&password=%s", config.HTTPPort, test.input.userID, test.input.password), nil)
			s.NoError(err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			client := http.Client{}
			response, err := client.Do(req)
			s.NoError(err)
			s.Equal(test.httpCode, response.StatusCode)
			response.Body.Close()
		})
	}
}

func (s *e2eTestSuite) Test_EndToEnd_ListTasksSuccess() {
	userIDTest, passwordTest := generator.NewUUID(), "example"
	content := "123"
	tearDown := createUser(s.T(), userIDTest, passwordTest, 5)
	defer tearDown()
	createdDate := time.Now().Format("2006-01-02")
	token, err := getTokenLogin(userIDTest, passwordTest)
	s.NoError(err)

	for i := 0; i < 4; i++ {
		response, err := createTask(s.T(), token, content)
		s.NoError(err)
		s.Equal(http.StatusOK, response.StatusCode)
	}

	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/tasks?created_date=%s", config.HTTPPort, createdDate), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	resp := &handler.ListTasksResult{}
	err = json.Unmarshal(byteBody, resp)
	if err != nil {
		s.T().Fatal(err)
	}
	s.Equal(4, len(resp.Tasks))
	response.Body.Close()
}

// Test_EndToEnd_CreateTaskHappyCase to test for creating task correctly.
func (s *e2eTestSuite) Test_EndToEnd_CreateTaskHappyCase() {

	userIDTest, passwordTest := generator.NewUUID(), "example"
	content := "content123"
	tearDown := createUser(s.T(), userIDTest, passwordTest, 2)
	defer tearDown()
	token, err := getTokenLogin(userIDTest, passwordTest)
	s.NoError(err)

	response, err := createTask(s.T(), token, content)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	resp := &handler.TaskResult{}
	err = json.Unmarshal(byteBody, resp)
	if err != nil {
		s.T().Fatal(err)
	}
	createdDate := time.Now().Format("2006-01-02")

	s.Equal(content, resp.Content)
	s.Equal(userIDTest, resp.UserID)
	s.Equal(createdDate, resp.CreatedDate)

	response.Body.Close()
}

// Test_EndToEnd_CreateTaskReachedOutTaskCreatedLimit
func (s *e2eTestSuite) Test_EndToEnd_CreateTaskReachedOutTaskCreatedLimit() {

	userIDTest, passwordTest, maxTodo := generator.NewUUID(), "example", 2
	content := "content123"
	tearDown := createUser(s.T(), userIDTest, passwordTest, maxTodo)
	defer tearDown()
	token, err := getTokenLogin(userIDTest, passwordTest)
	s.NoError(err)

	// create task to limit
	for i := 0; i < maxTodo; i++ {
		response, err := createTask(s.T(), token, content)
		s.NoError(err)
		s.Equal(http.StatusOK, response.StatusCode)
	}

	response, err := createTask(s.T(), token, content)
	s.NoError(err)

	s.NoError(err)
	s.Equal(http.StatusTooManyRequests, response.StatusCode)
	response.Body.Close()
}

func createTask(t *testing.T, token, content string) (*http.Response, error) {
	bodyString := fmt.Sprintf(`{"content":"%s"}`, content)

	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/tasks", config.HTTPPort), strings.NewReader(bodyString))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", token)
	client := http.Client{}
	return client.Do(req)
}

func getTokenLogin(userID, password string) (string, error) {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/login?user_id=%s&password=%s", config.HTTPPort, userID, password), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	byteBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	resp := &handler.LoginResult{}
	err = json.Unmarshal(byteBody, resp)
	if err != nil {
		return "", err
	}
	return resp.Data, nil
}

func createUser(t *testing.T, userID, password string, maxTodo int) func() {
	db := getPostgresConnection(config.PostgreSQL)
	stmt := `INSERT INTO users (id, password, max_todo) VALUES($1, $2, $3)`
	_, err := db.Exec(stmt, userID, password, maxTodo)
	if err != nil {
		t.Fatal(err)
	}
	return func() {
		clearAllData(db)
	}
}

func clearAllData(db *sql.DB) {
	db.Exec("TRUNCATE TABLE tasks")
	db.Exec("TRUNCATE TABLE users")
}
