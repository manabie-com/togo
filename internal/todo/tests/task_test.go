package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	d "github.com/manabie-com/togo/internal/todo/domain"
	"github.com/stretchr/testify/assert"
)

func Test_createTask(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		if err := testDB.Truncate(); err != nil {
			t.Errorf("%+v", err)
		}
	}()

	users, err := testDB.SeedUser()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	user := users[0]

	tests := []struct {
		name         string
		requestBody  map[string]string
		authenicated bool
		wantErr      bool
		expectedCode int
	}{
		{
			"Empty content",
			map[string]string{"content": ""},
			true,
			true,
			http.StatusBadRequest,
		},
		{
			"Invalid content",
			map[string]string{"bcd": ""},
			true,
			true,
			http.StatusBadRequest,
		},
		{
			"Valid",
			map[string]string{"content": "Test 123"},
			true,
			false,
			http.StatusOK,
		},
		{
			"Not Login",
			map[string]string{"content": "Test 123"},
			false,
			false,
			http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := ""
			if tt.authenicated {
				token, err = getToken(user)
				if err != nil {
					t.Errorf("%+v", err)
				}
			}

			rBody, _ := json.Marshal(tt.requestBody)
			req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(rBody))
			if err != nil {
				t.Errorf("%+v", err)
			}
			req.Header.Set("Authorization", "BEARER "+token)

			w := httptest.NewRecorder()
			appHandler.ServeHTTP(w, req)

			if tt.authenicated {
				if tt.wantErr {
					var m map[string]string
					if err := json.NewDecoder(w.Body).Decode(&m); err != nil {
						t.Errorf("%+v", err)
					}
					val, ok := m["error"]
					assert.True(ok)
					assert.NotEmpty(val)
				} else {
					var m map[string]*d.Task
					if err := json.NewDecoder(w.Body).Decode(&m); err != nil {
						t.Errorf("%+v", err)
					}
					val, ok := m["data"]
					assert.True(ok)
					assert.NotEmpty(val)
				}
			}

			assert.Equal(tt.expectedCode, w.Code)
		})
	}
}

func Test_listTasks(t *testing.T) {
	assert := assert.New(t)
	defer func() {
		if err := testDB.Truncate(); err != nil {
			t.Errorf("%+v", err)
		}
	}()

	users, err := testDB.SeedUser()
	if err != nil {
		t.Fatalf("%+v", err)
	}
	user := users[0]
	_, err = testDB.SeedTask()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	tests := []struct {
		name         string
		dateStr      string
		authenicated bool
		wantErr      bool
		expectedCode int
		numData      int
	}{
		{
			"Not Login",
			"2021-03-20",
			false,
			false,
			http.StatusUnauthorized,
			0,
		},
		{
			"Invalid date",
			"abcdef",
			true,
			false,
			http.StatusOK,
			1,
		},
		{
			"Empty date",
			"",
			true,
			false,
			http.StatusOK,
			1,
		},
		{
			"Date No data",
			"2019-09-09",
			true,
			false,
			http.StatusOK,
			0,
		},
		{
			"Date 1",
			"2021-03-15",
			true,
			false,
			http.StatusOK,
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := ""
			if tt.authenicated {
				token, err = getToken(user)
				if err != nil {
					t.Errorf("%+v", err)
				}
			}

			req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
			if err != nil {
				t.Errorf("%+v", err)
			}
			q := req.URL.Query()
			q.Add("created_date", tt.dateStr)
			req.URL.RawQuery = q.Encode()
			req.Header.Set("Authorization", "BEARER "+token)

			w := httptest.NewRecorder()
			appHandler.ServeHTTP(w, req)

			if tt.authenicated {
				if tt.wantErr {
					var m map[string]string
					if err := json.NewDecoder(w.Body).Decode(&m); err != nil {
						t.Errorf("%+v", err)
					}
					val, ok := m["error"]
					assert.True(ok)
					assert.NotEmpty(val)
				} else {
					var m map[string][]*d.Task
					if err := json.NewDecoder(w.Body).Decode(&m); err != nil {
						t.Errorf("%+v", err)
					}
					val, ok := m["data"]
					assert.Equal(tt.numData, len(val))
					assert.True(ok)
				}
			}

			assert.Equal(tt.expectedCode, w.Code)
		})
	}
}

func getToken(user *d.User) (string, error) {
	authParam := map[string]string{"username": user.Username, "password": user.Password}
	rBody, _ := json.Marshal(authParam)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(rBody))
	if err != nil {
		return "", err
	}
	w := httptest.NewRecorder()
	appHandler.ServeHTTP(w, req)

	var m map[string]string
	if err := json.NewDecoder(w.Body).Decode(&m); err != nil {
		return "", err
	}

	return m["data"], nil
}
