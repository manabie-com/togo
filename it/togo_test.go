package it

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

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/utils"
	"github.com/stretchr/testify/suite"
)

type Data struct {
	Data string `json:"data"`
}

type togoTestSuite struct {
	suite.Suite
	DB      *sql.DB
	appPort string
	jwtKey  string
	autKey  string
}

func TestTogoTestSuite(t *testing.T) {
	suite.Run(t, &togoTestSuite{})
}

func (s *togoTestSuite) SetupSuite() {
	err := utils.InitEnv()
	s.Require().NoError(err)

	db, err := utils.InitDB()
	s.Require().NoError(err)

	s.appPort = os.Getenv("APP_PORT")
	s.jwtKey = os.Getenv("APP_JWTKEY")

	go http.ListenAndServe(":"+s.appPort, &services.TransportService{
		JWTKey: s.jwtKey,
		DB:     db,
	})
}

func (s *togoTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *togoTestSuite) SetupTest() {

}

func (s *togoTestSuite) TearDownTest() {

}

func (s *togoTestSuite) Test_EndToEnd_GetAuthToken() {
	jsonStr := `{"user_id":"itestUser","password":"example"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%v/login", s.appPort), strings.NewReader(jsonStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	var resp Data
	json.Unmarshal(byteBody, &resp)
	s.autKey = resp.Data
}

func (s *togoTestSuite) Test_EndToEnd_ListTasks() {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%v/tasks?created_date=%v", s.appPort, "2021-08-30"), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderAuthorization, s.autKey)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"data":[{"id":"000c824f-5beb-4738-969b-e8b00c9a67d7","content":"some integration testing tasks","user_id":"itestUser","created_date":"2021-08-30"}]}`,
		strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}
