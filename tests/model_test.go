package app_test

import (
	"net/http"
	"testing"
)

func TestModel_ShouldReturnErrorOnPostRequestsWithEmptyPayload(t *testing.T) {
	responseRecorder := makeRequestTo("/todo", "POST", []byte{})

	if responseRecorder.Code != http.StatusBadRequest {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusBadRequest, responseRecorder.Code)
	}
}

func TestModel_ShouldCreateTodoRecordSucessfully(t *testing.T) {
	var jsonStr = []byte(`{"content":"some content", "user_id":1}`)
	responseRecorder := makeRequestTo("/todo", "POST", jsonStr)

	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusCreated, responseRecorder.Code)
	}
}
