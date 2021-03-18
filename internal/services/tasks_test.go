package services

import (
	"bytes"
	"github.com/manabie-com/togo/internal/storages"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTask(t *testing.T) {
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTM0Mjc4MjEsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.zJdFGhDFwPcR1-VZyUGmYgNMEDHpNYsttYYvJPcI7h4")
	res := httptest.NewRecorder()

	db, err := storages.GetConnection("postgres", "localhost", 5432, "postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	service := &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &storages.LiteDB{
			DB: db,
		},
	}

	handler := http.HandlerFunc(service.ListTasks)
	handler.ServeHTTP(res, req)

	if statusCode := res.Code; statusCode != http.StatusOK {
		t.Errorf("ListTask() returned wrong status code: got %d - expect %d", statusCode, http.StatusOK)
	}
}

func TestAddTask(t *testing.T) {
	jsonBody := []byte(`{"content":"Second test insert"}`)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU5NjMxNjcsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.-jvGJTJuJ6f_oiLk4MDKOWZTSH8HqpzGO3UXEmDWht")
	res := httptest.NewRecorder()

	db, err := storages.GetConnection("postgres", "localhost", 5432, "postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	service := &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &storages.LiteDB{
			DB: db,
		},
	}

	handler := http.HandlerFunc(service.AddTask)
	handler.ServeHTTP(res, req)

	if statusCode := res.Code; statusCode != http.StatusOK {
		t.Error(res.Body.String())
		t.Errorf("AddTask() returned wrong status code: got %d - expect %d", statusCode, http.StatusOK)
	}
}
