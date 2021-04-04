package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/util"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	testCases := []struct {
		name     string
		userId   string
		password string
		check    func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "Ok",
			userId:   "firstUser",
			password: "example",
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.NotEmpty(t, recorder.Body)
			},
		},
		{
			name:     "BadRequest: dont have id",
			userId:   "",
			password: "example",
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:     "BadRequest: password is incorrect",
			userId:   "firstUser",
			password: "example123",
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := getServer(t)

			// make request
			path := fmt.Sprintf("/login")
			param := loginParams{
				Id:       tc.userId,
				Password: tc.password,
			}

			recorder := makeRequest(t, server, http.MethodPost, path, "", param)
			tc.check(t, recorder)
		})
	}
}

// TestAddTask is function integration test
// We have to login success then we can create a task
func TestAddTask(t *testing.T) {
	testCases := []struct {
		name     string
		userId   string
		password string
		content  string
		check    func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "Ok",
			userId:   "firstUser",
			password: "example",
			content:  "test add success",
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:     "BadRequest: Content is empty",
			userId:   "firstUser",
			password: "example",
			content:  "",
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},

		{
			name:     "BadRequest: maxtodo",
			userId:   "fourthUser",
			password: "example",
			content:  "test",
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := getServer(t)

			token := login(t, tc.userId, tc.password, server)

			path := fmt.Sprintf("/tasks")
			params := createTaskParams{
				Content: tc.content,
			}
			fmt.Println("params: ", params)
			recorder := makeRequest(t, server, http.MethodPost, path, token, &params)
			fmt.Println("recorder: ", recorder)

			tc.check(t, recorder)
		})
	}

}

func login(t *testing.T, userId, password string, s *Server) string {
	path := fmt.Sprintf("/login")
	param := loginParams{
		Id:       userId,
		Password: password,
	}

	recorder := makeRequest(t, s, http.MethodPost, path, "", param)
	require.Equal(t, http.StatusOK, recorder.Code)

	bodyBytes, err := ioutil.ReadAll(recorder.Body)
	require.NoError(t, err)
	require.NotEmpty(t, bodyBytes)

	var mapDate map[string]interface{}
	err = json.Unmarshal(bodyBytes, &mapDate)
	require.NoError(t, err)

	token, ok := mapDate["data"]
	require.Equal(t, true, ok)

	return token.(string)
}

func getServer(t *testing.T) *Server {
	err := util.LoadConfig("../../configs")
	require.NoError(t, err)

	testStore := postgres.NewPostgres()
	server := NewServer(testStore)

	return server
}

func makeRequest(t *testing.T, s *Server, method, path, token string, body interface{}) *httptest.ResponseRecorder {
	bodyBytes, err := json.Marshal(body)
	require.NoError(t, err)
	require.NotEmpty(t, body)
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, bytes.NewBuffer([]byte(bodyBytes)))
	req.Header.Add(authorizationHeaderKey, "bearer "+token)
	require.NoError(t, err)
	s.router.ServeHTTP(recorder, req)

	return recorder
}
