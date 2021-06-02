package e2e_test

// Basic imports
import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type E2ETestSuite struct {
	suite.Suite
	DriverName     string
	DataSourceName string
	DB             *sql.DB
	Port           string
	JWTKey         string
	CurrentUserJWT string
}

func (suite *E2ETestSuite) SetupSuite() {
	suite.DriverName = "sqlite3"
	suite.DataSourceName, _ = filepath.Abs("../data.db")
	suite.Port = "8050"
	suite.JWTKey = "wqGyEBBfPK9w3Lxw"

	var err error
	suite.DB, err = sql.Open(suite.DriverName, suite.DataSourceName)
	suite.Require().NoError(err)

	go http.ListenAndServe(fmt.Sprintf(":%s", suite.Port), &services.ToDoService{
		JWTKey: suite.JWTKey,
		Store: &sqllite.LiteDB{
			DB: suite.DB,
		},
	})
}

func (suite *E2ETestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

type AuthenticationResponse struct {
	JWTToken string `json:"data"`
}

func (suite *E2ETestSuite) Test_Authentication_Request() {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%s/login?user_id=firstUser&password=example", suite.Port), nil)
	suite.NoError(err)
	client := http.Client{}
	resp, err := client.Do(req)
	suite.NoError(err)

	suite.Equal(http.StatusOK, resp.StatusCode)
	byte, err := ioutil.ReadAll(resp.Body)
	suite.NoError(err)

	var jsonResp AuthenticationResponse
	err = json.Unmarshal(byte, &jsonResp)
	suite.NoError(err)
	token, _ := jwt.Parse(jsonResp.JWTToken, nil)
	// TODO exp is later than now
	suite.Equal(jsonResp.JWTToken, token.Raw)
	suite.CurrentUserJWT = jsonResp.JWTToken
	resp.Body.Close()
}

type CreateNewTaskResponse struct {
	Task storages.Task `json:"data"`
}

func (suite *E2ETestSuite) Test_Create_New_Task_Request() {
	suite.Test_Authentication_Request()
	var tasks = []byte(`{"contents": "example"}`)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%s/tasks", suite.Port), bytes.NewBuffer(tasks))
	suite.NoError(err)

	req.Header.Set("Authorization", suite.CurrentUserJWT)

	client := http.Client{}
	resp, err := client.Do(req)
	suite.NoError(err)

	suite.Equal(http.StatusOK, resp.StatusCode)
	byte, err := ioutil.ReadAll(resp.Body)
	suite.NoError(err)

	var jsonResp CreateNewTaskResponse
	err = json.Unmarshal(byte, &jsonResp)
	suite.NoError(err)

	suite.Equal(jsonResp.Task.UserID, "firstUser")
	resp.Body.Close()
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}
