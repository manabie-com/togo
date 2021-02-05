package services

import (
	"encoding/json"
	mockup_db "github.com/manabie-com/togo/internal/storages/mockup-db"
	"github.com/manabie-com/togo/model"
	"github.com/manabie-com/togo/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func setupMockupServer() *ToDoService {
	return &ToDoService{
		JWTKey: os.Getenv("ENCRYPTION"),
		Store:  &mockup_db.MockupDB{},
	}
}

func makeReqLogin(t *testing.T, s *ToDoService, input *model.LoginInput) *httptest.ResponseRecorder {
	params := utils.ToMap(input)
	writer := httptest.NewRecorder()
	request := utils.InitHTTPRequest("GET", "/login", nil, params, nil)
	s.ServeHTTP(writer, request)
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
	s := &ToDoService{JWTKey: os.Getenv("ENCRYPTION")}
	token, err := s.CreateTokenWithExpireTime(id, d)
	if err != nil {
		t.Error(err)
		t.FailNow()
		return ""
	}

	return token
}

func makeReqGetTasks(t *testing.T, s *ToDoService, input *model.GetTaskInput, authorization string) *httptest.ResponseRecorder {
	params := utils.ToMap(input)
	headers := map[string]string{
		"Authorization": authorization,
	}
	writer := httptest.NewRecorder()
	request := utils.InitHTTPRequest("GET", "/tasks", headers, params, nil)
	s.ServeHTTP(writer, request)
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

func makeReqCreateTasks(t *testing.T, s *ToDoService, input *model.CreateTaskInput, authorization string) *httptest.ResponseRecorder {
	headers := map[string]string{
		"Authorization": authorization,
	}
	writer := httptest.NewRecorder()
	request := utils.InitHTTPRequest("POST", "/tasks", headers, nil, input)
	s.ServeHTTP(writer, request)
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

func TestMethodOptions(t *testing.T) {
	var (
		expectedCode = http.StatusOK
		s            = setupMockupServer()
	)

	writer := httptest.NewRecorder()
	request := utils.InitHTTPRequest("OPTIONS", "/", nil, nil, nil)
	s.ServeHTTP(writer, request)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Response code is %v", writer.Code)
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)
}

func TestLoginFailed(t *testing.T) {
	var (
		expectedCode = http.StatusUnauthorized
		s            = setupMockupServer()
	)

	writer := httptest.NewRecorder()
	input := &model.LoginInput{
		UserID:   "firstUser",
		Password: "incorrect-password",
	}
	params := utils.ToMap(input)
	req := utils.InitHTTPRequest("GET", "/login", nil, params, nil)
	s.ServeHTTP(writer, req)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)
}

func TestLoginSuccessfully(t *testing.T) {
	var (
		expectedCode = http.StatusOK
		s            = setupMockupServer()
	)

	input := &model.LoginInput{
		UserID:   "firstUser",
		Password: "example",
	}
	writer := makeReqLogin(t, s, input)

	// Validate result
	if writer.Code != expectedCode {
		t.Errorf("Expectation: status is %d but server response with status %d\n", expectedCode, writer.Code)
		t.FailNow()
		return
	}

	t.Logf("Server responses status code %d as expected\n", expectedCode)

	getLoginResponse(t, writer)
}

func TestGetTasksWithInvalidToken(t *testing.T) {
	var (
		expectedCode = http.StatusUnauthorized
		s            = setupMockupServer()
	)

	input := &model.GetTaskInput{
		CreatedDate: time.Now().Format("2006-01-02"),
	}
	writer := makeReqGetTasks(t, s, input, "this is an invalid token")

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
		s            = setupMockupServer()
	)

	token := getServiceToken(t, "firstUser", expiredTime)

	// Wait to let the token expired
	time.Sleep(expiredTime + time.Minute)

	input := &model.GetTaskInput{
		CreatedDate: time.Now().Format("2006-01-02"),
	}
	writer := makeReqGetTasks(t, s, input, token)

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
		s            = setupMockupServer()
	)

	token := getServiceToken(t, "firstUser", expiredTime)

	input := &model.GetTaskInput{
		CreatedDate: createdDate,
	}
	writer := makeReqGetTasks(t, s, input, token)

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
		s            = setupMockupServer()
	)

	token := getServiceToken(t, "thirdUser", expiredTime)

	input := &model.GetTaskInput{
		CreatedDate: createdDate,
	}
	writer := makeReqGetTasks(t, s, input, token)

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
		s            = setupMockupServer()
	)

	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	writer := makeReqCreateTasks(t, s, input, "this is an invalid token")

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
		s            = setupMockupServer()
	)

	token := getServiceToken(t, "firstUser", expiredTime)

	// Wait to let the token expired
	time.Sleep(expiredTime + time.Minute)

	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	writer := makeReqCreateTasks(t, s, input, token)

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
		expiredTime  = 15 * time.Minute
		s            = setupMockupServer()
	)

	token := getServiceToken(t, "firstUser", expiredTime)
	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	writer := makeReqCreateTasks(t, s, input, token)

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

	t.Logf("Response data is not empty as expected\n")
}

func TestCreateTheTaskWithValidTokenButNotAccepted(t *testing.T) {
	var (
		expectedCode = http.StatusBadRequest
		expiredTime  = 15 * time.Minute
		mux          = setupMockupServer()
	)

	token := getServiceToken(t, "secondUser", expiredTime)
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
		expiredTime  = 15 * time.Minute
		mux          = setupMockupServer()
	)

	token := getServiceToken(t, "thirdUser", expiredTime)
	input := &model.CreateTaskInput{
		Content: "this is a content",
	}
	for i := 1; i <= 6; i++ {
		writer := makeReqCreateTasks(t, mux, input, token)

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

			t.Logf("New task has been created as expected\n")

			time.Sleep(10 *time.Second)
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
