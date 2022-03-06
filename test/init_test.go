package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"togo/models/dbcon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"togo/models"
	"togo/pkg/setting"
	"togo/pkg/util"
	"togo/routers"
)

// https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
// https://dev.to/jacobsngoodwin/04-testing-first-gin-http-handler-9m0
func init() {
	setting.Setup()

	dbcon.Setup()
	models.Migrate()

	util.Setup()
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	log.Println("Everything above here run before ALL test")
	// Run test suites
	exitVal := m.Run()
	log.Println("Everything below run after ALL test")
	// we can do clean up code here
	os.Exit(exitVal)
}

// EnRSA Encrypt
func TestEnRSA(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	t.Run("Success", func(t *testing.T) {
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()
		router := routers.InitRouter()

		var body struct {
			Text string `json:"text"`
		}
		body.Text = "password"
		byte, _ := json.Marshal(body)
		request, _ := http.NewRequest(http.MethodPost, "/api/dev/en_rsa", bytes.NewBuffer(byte))
		//assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		//fmt.Print(rr.Body)
		var resp struct {
			Data string `json:"data"`
		}
		json.Unmarshal(rr.Body.Bytes(), &resp)
		//fmt.Print(resp.Data)

		var loginBody struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		loginBody.Username = "dev"
		loginBody.Password = resp.Data
		byte, _ = json.Marshal(loginBody)
		fmt.Print(string(byte))
		requestLogin, _ := http.NewRequest(http.MethodPost, "/api/admin/sign-in", bytes.NewBuffer(byte))
		router.ServeHTTP(rr, requestLogin)
		fmt.Print(rr.Body)
		//respBody, err := json.Marshal(gin.H{
		//	"data": "encryptRSA(form.Text)",
		//})
		//assert.NoError(t, err)

		assert.Equal(t, 200, rr.Code)
		//assert.Equal(t, respBody, rr.Body.Bytes())
		fmt.Print(t)
		//fmt.Print(rr.Body.Bytes())
		//mockUserService.AssertExpectations(t) // assert that UserService.Get was called
	})
}
