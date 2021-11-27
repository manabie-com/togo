package test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	svc := Init(t)
	req := httptest.NewRequest("POST", "http://localhost:9000/api/v1/todo", strings.NewReader(`
									{
										"title": "make breakfast",
										"user_id": 1
									}`))
	expected := `{"todo":{"id":1,"title":"make breakfast","user_id":1}}`

	w := httptest.NewRecorder()
	handler := svc.CreateTodoHandler()
	handler.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	compare(t, expected, string(body))
}
