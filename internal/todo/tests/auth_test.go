package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_login(t *testing.T) {
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
		wantErr      bool
		expectedCode int
	}{
		{
			"Invalid",
			map[string]string{"username": "fail", "password": "fail"},
			true,
			http.StatusBadRequest,
		},
		{
			"Valid",
			map[string]string{"username": user.Username, "password": user.Password},
			false,
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rBody, _ := json.Marshal(tt.requestBody)
			req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(rBody))

			if err != nil {
				t.Errorf("%+v", err)
			}

			w := httptest.NewRecorder()
			appHandler.ServeHTTP(w, req)

			var m map[string]string
			if err := json.NewDecoder(w.Body).Decode(&m); err != nil {
				t.Errorf("%+v", err)
			}

			if tt.wantErr {
				val, ok := m["error"]
				assert.True(ok)
				assert.NotEmpty(val)
			} else {
				val, ok := m["data"]
				assert.True(ok)
				assert.NotEmpty(val)
			}

			assert.Equal(tt.expectedCode, w.Code)
		})
	}
}
