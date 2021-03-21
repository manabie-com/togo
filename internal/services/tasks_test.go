package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/manabie-com/togo/internal/ratelimiters"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

var (
	JWTKey = "wqGyEBBfPK9w3Lxw"
)

func TestUserIDFromCtx(t *testing.T) {
	testCases := []struct {
		Input   context.Context
		Output1 string
		Output2 bool
	}{
		{context.WithValue(context.Background(), userAuthKey(0), "1"), "1", true},
		{context.WithValue(context.Background(), userAuthKey(0), ""), "", true},
		{context.Background(), "", false},
	}

	for i, testCase := range testCases {
		output1, output2 := userIDFromCtx(testCase.Input)
		if output1 != testCase.Output1 {
			t.Errorf("Error at testcase: %d\nExpect output1: %v\nReceive: %v\n", i, testCase.Output1, output1)
		}
		if output2 != testCase.Output2 {
			t.Errorf("Error at testcase: %d\nExpect output2: %v\nReceive: %v\n", i, testCase.Output2, output2)
		}
	}
}

func TestCreateToken(t *testing.T) {
	testCases := []struct {
		Input   string
		IsValid bool
		Err     error
	}{
		{"test", true, nil},
		{"", true, nil},
	}
	whiteMouse := ToDoService{
		JWTKey: JWTKey,
	}
	for i, testCase := range testCases {
		token, err := whiteMouse.createToken(testCase.Input)

		// Valid token
		req := &http.Request{
			Header: http.Header{"Authorization": []string{token}},
		}

		_, isValid := whiteMouse.validToken(req)
		if isValid != testCase.IsValid {
			t.Errorf("Error at testcase: %d\nExpect isValid: %v\nReceive: %v\n", i, testCase.IsValid, isValid)
		}
		if err != testCase.Err {
			t.Errorf("Error at testcase: %d\nExpect err: %v\nReceive: %v\n", i, testCase.Err, err)
		}
	}
}

func TestValidToken(t *testing.T) {
	token := genToken("Test")
	req1 := &http.Request{Header: http.Header{"Authorization": []string{token}}}
	req2 := req1.WithContext(context.WithValue(req1.Context(), userAuthKey(0), "Test"))

	testCases := []struct {
		Input   *http.Request
		Output1 *http.Request
		IsValid bool
	}{
		{ // Wrong token
			Input:   &http.Request{Header: http.Header{"Authorization": []string{"1"}}},
			Output1: &http.Request{Header: http.Header{"Authorization": []string{"1"}}},
			IsValid: false,
		},
		{ // Time out
			Input:   &http.Request{Header: http.Header{"Authorization": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTYyNDE0NjgsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.4fbKJSHlVjlAwJilv0M_5pgekhE-T5HfZgPMMLmYKoY"}}},
			Output1: &http.Request{Header: http.Header{"Authorization": []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTYyNDE0NjgsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.4fbKJSHlVjlAwJilv0M_5pgekhE-T5HfZgPMMLmYKoY"}}},
			IsValid: false,
		},
		{ // Success
			Input:   req1,
			Output1: req2,
			IsValid: true,
		},
	}

	whiteMouse := ToDoService{
		JWTKey: JWTKey,
	}
	for i, testCase := range testCases {
		output1, isValid := whiteMouse.validToken(testCase.Input)
		if isValid != testCase.IsValid {
			t.Errorf("Error at testcase: %d\nExpect isValid: %v\nReceive: %v\n", i, testCase.IsValid, isValid)
		}
		if !reflect.DeepEqual(output1, testCase.Output1) {
			t.Errorf("Error at testcase: %d\nExpect output1: %+v\nReceive: %+v\n", i, testCase.Output1, output1)
		}
	}
}

type check func(*httptest.ResponseRecorder) bool

func checkCodeStatus(code int) check {
	return func(w *httptest.ResponseRecorder) bool {
		if w.Code != code {
			return false
		}
		return true
	}
}

func checkValidToken() check {
	return func(w *httptest.ResponseRecorder) bool {
		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		var tokeData struct {
			Data string `json:"data"`
		}
		json.Unmarshal(body, &tokeData)
		whiteMouse := ToDoService{
			JWTKey: JWTKey,
		}

		reqWhiteMouse := &http.Request{Header: http.Header{"Authorization": []string{tokeData.Data}}}
		if _, valid := whiteMouse.validToken(reqWhiteMouse); valid {
			return true
		}
		return false
	}
}

func checkBody(body string) check {
	return func(w *httptest.ResponseRecorder) bool {
		resp := w.Result()
		respBody, _ := ioutil.ReadAll(resp.Body)
		if string(respBody) != body {
			return false
		}
		return true
	}
}

func genCreateTaskRequest() *http.Request {
	payload := strings.NewReader(`{
    "content": "another content"
}`)
	reqCreateTask := httptest.NewRequest("POST", "http://example.com/tasks", payload)
	reqCreateTask.Header.Add("Authorization", genToken(""))
	return reqCreateTask
}

func TestServeHTTP(t *testing.T) {
	// Init
	db, err := sql.Open("sqlite3", "../../data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	ratelimiters.InitLocalRatelimiter(&sqllite.LiteDB{db})

	reqGetTask := httptest.NewRequest("GET", "http://example.com/tasks", nil)
	reqGetTask.Header.Add("Authorization", genToken(""))

	testCases := []struct {
		Req    *http.Request
		Resp   *httptest.ResponseRecorder
		Checks []check
	}{
		{
			httptest.NewRequest("GET", "http://example.com/login?user_id=firstUser&password=example", nil),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(200), checkValidToken()},
		},
		{
			httptest.NewRequest("GET", "http://example.com/login?user_id=firstUser&password=example1", nil),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(401), checkBody("{\"error\":\"incorrect user_id/pwd\"}\n")},
		},
		{
			httptest.NewRequest("GET", "http://example.com/tasks", nil),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(401)},
		},
		{
			reqGetTask,
			httptest.NewRecorder(),
			[]check{checkCodeStatus(200)},
		},
		{
			genCreateTaskRequest(),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(200)},
		},
		{
			genCreateTaskRequest(),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(200)},
		},
		{
			genCreateTaskRequest(),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(200)},
		},
		{
			genCreateTaskRequest(),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(200)},
		},
		{
			genCreateTaskRequest(),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(200)},
		},
		{ // Reach limit
			genCreateTaskRequest(),
			httptest.NewRecorder(),
			[]check{checkCodeStatus(429)},
		},
	}

	whiteMouse := ToDoService{
		JWTKey: JWTKey,
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}
	for i, testCase := range testCases {
		whiteMouse.ServeHTTP(testCase.Resp, testCase.Req)
		for _, check := range testCase.Checks {
			if ok := check(testCase.Resp); !ok {
				t.Errorf("Error at testcase: %d. Check %v", i, runtime.FuncForPC(reflect.ValueOf(check).Pointer()).Name())
			}
		}
	}
}

func genToken(userId string) string {
	if userId == "" {
		userId = "firstUser"
	}
	whiteMouse := ToDoService{
		JWTKey: JWTKey,
	}

	token, _ := whiteMouse.createToken(userId)
	return token
}
