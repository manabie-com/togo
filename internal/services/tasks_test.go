package services

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
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
	whiteMouse := ToDoService{}
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
	whiteMouse := ToDoService{}

	token, _ := whiteMouse.createToken("Test")
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

func TestServeHTTP(t *testing.T) {
	testCases := []struct {
		Req  *http.Request
		Resp *httptest.ResponseRecorder
	}{
		{
			httptest.NewRequest("GET", "http://example.com/login?user_id=firstUser&password=example", nil),
			httptest.NewRecorder(),
		},
	}

	whiteMouse := ToDoService{}
	for _, testCase := range testCases {
		whiteMouse.ServeHTTP(testCase.Resp, testCase.Req)
	}
}
