package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/stretchr/testify/require"
)

var (
	ts          *httptest.Server
	accessToken string
	client      *http.Client

	testUser        = storages.User{ID: "firstUser", Password: "example"}
	testUserMaxTodo = 5
	testCreatedDate = "2020-06-29"
	testTasks       = []*storages.Task{
		{ID: "e1da0b9b-7ecc-44f9-82ff-4623cc50446a", Content: "first content", UserID: testUser.ID, CreatedDate: testCreatedDate},
		{ID: "055261ab-8ba8-49e1-a9e8-e9f725ba9104", Content: "second content", UserID: testUser.ID, CreatedDate: testCreatedDate},
		{ID: "2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a", Content: "another content", UserID: testUser.ID, CreatedDate: testCreatedDate},
	}
	testNewContent = "today content"
)

func TestMain(m *testing.M) {
	db, err := sql.Open("pgx", sqllite.DBConnectionURL())
	if err != nil {
		log.Fatal("error opening db: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("error pinging: ", err)
	}
	defer db.Close()

	ts = httptest.NewServer(&services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})
	defer ts.Close()

	client = &http.Client{}

	log.Println("Finished initializing. Running test...")
	os.Exit(m.Run())
}

func TestLogin_Success(t *testing.T) {
	require := require.New(t)
	jsonResp := makeLoginRequest(require, testUser.ID, testUser.Password)
	require.Equal(http.StatusOK, jsonResp.StatusCode)
	require.NotEmpty(jsonResp.Data)
	require.Empty(jsonResp.Error)

	// Save the access token for authorized requests
	accessToken = jsonResp.Data
}

func TestLogin_WrongCredentials(t *testing.T) {
	testCases := []storages.User{
		{ID: "wrongUserID", Password: testUser.Password},
		{ID: testUser.ID, Password: "wrongPassword"},
	}
	require := require.New(t)
	for _, u := range testCases {
		jsonResp := makeLoginRequest(require, u.ID, u.Password)
		require.Equal(http.StatusUnauthorized, jsonResp.StatusCode)
		require.NotEmpty(jsonResp.Error)
		require.Empty(jsonResp.Data)
	}
}

func TestListTasks_Success(t *testing.T) {
	require := require.New(t)
	jsonResp := makeListTasksRequest(require, testCreatedDate)
	require.Equal(http.StatusOK, jsonResp.StatusCode)
	require.ElementsMatch(testTasks, jsonResp.Data)
	require.Empty(jsonResp.Error)
}

func TestListTasks_SuccessEmpty(t *testing.T) {
	emptyTaskDate := "2020-06-30"
	require := require.New(t)
	jsonResp := makeListTasksRequest(require, emptyTaskDate)
	require.Equal(http.StatusOK, jsonResp.StatusCode)
	require.ElementsMatch([]*storages.Task{}, jsonResp.Data)
	require.Empty(jsonResp.Error)
}

func TestListTasks_Unauthorized(t *testing.T) {
	require := require.New(t)
	// Create HTTP request
	req, err := http.NewRequest("GET", ts.URL+"/tasks", nil)
	require.NoError(err)
	q := req.URL.Query()
	q.Add("created_date", testCreatedDate)
	req.URL.RawQuery = q.Encode()

	// Perform the request and parse the response
	resp, err := client.Do(req)
	require.NoError(err)
	defer resp.Body.Close()

	require.Equal(http.StatusUnauthorized, resp.StatusCode)
}

// TestAddTasks assumes that there are no task entries created today
func TestAddTasks(t *testing.T) {
	require := require.New(t)
	for i := 0; i < testUserMaxTodo; i++ {
		jsonResp := makeAddTaskRequest(require, testNewContent)
		require.Equal(http.StatusOK, jsonResp.StatusCode)
		require.Equal(testNewContent, jsonResp.Data.Content)
		require.Empty(jsonResp.Error)
	}

	// The following add task request will fail due to exceeding quota limit
	jsonResp := makeAddTaskRequest(require, testNewContent)
	require.Equal(http.StatusTooManyRequests, jsonResp.StatusCode)
	require.Empty(jsonResp.Data)
	require.NotEmpty(jsonResp.Error)
}

func makeLoginRequest(require *require.Assertions, userID, password string) *storages.LoginJSONResponse {
	resp, err := http.PostForm(
		ts.URL+"/login",
		url.Values{
			"user_id":  {userID},
			"password": {password},
		},
	)
	require.NoError(err)
	defer resp.Body.Close()

	var jsonResp *storages.LoginJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	jsonResp.StatusCode = resp.StatusCode
	return jsonResp
}

// makeListTasksRequest uses the same token from the successful login request
func makeListTasksRequest(require *require.Assertions, createdDate string) *storages.ListTasksJSONResponse {
	// Create HTTP request
	req, err := http.NewRequest("GET", ts.URL+"/tasks", nil)
	require.NoError(err)
	q := req.URL.Query()
	q.Add("created_date", createdDate)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", accessToken)

	// Perform the request and parse the response
	resp, err := client.Do(req)
	require.NoError(err)
	defer resp.Body.Close()

	var jsonResp *storages.ListTasksJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	jsonResp.StatusCode = resp.StatusCode
	return jsonResp
}

func makeAddTaskRequest(require *require.Assertions, content string) *storages.AddTasksJSONResponse {
	// Create HTTP request
	payload, err := json.Marshal(struct {
		Content string `json:"content"`
	}{
		Content: content,
	})
	require.NoError(err)
	req, err := http.NewRequest("POST", ts.URL+"/tasks", bytes.NewBuffer(payload))
	require.NoError(err)
	req.Header.Add("Authorization", accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Perform the request and parse the response
	resp, err := client.Do(req)
	require.NoError(err)
	defer resp.Body.Close()

	var jsonResp *storages.AddTasksJSONResponse
	require.NoError(json.NewDecoder(resp.Body).Decode(&jsonResp))
	jsonResp.StatusCode = resp.StatusCode
	return jsonResp
}
