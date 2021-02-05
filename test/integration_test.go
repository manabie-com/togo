package test

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/model"
	"github.com/manabie-com/togo/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func initSQLite(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		t.Errorf("error opening db: %s", err.Error())
		t.FailNow()
		return nil
	}

	return db
}

func setupTestServerWithSQLite(t *testing.T) *http.ServeMux {
	db := initSQLite(t)

	mux := http.NewServeMux()
	mux.Handle("/", &services.ToDoService{
		JWTKey: os.Getenv("ENCRYPTION"),
		Store: &sqllite.LiteDB{
			DB: db,
		},
	})

	return mux
}

func initPostgres(t *testing.T) *sql.DB {
	pgInfo := config.GetConfig().GetConnString()
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		t.Errorf("error opening db: %s", err.Error())
		t.FailNow()
		return nil
	}

	err = db.Ping()
	if err != nil {
		t.Errorf("error opening db: %s", err.Error())
		t.FailNow()
		return nil
	}

	return db
}

func setupTestServerWithPostgres(t *testing.T) *http.ServeMux {
	connStr := config.GetConfig().GetConnString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return nil
	}

	mux := http.NewServeMux()
	mux.Handle("/", &services.ToDoService{
		JWTKey: os.Getenv("ENCRYPTION"),
		Store:  &postgres.Postgres{DB: db},
	})

	return mux
}

func makeReqLogin(t *testing.T, mux *http.ServeMux, input *model.LoginInput) *httptest.ResponseRecorder {
	params := utils.ToMap(input)
	writer := httptest.NewRecorder()
	request := utils.InitHTTPRequest("GET", "/login", nil, params, nil)
	mux.ServeHTTP(writer, request)
	return writer
}

func getLoginResponse(t *testing.T, writer *httptest.ResponseRecorder) *model.LoginResponse {
	b, err := ioutil.ReadAll(writer.Body)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return nil
	}

	var response *model.LoginResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return nil
	}

	if response.Data == "" {
		t.Errorf("Expected output contains value but it does not\n")
		t.FailNow()
		return nil
	}

	t.Logf("Response contains value as expected\n")
	return response
}

func getServiceToken(t *testing.T, id string, d time.Duration) string {
	s := &services.ToDoService{JWTKey: os.Getenv("ENCRYPTION")}
	token, err := s.CreateTokenWithExpireTime(id, d)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return ""
	}

	return token
}

func makeReqGetTasks(t *testing.T, mux *http.ServeMux, input *model.GetTaskInput, authorization string) *httptest.ResponseRecorder {
	params := utils.ToMap(input)
	headers := map[string]string{
		"Authorization": authorization,
	}
	writer := httptest.NewRecorder()
	request := utils.InitHTTPRequest("GET", "/tasks", headers, params, nil)
	mux.ServeHTTP(writer, request)
	return writer
}

func getListTasks(t *testing.T, writer *httptest.ResponseRecorder) *model.GetTaskResponse {
	b, err := ioutil.ReadAll(writer.Body)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return nil
	}

	var output *model.GetTaskResponse
	err = json.Unmarshal(b, &output)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return nil
	}

	return output
}

func makeReqCreateTasks(t *testing.T, mux *http.ServeMux, input *model.CreateTaskInput, authorization string) *httptest.ResponseRecorder {
	headers := map[string]string{
		"Authorization": authorization,
	}
	writer := httptest.NewRecorder()
	request := utils.InitHTTPRequest("POST", "/tasks", headers, nil, input)
	mux.ServeHTTP(writer, request)
	return writer
}

func getCreateTaskResponse(t *testing.T, writer *httptest.ResponseRecorder) *model.CreateTaskResponse {
	b, err := ioutil.ReadAll(writer.Body)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return nil
	}

	var output *model.CreateTaskResponse
	err = json.Unmarshal(b, &output)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return nil
	}

	return output
}

