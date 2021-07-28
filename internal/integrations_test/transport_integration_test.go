package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	testfixture "github.com/manabie-com/togo/internal/database/testfixtures"
	errutil "github.com/manabie-com/togo/internal/pkg/error_utils"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/transport"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

var fixturePath = "../database/testfixtures/fixtures"

func init() {
	os.Setenv("SECRET_KEY_JWT", "secret_key")
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		context      string
		user         storages.User
		params       url.Values
		expectedErr  errutil.MessageErr
		expectedCode int
	}{
		{
			context:      "success",
			params:       url.Values{"user_id": []string{"1000"}, "password": []string{"password"}},
			expectedCode: http.StatusOK,
		},
		{
			context:      "invalid user",
			params:       url.Values{"user_id": []string{"firstUser"}},
			expectedCode: http.StatusUnauthorized,
			expectedErr:  errutil.NewUnauthorizedError("record not found"),
		},
		{
			context:      "invalid password",
			params:       url.Values{"user_id": []string{"1000"}, "password": []string{"abcd"}},
			expectedCode: http.StatusUnauthorized,
			expectedErr:  errutil.NewUnauthorizedError("record not found"),
		},
	}

	for _, c := range testCases {

		t.Run(c.context, func(t *testing.T) {
			db := testfixture.SetupRepo(fixturePath)
			transport := transport.NewTransport(db)

			router := mux.NewRouter()
			router.HandleFunc("/login", transport.Login).Methods("GET")
			req, _ := http.NewRequest(http.MethodGet, "/login", nil)
			req.URL.RawQuery = c.params.Encode()

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, c.expectedCode, rr.Code)
			if c.expectedCode != http.StatusOK {
				message, err := errutil.NewApiErrFromBytes(rr.Body.Bytes())
				assert.NoError(t, err)
				assert.Equal(t, c.expectedErr, message)
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	testCases := []struct {
		context      string
		params       interface{}
		userID       interface{}
		expectedErr  errutil.MessageErr
		expectedCode int
		expected     storages.Task
	}{
		{
			context:      "success",
			params:       storages.Task{Content: "test task"},
			expectedCode: http.StatusOK,
			userID:       "1000",
			expected:     storages.Task{Content: "test task"},
		},
		{
			context:      "parse params error",
			params:       `example`,
			expectedCode: http.StatusBadRequest,
			expectedErr:  errutil.NewBadRequestError("json: cannot unmarshal string into Go value of type storages.Task"),
			userID:       "1000",
		},
		{
			context:      "user_id is empty",
			expectedCode: http.StatusBadRequest,
			expectedErr:  errutil.NewBadRequestError("user_id is empty"),
			userID:       nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.context, func(t *testing.T) {
			db := testfixture.SetupRepo(fixturePath)
			transport := transport.NewTransport(db)

			jsonStr, err := json.Marshal(c.params)
			assert.NoError(t, err)

			router := mux.NewRouter()
			router.HandleFunc("/tasks", transport.CreateTask).Methods("POST")

			req, err := http.NewRequestWithContext(context.WithValue(context.Background(), 0, c.userID), http.MethodPost, "/tasks", bytes.NewBuffer(jsonStr))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, c.expectedCode, rr.Code)

			if c.expectedCode == http.StatusOK {
				task := storages.Task{}
				err := json.Unmarshal(rr.Body.Bytes(), &task)
				assert.NoError(t, err)
				assert.Equal(t, c.expected.Content, task.Content)
			} else {
				message, err := errutil.NewApiErrFromBytes(rr.Body.Bytes())
				assert.NoError(t, err)
				assert.Equal(t, c.expectedErr, message)
			}
		})
	}
}

var (
	task1 = storages.Task{
		ID:      "10001000",
		UserID:  "1000",
		Content: "first task",
	}
)

func TestListTasks(t *testing.T) {
	testCases := []struct {
		context      string
		params       url.Values
		userID       interface{}
		expectedErr  errutil.MessageErr
		expectedCode int
		expect       []*storages.Task
	}{
		{
			context:      "success",
			expectedCode: http.StatusOK,
			userID:       "1000",
			expect:       []*storages.Task{&task1, &task1},
		},
		{
			context:      "created date invalid",
			expectedCode: http.StatusOK,
			params:       url.Values{"created_date": []string{"2021-07-27"}},
			userID:       "1000",
			expect:       []*storages.Task{},
		},
		{
			context:      "cannot get user_id",
			params:       url.Values{"created_date": []string{"2021-07-27"}},
			expectedCode: http.StatusBadRequest,
			expectedErr:  errutil.NewBadRequestError("user_id is empty"),
			userID:       nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.context, func(t *testing.T) {
			db := testfixture.SetupRepo(fixturePath)
			transport := transport.NewTransport(db)

			router := mux.NewRouter()
			router.HandleFunc("/tasks", transport.ListTasks).Methods("GET")

			req, err := http.NewRequestWithContext(context.WithValue(context.Background(), 0, c.userID), http.MethodGet, "/tasks", nil)
			assert.NoError(t, err)
			req.URL.RawQuery = c.params.Encode()

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, c.expectedCode, rr.Code)

			if c.expectedCode == http.StatusOK {
				tasks := []*storages.Task{}
				err := json.Unmarshal(rr.Body.Bytes(), &tasks)
				assert.NoError(t, err)
				assert.Equal(t, len(c.expect), len(tasks))
			} else {
				message, err := errutil.NewApiErrFromBytes(rr.Body.Bytes())
				assert.NoError(t, err)
				assert.Equal(t, c.expectedErr, message)
			}
		})
	}
}
