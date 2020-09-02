package integration_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
)

func TestLogin(t *testing.T) {
	//Create new user
	u, err := seedUser(toDoService.Store.DB)
	if err != nil {
		t.Errorf("Can not create new user %v\n", err)
	}

	samples := []struct {
		url        string
		statusCode int
	}{
		{
			url:        fmt.Sprintf("http://localhost:5050/login?user_id=%v&password=%v", u.ID, u.Password),
			statusCode: http.StatusOK,
		},
		{
			url:        `http://localhost:5050/login?user_id=first&password=example`,
			statusCode: http.StatusUnauthorized,
		},
		{
			url:        `http://localhost:5050/login?user_id=1`,
			statusCode: http.StatusUnauthorized,
		},
		{
			url:        `http://localhost:5050/login?password=example`,
			statusCode: http.StatusUnauthorized,
		},
		{
			url:        `http://localhost:5050/login?some=thing`,
			statusCode: http.StatusUnauthorized,
		},
		{
			url:        `http://localhost:5050/login?password=`,
			statusCode: http.StatusUnauthorized,
		},
		{
			url:        `http://localhost:5050/login?`,
			statusCode: http.StatusUnauthorized,
		},
		// {
		// 	url:        `http://localhost:5050/login/`,
		// 	statusCode: http.StatusUnauthorized,
		// },
	}

	for i, sample := range samples {
		fmt.Printf("run test: %v/%v\n", i+1, len(samples))
		r := httptest.NewRequest("GET", sample.url, nil)
		w := httptest.NewRecorder()
		toDoService.ServeHTTP(w, r)

		if w.Code != sample.statusCode {
			t.Errorf("wrong status code want %v but get %v", sample.statusCode, w.Code)
		}
	}

	// Refresh the tables
	if err = truncate(toDoService.Store.DB); err != nil {
		t.Errorf("cant not truncate database tables %v\n", err)
	}
}

func TestGetTaskList(t *testing.T) {

	user, err := seedUser(toDoService.Store.DB)
	if err != nil {
		t.Errorf("Can not create new user %v\n", err)
	}

	// taskitems belong to new user
	taskItems, err := seedTaskItems(toDoService.Store.DB, user)
	if err != nil {
		t.Errorf("Can not create task items sample %v\n", err)
	}

	samples := []struct {
		url          string
		numberOfTask int
		statusCode   int
	}{
		{
			url:          fmt.Sprintf(`http://localhost:5050/tasks?created_date=%v`, taskItems[0].CreatedDate),
			numberOfTask: len(taskItems),
			statusCode:   http.StatusOK,
		},
		{
			url:          `http://localhost:5050/tasks?created_date=2020--11`,
			numberOfTask: 0,
			statusCode:   http.StatusOK,
		},
		{
			url:          `http://localhost:5050/tasks?yesterday`,
			numberOfTask: 0,
			statusCode:   http.StatusOK,
		},
	}

	for i, sample := range samples {
		fmt.Printf("run test: %v/%v\n", i+1, len(samples))
		// Login to get token of new user
		validLoginUrl := fmt.Sprintf("http://localhost:5050/login?user_id=%v&password=%v", user.ID, user.Password)
		token, err := login(http.MethodGet, validLoginUrl)
		if err != nil {
			t.Errorf("error can not get token %v\n", err)
		}

		// Create request and response
		r := httptest.NewRequest(http.MethodGet, sample.url, nil)
		r.Header.Add("Authorization", token)
		w := httptest.NewRecorder()
		toDoService.ServeHTTP(w, r)

		// Validate answer
		if w.Code != sample.statusCode {
			t.Errorf("wrong status code want %v but get %v\n", sample.statusCode, w.Code)
		}

		dataJSON := make(map[string][]storages.Task)
		err = json.Unmarshal(w.Body.Bytes(), &dataJSON)
		if err != nil {
			t.Errorf("can not parse data response %v\n", err)
		}

		if sample.numberOfTask != len(dataJSON["data"]) {
			t.Errorf("get task list fail want %v items but only get %v items\n", len(taskItems), len(dataJSON["data"]))
		}
	}

	// Refresh the tables
	if err = truncate(toDoService.Store.DB); err != nil {
		t.Errorf("cant not truncate database tables %v\n", err)
	}
}

func TestAddTask(t *testing.T) {
	samples := []struct {
		url        string
		inputJSON  string
		content    string
		statusCode int
	}{
		{
			url:        `http://localhost:5050/tasks`,
			inputJSON:  `{"content": "another content again"}`,
			content:    `another content again`,
			statusCode: http.StatusOK,
		},
		{
			url:        `http://localhost:5050/tasks`,
			inputJSON:  `{"content": "sad of content"}`,
			content:    `sad of content`,
			statusCode: http.StatusOK,
		},
		{
			url:        `http://localhost:5050/tasks`,
			inputJSON:  `{"content": "happy to see content"}`,
			content:    `happy to see content`,
			statusCode: http.StatusOK,
		},
	}

	// create new user in database users table
	u, err := seedUser(toDoService.Store.DB)
	if err != nil {
		t.Errorf("Can not create new user %v\n", err)
	}

	for i, sample := range samples {
		fmt.Printf("run test: %v/%v\n", i+1, len(samples))
		// Login to get token of new user
		validLoginUrl := fmt.Sprintf("http://localhost:5050/login?user_id=%v&password=%v", u.ID, u.Password)
		token, err := login(http.MethodGet, validLoginUrl)
		if err != nil {
			t.Errorf("error can not get token %v\n", err)
		}

		// Create new request
		r := httptest.NewRequest(http.MethodPost, sample.url, strings.NewReader(sample.inputJSON))
		r.Header.Add("Authorization", token)
		w := httptest.NewRecorder()
		toDoService.ServeHTTP(w, r)

		// Validate answer
		if w.Code != sample.statusCode {
			t.Errorf("wrong status code want %v but get %v\n", sample.statusCode, w.Code)
		}

		dataJSON := make(map[string]storages.Task)
		err = json.Unmarshal(w.Body.Bytes(), &dataJSON)
		if err != nil {
			t.Errorf("can not parse data response %v\n", err)
		}

		if sample.content != dataJSON["data"].Content {
			t.Errorf("not the same content expect \"%v\" but got \"%v\" \n", sample.content, dataJSON["content"])
		}
	}

	// Refresh the tables
	if err = truncate(toDoService.Store.DB); err != nil {
		t.Errorf("cant not truncate database tables %v\n", err)
	}
}

func login(method string, validLoginUrl string) (string, error) {
	r := httptest.NewRequest(method, validLoginUrl, nil)
	w := httptest.NewRecorder()
	toDoService.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		return "", fmt.Errorf("wrong status code want %v but get %v\n", http.StatusOK, w.Code)
	}
	dataJson := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &dataJson)
	if err != nil {
		return "", fmt.Errorf("can not parse response data %v\n", err)
	}
	token, ok := dataJson["data"].(string)
	if !ok {
		return "", fmt.Errorf("data is invalid, get %v\n", token)
	}

	return token, nil
}