func testLoginSuccessfully(t *testing.T, input *model.LoginInput) *model.LoginResponse {
	var (
		expectedCode = http.StatusOK
		mux          = setupTestServerWithPostgres(t)
	)

	writer := makeReqLogin(t, mux, input)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return nil
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)

	return getLoginResponse(t, writer)
}

func TestDBConnected(t *testing.T) {
	db := initPostgres(t)
	pg := &postgres.Postgres{
		DB: db,
	}
	tasks, err := pg.RetrieveTasks(context.Background(),
		sql.NullString{
			String: "firstUser",
			Valid:  true,
		},
		sql.NullString{
			String: "2020-06-29",
			Valid:  true,
		},
	)
	if err != nil {
		t.Error(err)
	}
	if tasks == nil || len(tasks) == 0 {
		t.Errorf("RetrieveTasks does not return data as expected")
		t.FailNow()
	}
}

func TestMethodOptions(t *testing.T) {
	var (
		expectedCode = http.StatusOK
		mux          = setupTestServerWithPostgres(t)
	)

	writer := httptest.NewRecorder()
	request := utils.InitHTTPRequest("OPTIONS", "/", nil, nil, nil)
	mux.ServeHTTP(writer, request)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Response code is %v", writer.Code)
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)
}

func TestLoginFailed(t *testing.T) {
	var (
		expectedCode = http.StatusUnauthorized
		mux          = setupTestServerWithPostgres(t)
	)

	input := &model.LoginInput{
		UserID:   "firstUser",
		Password: "incorrect-password",
	}
	writer := makeReqLogin(t, mux, input)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)
}

func TestLoginSuccessfully(t *testing.T) {
	testLoginSuccessfully(t, &model.LoginInput{
		UserID:   "firstUser",
		Password: "example",
	})
}

func TestGetTasksWithInvalidToken(t *testing.T) {
	var (
		expectedCode = http.StatusUnauthorized
		mux          = setupTestServerWithPostgres(t)
	)

	input := &model.GetTaskInput{
		CreatedDate: time.Now().Format("2006-01-02"),
	}
	writer := makeReqGetTasks(t, mux, input, "this is an invalid token")

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)
}

func TestGetTasksWithExpiredToken(t *testing.T) {
	var (
		expectedCode = http.StatusUnauthorized
		expiredTime  = time.Minute
		mux          = setupTestServerWithPostgres(t)
	)

	token := getServiceToken(t, "firstUser", expiredTime)

	// Wait to let the token expired
	time.Sleep(expiredTime + time.Minute)

	input := &model.GetTaskInput{
		CreatedDate: time.Now().Format("2006-01-02"),
	}
	writer := makeReqGetTasks(t, mux, input, token)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)
}

func TestGetTasksWithValidToken(t *testing.T) {
	var (
		expectedCode = http.StatusOK
		expiredTime  = 15 * time.Minute
		createdDate  = "2020-06-29"
		mux          = setupTestServerWithPostgres(t)
	)

	token := getServiceToken(t, "firstUser", expiredTime)

	input := &model.GetTaskInput{
		CreatedDate: createdDate,
	}
	writer := makeReqGetTasks(t, mux, input, token)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)

	response := getListTasks(t, writer)
	if response.Data == nil || len(response.Data) == 0 {
		t.Errorf("Expectation: data is not empty data but it's not\n")
		t.FailNow()
		return
	}

	t.Logf("Response data is a list of tasks as expected\n")
}

func TestGetTasksWithValidTokenButEmpty(t *testing.T) {
	var (
		expectedCode = http.StatusOK
		expiredTime  = 15 * time.Minute
		createdDate  = "2021-02-30"
		mux          = setupTestServerWithPostgres(t)
	)

	token := getServiceToken(t, "thirdUser", expiredTime)

	input := &model.GetTaskInput{
		CreatedDate: createdDate,
	}
	writer := makeReqGetTasks(t, mux, input, token)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)

	response := getListTasks(t, writer)
	if response.Data != nil && len(response.Data) > 0 {
		t.Errorf("Expectation: data is empty but it's not\n")
		t.FailNow()
		return
	}

	t.Logf("Response data is empty as expected\n")
}

