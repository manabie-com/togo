package integration_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestLogin(t *testing.T) {
	samples := []struct {
		url        string
		statusCode int
	}{
		{
			url:        `http://localhost:5050/login?user_id=firstUser&password=example`,
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
		fmt.Printf("run test: %v/%v\n", i, len(samples))
		r := httptest.NewRequest("GET", sample.url, nil)
		w := httptest.NewRecorder()
		toDoService.ServeHTTP(w, r)

		if w.Code != sample.statusCode {
			t.Errorf("wrong status code want %v but get %v", sample.statusCode, w.Code)
		}
	}
}

func TestGetTaskList(t *testing.T) {
	samples := []struct {
		url        string
		statusCode int
	}{
		{
			url:        `http://localhost:5050/tasks?created_date=2020-06-29`,
			statusCode: http.StatusOK,
		},
		{
			url:        `http://localhost:5050/tasks?created_date=2020--11`,
			statusCode: http.StatusOK,
		},
	}

	for i, sample := range samples {
		fmt.Printf("run test: %v/%v\n", i, len(samples))
		validLoginUrl := "http://localhost:5050/login?user_id=firstUser&password=example"
		token, err := getToken(http.MethodGet, validLoginUrl)
		if err != nil {
			t.Errorf("error can not get token %v", err)
		}
		r := httptest.NewRequest("GET", sample.url, nil)
		r.Header.Add("Authorization", token)
		w := httptest.NewRecorder()
		toDoService.ServeHTTP(w, r)
		if w.Code != sample.statusCode {
			t.Errorf("wrong status code want %v but get %v", sample.statusCode, w.Code)
		}
	}
}
