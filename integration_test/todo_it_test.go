package it

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"

	server "github.com/HoangMV/todo/src"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

func init() {
	viper.SetConfigFile(`../config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

type todoTestSuite struct {
	suite.Suite
}

func TestTodoSuite(t *testing.T) {
	suite.Run(t, &todoTestSuite{})
}

func (s *todoTestSuite) SetupSuite() {
	server := server.New()
	server.SetupConfig()
	go server.Run()
}

func (s *todoTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

func (s *todoTestSuite) TestAPICreateTodo() {

	reqLoginStr := `{
		"content":"test 123"
	}`
	req, err := http.NewRequest(fiber.MethodPost, fmt.Sprintf("http://localhost%s/api/v1/todo", viper.GetString("Server.Port")), strings.NewReader(reqLoginStr))
	s.NoError(err)

	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Set(fiber.HeaderAuthorization, viper.GetString("Test.Token"))

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"Message":"Success","Status":200}`, strings.Trim(string(byteBody), "\n"))

	response.Body.Close()
}
