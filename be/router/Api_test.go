package router

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
	"todo/be/utils"
)

func TestApiTodo_add(t *testing.T) {
	response, err := http.Post(
		"http://localhost:8008/api/todo",
		"application/json",
		bytes.NewReader([]byte("{\"Text\":\"Todo taks text\"}")),
	)
	if utils.IsError(err) {
		t.Errorf("Please start server")
	}
	token := response.Header.Get(key_TokenHeader)
	response.Body.Close()
	if len(token) == 0 {
		t.Errorf("Response expect to have header Token")
	}
	response1, err1 := callApiTodo("", token)
	if utils.IsError(err1) {
		t.Errorf("Response1 is error")
	}
	resBytes, _ := io.ReadAll(response1.Body)
	if !strings.Contains(string(resBytes), "Invald Text") {
		t.Errorf("Resonse string expect to contain 'Invald Text'")
	}
	response1.Body.Close()

	response2, err2 := callApiTodo("Text of todo task", token)
	if utils.IsError(err2) {
		t.Errorf("Response2 is error")
		return
	}
	resBytes, _ = io.ReadAll(response2.Body)
	if !strings.Contains(string(resBytes), "Add todo task success") {
		t.Errorf("Resonse string expect to contain 'Add todo task success'")
	}
	response2.Body.Close()
}

func callApiTodo(text string, token string) (*http.Response, error) {
	client := http.DefaultClient
	request, _ := http.NewRequest(
		http_POST,
		"http://localhost:8008/api/todo",
		bytes.NewReader([]byte("{\"Text\":\""+text+"\"}")),
	)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Token", token)
	return client.Do(request)
}
