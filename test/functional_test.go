package test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	_ = httptest.NewRequest("GET", "http://localhost:9000/api/v1/todo", strings.NewReader(`
									{
										"title": "make breakfast",
										"user_id": 1
									}`))

	w := httptest.NewRecorder()

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
