package services

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	cmsqlmock "github.com/manabie-com/togo/pkg/common/cmsql/mock"
	"github.com/manabie-com/togo/pkg/common/crypto"
	"github.com/manabie-com/togo/up"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LoginResponse struct {
	Data string `json:"data"`
}

type RegisterResponse struct {
	Data *up.RegisterResponse `json:"data"`
}

func TestToDoService_Login(t *testing.T) {
	mock := integrationTest.mock

	t.Run("TestToDoService_Login", func(t *testing.T) {
		passwordHashed, _ := crypto.HashPassword("example")
		mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
			WithArgs("000001").
			WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}).
				AddRow("000001", passwordHashed, 10))

		loginBody, err := json.Marshal(map[string]string{
			"user_id":  "000001",
			"password": "example",
		})
		assert.Nil(t, err)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(loginBody))
		w := httptest.NewRecorder()
		integrationTest.hander.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("TestToDoService_Login_WrongPassword", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
			WithArgs("000001").
			WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}).
				AddRow("000001", "abc", 10))

		loginBody, err := json.Marshal(map[string]string{
			"user_id":  "000001",
			"password": "password",
		})
		assert.Nil(t, err)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(loginBody))
		w := httptest.NewRecorder()
		integrationTest.hander.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("TestToDoService_Login_WrongID", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
			WithArgs("01234").
			WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}))

		loginBody, err := json.Marshal(map[string]string{
			"user_id":  "01234",
			"password": "password",
		})
		assert.Nil(t, err)

		req := httptest.NewRequest("POST", "/login", bytes.NewReader(loginBody))
		w := httptest.NewRecorder()
		integrationTest.hander.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}

func TestToDoService_Register(t *testing.T) {
	mock := integrationTest.mock

	t.Run("TestToDoService_Register", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
			WithArgs("1").
			WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}))

		mock.ExpectExec("INSERT INTO users (id, password, max_todo) VALUES ($1, $2, $3)").
			WithArgs("1", cmsqlmock.AnyString{}, 11).
			WillReturnResult(sqlmock.NewResult(1, 1))

		registerBody, err := json.Marshal(&up.RegisterRequest{
			ID:       "1",
			Password: "password",
			MaxTodo:  11,
		})

		req := httptest.NewRequest("POST", "/register", bytes.NewReader(registerBody))
		w := httptest.NewRecorder()
		integrationTest.hander.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		bodyStr, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		var registerResp RegisterResponse

		err = json.Unmarshal(bodyStr, &registerResp)
		assert.Nil(t, err)

		assert.Equal(t, "1", registerResp.Data.ID)
		assert.Equal(t, 11, registerResp.Data.MaxTodo)
	})

	t.Run("TestToDoService_Register_UserExists", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, password, max_todo FROM users WHERE id = $1").
			WithArgs("2").
			WillReturnRows(sqlmock.NewRows([]string{"id", "password", "max_todo"}).
				AddRow("2", "password1", 5))

		registerBody, err := json.Marshal(&up.RegisterRequest{
			ID:       "2",
			Password: "password2",
		})

		req := httptest.NewRequest("POST", "/register", bytes.NewReader(registerBody))
		w := httptest.NewRecorder()
		integrationTest.hander.ServeHTTP(w, req)

		resp := w.Result()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		bodyStr, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		var errResponse ErrorResponse

		err = json.Unmarshal(bodyStr, &errResponse)
		assert.Nil(t, err)

		assert.Equal(t, "user exists with the given id", errResponse.Error)
	})
}
