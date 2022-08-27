package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	ctrl "backend_test/task/controllers"
	"backend_test/task/services"
)

func TestLimit(t *testing.T) {
	var testData = []struct {
		input1   int
		input2   int
		expected bool
	}{
		{2, 5, true},
		{5, 5, false},
		{6, 5, false},
	}

	for _, td := range testData {
		if output := services.LimitValidator(td.input1, td.input2); output != td.expected {
			t.Error("Test Failed: {} {} inputted, expected {}", td.input1, td.input2, td.expected)
		}
	}
}

func TestGetAllTasks(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/task/list", nil)
	response := httptest.NewRecorder()

	handler := http.HandlerFunc(ctrl.GetAllPaymentAPI)
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Error("Test Failed: expected {}", http.StatusOK)
	}
}
