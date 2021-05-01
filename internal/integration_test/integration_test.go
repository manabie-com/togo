package integration_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/sqldb"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/stretchr/testify/assert"
)

type tokenResponse struct {
	Token string `json:"data"`
}

type getTasksResponse struct {
	Tasks []*storages.Task `json:"data"`
}

type addTasksRequest struct {
	Content string `json:"content"`
}

type addTasksResponse struct {
	Task *storages.Task `json:"data"`
}

const (
	testUserId  = "testUser"
	testUserPwd = "testPwd"
)

type TestConfig struct {
	StorageConfig storages.Config `json:"storages"`
}

var testConfig TestConfig

func TestLogin(t *testing.T) {
	jwtToken := "wqGyEBBfPK9w3Lxw"
	dbPool, err := openDbForTest()
	assert.Nil(t, err)
	ctx := context.Background()
	err = addTestUser(dbPool, ctx)
	assert.Nil(t, err)
	defer func() {
		_ = removeTestUser(dbPool, ctx)
		dbPool.Close()
	}()

	storeRepo := &sqldb.SqlDB{
		DB: dbPool,
	}
	todoService := services.NewToDoService(storeRepo)
	httpHandler := transport.NewHttpHandler(jwtToken, todoService)

	token, _ := login(httpHandler, "someUserNotInDb9739dgetw3t3ifheh", "somePwdNotInDb9rfhsfh93yrhe")
	assert.Empty(t, token)

	token, _ = login(httpHandler, testUserId, "somePwdNotInDb9rfhsfh93yrhe")
	assert.Empty(t, token)

	token, _ = login(httpHandler, "someUserNotInDb9739dgetw3t3ifheh", testUserPwd)
	assert.Empty(t, token)

	token, err = login(httpHandler, testUserId, testUserPwd)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestAddAndGetTasks(t *testing.T) {
	jwtToken := "wqGyEBBfPK9w3Lxw"
	dbPool, err := openDbForTest()
	assert.Nil(t, err)
	ctx := context.Background()
	err = addTestUser(dbPool, ctx)
	assert.Nil(t, err)
	defer func() {
		_ = removeTasksByTestUser(dbPool, ctx)
		_ = removeTestUser(dbPool, ctx)
		dbPool.Close()
	}()

	storeRepo := &sqldb.SqlDB{
		DB: dbPool,
	}
	todoService := services.NewToDoService(storeRepo)
	httpHandler := transport.NewHttpHandler(jwtToken, todoService)

	statusCode, tasks, err := getTasks(httpHandler, "")
	assert.EqualValues(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, tasks)
	assert.NotNil(t, err)

	token, _ := login(httpHandler, testUserId, testUserPwd)
	assert.NotEmpty(t, token)

	req, err := json.Marshal(addTasksRequest{Content: "c1"})
	assert.Nil(t, err)

	statusCode, task, err := addTask(httpHandler, "", req)
	assert.EqualValues(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, task)
	assert.NotNil(t, err)

	statusCode, task, err = addTask(httpHandler, token, req)
	assert.EqualValues(t, http.StatusOK, statusCode)
	assert.NotNil(t, task)
	assert.Nil(t, err)

	req, err = json.Marshal(addTasksRequest{Content: "c2"})
	assert.Nil(t, err)

	statusCode, task, err = addTask(httpHandler, token, req)
	assert.EqualValues(t, http.StatusOK, statusCode)
	assert.NotNil(t, task)
	assert.Nil(t, err)

	statusCode, tasks, err = getTasks(httpHandler, token)
	assert.EqualValues(t, http.StatusOK, statusCode)
	assert.Nil(t, err)
	assert.NotNil(t, tasks)
	assert.EqualValues(t, 2, len(tasks))
	assert.EqualValues(t, "c1", tasks[0].Content)
	assert.EqualValues(t, testUserId, tasks[0].UserID)
	assert.EqualValues(t, time.Now().Format("2006-01-02"), tasks[0].CreatedDate)
	assert.EqualValues(t, "c2", tasks[1].Content)
	assert.EqualValues(t, testUserId, tasks[1].UserID)
	assert.EqualValues(t, time.Now().Format("2006-01-02"), tasks[1].CreatedDate)

	statusCode, task, err = addTask(httpHandler, token, req)
	assert.EqualValues(t, http.StatusOK, statusCode)
	assert.NotNil(t, task)
	assert.Nil(t, err)

	statusCode, task, err = addTask(httpHandler, token, req)
	assert.EqualValues(t, http.StatusForbidden, statusCode)
	assert.Nil(t, task)
}

func login(httpHandler http.Handler, userID, password string) (string, error) {
	request, _ := http.NewRequest(http.MethodGet, "/login?user_id="+userID+"&password="+password, nil)
	response := httptest.NewRecorder()
	httpHandler.ServeHTTP(response, request)
	tokenResp := tokenResponse{}
	err := json.NewDecoder(response.Body).Decode(&tokenResp)
	if err != nil {
		return "", err
	}
	return tokenResp.Token, nil
}

func getTasks(httpHandler http.Handler, token string) (int, []*storages.Task, error) {
	request, _ := http.NewRequest(http.MethodGet, "/tasks?created_date="+time.Now().Format("2006-01-02"), nil)
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	httpHandler.ServeHTTP(response, request)
	getTasksResp := getTasksResponse{}
	err := json.NewDecoder(response.Body).Decode(&getTasksResp)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return response.Result().StatusCode, getTasksResp.Tasks, nil
}

func addTask(httpHandler http.Handler, token string, req []byte) (int, *storages.Task, error) {
	request, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(req))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	httpHandler.ServeHTTP(response, request)
	addTasksResp := addTasksResponse{}
	err := json.NewDecoder(response.Body).Decode(&addTasksResp)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return response.Result().StatusCode, addTasksResp.Task, nil
}

func openDbForTest() (*sql.DB, error) {
	config.Load("./test_config.json", &testConfig)
	postgresConfig := testConfig.StorageConfig.Postgres
	postgresConnStr, err := postgresConfig.Build()
	if err != nil {
		return nil, err
	}
	dbPool, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		return nil, err
	}
	if postgresConfig.MaxIdleConns > 0 {
		dbPool.SetMaxIdleConns(postgresConfig.MaxIdleConns)
	}
	return dbPool, nil
}

func addTestUser(dbConn *sql.DB, ctx context.Context) error {
	stmt := `INSERT INTO users (id, password, max_todo) VALUES('` + testUserId + `', '` + testUserPwd + `', 3);`
	_, err := dbConn.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}

func removeTestUser(dbConn *sql.DB, ctx context.Context) error {
	stmt := `DELETE FROM users WHERE id ='` + testUserId + `';`
	_, err := dbConn.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}

func removeTasksByTestUser(dbConn *sql.DB, ctx context.Context) error {
	stmt := `DELETE FROM tasks WHERE user_id ='` + testUserId + `';`
	_, err := dbConn.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}
