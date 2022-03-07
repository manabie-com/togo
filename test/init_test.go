package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/khoale193/togo/models/dbcon"
	"github.com/khoale193/togo/models/migration"
	"github.com/khoale193/togo/pkg/app"
	"github.com/khoale193/togo/pkg/e"
	"github.com/khoale193/togo/pkg/setting"
	"github.com/khoale193/togo/pkg/util"
	"github.com/khoale193/togo/routers"
)

// https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
// https://dev.to/jacobsngoodwin/04-testing-first-gin-http-handler-9m0
func init() {
	setting.Setup()

	dbcon.Setup()
	migration.Migrate()

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

func TestSignIn(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		loginBody := []struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{{Username: "test1", Password: "123456"}, {Username: "test2", Password: "123456"}}
		for _, i := range loginBody {
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			router := routers.InitRouter()
			byte, _ := json.Marshal(i)
			requestLogin, _ := http.NewRequest(http.MethodPost, "/api/sign-in", bytes.NewBuffer(byte))
			router.ServeHTTP(rr, requestLogin)
			var response app.Response
			json.Unmarshal(rr.Body.Bytes(), &response)
			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, "success", response.Status)
		}
	})
	t.Run("Error", func(t *testing.T) {
		loginBody := []struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{{Username: "test", Password: "123456"}, {Username: "test1", Password: "12345"}}
		for _, i := range loginBody {
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			router := routers.InitRouter()
			byte, _ := json.Marshal(i)
			requestLogin, _ := http.NewRequest(http.MethodPost, "/api/sign-in", bytes.NewBuffer(byte))
			router.ServeHTTP(rr, requestLogin)
			var response app.Response
			json.Unmarshal(rr.Body.Bytes(), &response)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, "error", response.Status)
			assert.Equal(t, e.Msg[e.ERROR_AUTH], response.Message)
		}
	})
	t.Run("Validate", func(t *testing.T) {
		loginBody := []struct {
			Data struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			Expected struct {
				Status  string
				Message string
			}
		}{
			{Data: struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}{Username: "", Password: "123456"}, Expected: struct {
				Status  string
				Message string
			}{Status: "error", Message: "Username is a required field"}},
			{Data: struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}{Username: "test", Password: ""}, Expected: struct {
				Status  string
				Message string
			}{Status: "error", Message: "Password is a required field"}},
			{Data: struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}{Username: "", Password: ""}, Expected: struct {
				Status  string
				Message string
			}{Status: "error", Message: "Username is a required field"}},
		}
		for _, i := range loginBody {
			// a response recorder for getting written http response
			rr := httptest.NewRecorder()
			router := routers.InitRouter()
			byte, _ := json.Marshal(i.Data)
			requestLogin, _ := http.NewRequest(http.MethodPost, "/api/sign-in", bytes.NewBuffer(byte))
			router.ServeHTTP(rr, requestLogin)
			var response app.Response
			json.Unmarshal(rr.Body.Bytes(), &response)
			assert.Equal(t, http.StatusBadRequest, rr.Code)
			assert.Equal(t, i.Expected.Status, response.Status)
			assert.Equal(t, i.Expected.Message, response.Message)
		}
	})
}
