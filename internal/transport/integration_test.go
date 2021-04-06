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

var (
	server = getServer()
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
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

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
			userId:   "secondUser",
			password: "example",
			content:  "test add success",
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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

			token := login(t, tc.userId, tc.password, server)

			path := fmt.Sprintf("/tasks")
			params := createTaskParams{
				Content: tc.content,
			}
			recorder := makeRequest(t, server, http.MethodPost, path, token, &params)

			tc.check(t, recorder)
		})
	}

}

func TestListTask(t *testing.T) {

	testCases := []struct {
		name        string
		userId      string
		password    string
		createdDate string
		total       int
		page        int
		check       func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "Ok",
			userId:      "secondUser",
			password:    "example",
			createdDate: util.GetDate(),
			total:       10,
			page:        1,
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:        "BadRequest: createdDate is invalid",
			userId:      "secondUser",
			password:    "example",
			createdDate: "",
			total:       10,
			page:        1,
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:        "BadRequest: total is negative",
			userId:      "secondUser",
			password:    "example",
			createdDate: util.GetDate(),
			total:       -1,
			page:        1,
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
		{
			name:        "BadRequest: page is negative",
			userId:      "secondUser",
			password:    "example",
			createdDate: util.GetDate(),
			total:       10,
			page:        -1,
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			token := login(t, tc.userId, tc.password, server)

			path := fmt.Sprintf("/tasks/%v/%v/%v", tc.createdDate, tc.total, tc.page)
			fmt.Println("path: ", path)
			recorder := makeRequest(t, server, http.MethodGet, path, token, nil)

			tc.check(t, recorder)
		})
	}
}

func TestRatelimit(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/login"), nil)
	req.Header.Add("X-Forwarded-For", "100.100.100.100")
	require.NoError(t, err)
	require.NotEmpty(t, req)
	code := http.StatusOK
	for i := 0; i < 10; i++ {
		recorder := httptest.NewRecorder()
		server.httpServer.Handler.ServeHTTP(recorder, req)
		if i == 9 {
			code = recorder.Code
		}
	}

	require.Equal(t, http.StatusTooManyRequests, code)
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

func getServer() *Server {
	util.LoadConfig("../../configs")
	testStore := postgres.NewPostgres()
	server := NewServer(testStore)

	return server
}

func makeRequest(t *testing.T, s *Server, method, path, token string, body interface{}) *httptest.ResponseRecorder {
	bodyBytes, err := json.Marshal(body)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, bytes.NewBuffer([]byte(bodyBytes)))
	req.Header.Add(authorizationHeaderKey, "bearer "+token)
	require.NoError(t, err)
	s.httpServer.Handler.ServeHTTP(recorder, req)

	return recorder
}
