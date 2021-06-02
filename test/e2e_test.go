package e2e_test

// Basic imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/dgrijalva/jwt-go"
	internal "github.com/manabie-com/togo/internal"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type E2ETestSuite struct {
	suite.Suite
	Server         *http.Server
	CurrentUserJWT string
}

func (suite *E2ETestSuite) SetupSuite() {
	dataSourceName, _ := filepath.Abs("../data.db")
	suite.Server = internal.NewServer("sqlite3", dataSourceName)
	go suite.Server.ListenAndServe()
}

func (suite *E2ETestSuite) TearDownSuite() {
	suite.Server.Close()
}

type AuthenticationResponse struct {
	JWTToken string `json:"data"`
}

func (suite *E2ETestSuite) Test_Authentication_Request() {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://localhost%s/login?user_id=firstUser&password=example", suite.Server.Addr),
		nil,
	)
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
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("http://localhost%s/tasks", suite.Server.Addr),
		bytes.NewBuffer(tasks),
	)
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
