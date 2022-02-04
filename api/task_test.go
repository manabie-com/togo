package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedTask struct {
	mock.Mock
}

func TestCreateTask(t *testing.T) {
	w := httptest.NewRecorder()
	var jsonString = []byte(`{"title": "Test","content": "API Test","is_complete": true,"fullname": "Roan Dino"}`)
	r := httptest.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(jsonString))
	r.Header.Set("Content-Type", "application/json")
	fmt.Println(w.Body.String())

	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"title":"Test","content":"API Test","is_complete":true,"fullname":"Roan Dino"}`
	assert.Contains(t, w.Body.String(), expected, "it should contain expected")
}
