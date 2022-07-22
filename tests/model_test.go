package app_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestModel_ShouldCreateTodoRecordSucessfully(t *testing.T) {
	var jsonStr = []byte(`{"content":"some content", "user_id":1}`)
	request, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()
	app.Router.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusCreated, responseRecorder.Code)
	}
}