func TestCreateTheTaskWithInvalidToken(t *testing.T) {
	var (
		expectedCode = http.StatusUnauthorized
		mux          = setupTestServerWithPostgres(t)
	)

	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	writer := makeReqCreateTasks(t, mux, input, "this is an invalid token")

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)
}

func TestCreateTheTaskWithExpiredToken(t *testing.T) {
	var (
		expectedCode = http.StatusUnauthorized
		expiredTime  = time.Minute
		mux          = setupTestServerWithPostgres(t)
	)

	token := getServiceToken(t, "firstUser", expiredTime)

	// Wait to let the token expired
	time.Sleep(expiredTime + time.Minute)

	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	writer := makeReqCreateTasks(t, mux, input, token)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)
}

func TestCreateTheTaskWithValidToken(t *testing.T) {
	var (
		expectedCode = http.StatusOK
		mux          = setupTestServerWithPostgres(t)
	)

	loginRes := testLoginSuccessfully(t, &model.LoginInput{
		UserID:   "firstUser",
		Password: "example",
	})

	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	writer := makeReqCreateTasks(t, mux, input, loginRes.Data)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)

	response := getCreateTaskResponse(t, writer)
	if response.Data == nil {
		t.Errorf("Expectation: data is not empty data but it's not\n")
		t.FailNow()
		return
	}

	t.Logf("New task has been created as expected\n")
}

func TestCreateTheTaskWithValidTokenButNotAccepted(t *testing.T) {
	var (
		expectedCode = http.StatusBadRequest
		mux          = setupTestServerWithPostgres(t)
	)

	loginRes := testLoginSuccessfully(t, &model.LoginInput{
		UserID:   "secondUser",
		Password: "example",
	})

	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	writer := makeReqCreateTasks(t, mux, input, loginRes.Data)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)

	response := getCreateTaskResponse(t, writer)
	if response.Data != nil {
		t.Errorf("Expectation: data is empty but it's not\n")
		t.FailNow()
		return
	}
	if response.Error == "" {
		t.Errorf("Expectation: response contains error but it does not\n")
		t.FailNow()
		return
	}
	t.Logf("Server does not accept as expected\n")
}

func TestCreateTheTasksToReachLimitWithValidToken(t *testing.T) {
	var (
		expectedCode int
		mux          = setupTestServerWithPostgres(t)
	)

	loginRes := testLoginSuccessfully(t, &model.LoginInput{
		UserID:   "thirdUser",
		Password: "example",
	})

	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	for i := 1; i <= 6; i++ {
		writer := makeReqCreateTasks(t, mux, input, loginRes.Data)

		// Validate result
		if i <= 5 {
			expectedCode = http.StatusOK
		} else {
			expectedCode = http.StatusBadRequest
		}

		if writer.Code != expectedCode {
			t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
			t.FailNow()
			return
		}

		t.Logf("Server responses status code %d as expected\n", expectedCode)

		response := getCreateTaskResponse(t, writer)
		if i <= 5 {
			if response.Data == nil {
				t.Errorf("Expectation: data is not empty data but it's not\n")
				t.FailNow()
				return
			}

			t.Logf("New task has been created with id '%s'\n", response.Data.ID)

			time.Sleep(10 * time.Second)
		} else {
			if response.Data != nil {
				t.Errorf("Expectation: data is empty but it's not\n")
				t.FailNow()
				return
			}
			if response.Error == "" {
				t.Errorf("Expectation: response contains error but it does not\n")
				t.FailNow()
				return
			}
			t.Logf("Server does not accept as expected\n")
		}
	}
}
