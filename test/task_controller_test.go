package test

import (
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/controller"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskControllerSuccess(t *testing.T) {
	db, err := config.GetPostgersDB()

	if err != nil{
		t.Fatal(err)
	}

	task :=controller.NewTaskController(db)
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(task.ListTasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}