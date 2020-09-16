package integration

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
)

type jsonTask struct {
  ID          string `json:"id"`
  Content     string `json:"content"`
  CreatedDate string `json:"created_date"`
}

const signupUrl = "/signup"
const taskUrl = "/tasks"

func TestSignup(t *testing.T) {
  t.Run("Signup new user and get tasks", func(t *testing.T) {
    reset()
    // signup
    jsonBody := `{
      "id": "firstUser",
      "password": "example",
      "max_todo": 5
    }`
    r, _ := http.NewRequest(http.MethodPost, signupUrl, bytes.NewBufferString(jsonBody))
    w := httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusOK {
      t.Errorf("Expect 200 OK code, got %v", w.Result().StatusCode)
    }
    tokenResponse := struct{
      Data string
    }{}
    err := json.NewDecoder(w.Result().Body).Decode(&tokenResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }

    // get tasks
    r, _ = http.NewRequest(http.MethodGet, taskUrl, bytes.NewBufferString(""))
    r.AddCookie(&http.Cookie{Name: "Authorization", Value: tokenResponse.Data, HttpOnly: true})
    w = httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusOK {
      t.Errorf("Expect 200 OK code, got %v", w.Result().StatusCode)
    }
    taskListResponse := struct{
      Data []*jsonTask
    }{}
    err = json.NewDecoder(w.Result().Body).Decode(&taskListResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }
    if len(taskListResponse.Data) != 0 {
      t.Errorf("Expect no tasks yet, got %v", len(taskListResponse.Data))
    }
  })

  t.Run("Signup existing user with different password", func(t *testing.T) {
    reset()
    // signup
    jsonBody := `{
      "id": "firstUser",
      "password": "example",
      "max_todo": 5
    }`
    r, _ := http.NewRequest(http.MethodPost, signupUrl, bytes.NewBufferString(jsonBody))
    w := httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusOK {
      t.Errorf("Expect 200 OK code, got %v", w.Result().StatusCode)
    }
    tokenResponse := struct{
      Data string
    }{}
    err := json.NewDecoder(w.Result().Body).Decode(&tokenResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }

    // signup again
    jsonBody = `{
      "id": "firstUser",
      "password": "password",
      "max_todo": 5
    }`
    r, _ = http.NewRequest(http.MethodPost, signupUrl, bytes.NewBufferString(jsonBody))
    w = httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusConflict {
      t.Errorf("Expect 409 Conflict code, got %v", w.Result().StatusCode)
    }
    errorResponse := struct{
      Error string
    }{}
    err = json.NewDecoder(w.Result().Body).Decode(&errorResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }
  })

  t.Run("Signup with empty id", func(t *testing.T) {
    reset()
    // signup
    jsonBody := `{
      "id": "",
      "password": "example",
      "max_todo": 5
    }`
    r, _ := http.NewRequest(http.MethodPost, signupUrl, bytes.NewBufferString(jsonBody))
    w := httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusBadRequest {
      t.Errorf("Expect 400 Bad Request code, got %v", w.Result().StatusCode)
    }
    errorResponse := struct{
      Error string
    }{}
    err := json.NewDecoder(w.Result().Body).Decode(&errorResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }
  })

  t.Run("Signup with empty password", func(t *testing.T) {
    reset()
    // signup
    jsonBody := `{
      "id": "firstUser",
      "password": "",
      "max_todo": 5
    }`
    r, _ := http.NewRequest(http.MethodPost, signupUrl, bytes.NewBufferString(jsonBody))
    w := httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusBadRequest {
      t.Errorf("Expect 400 Bad Request code, got %v", w.Result().StatusCode)
    }
    errorResponse := struct{
      Error string
    }{}
    err := json.NewDecoder(w.Result().Body).Decode(&errorResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }
  })

  t.Run("Signup with invalid max todo", func(t *testing.T) {
    reset()
    // signup
    jsonBody := `{
      "id": "firstUser",
      "password": "password",
      "max_todo": -5
    }`
    r, _ := http.NewRequest(http.MethodPost, signupUrl, bytes.NewBufferString(jsonBody))
    w := httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusBadRequest {
      t.Errorf("Expect 400 Bad Request code, got %v", w.Result().StatusCode)
    }
    errorResponse := struct{
      Error string
    }{}
    err := json.NewDecoder(w.Result().Body).Decode(&errorResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }
  })
}

func TestAddTask(t *testing.T) {
  // add a user
  reset()
  maxTodo := 2
  jsonBody := fmt.Sprintf(`{
      "id": "firstUser",
      "password": "example",
      "max_todo": %d
    }`, maxTodo)
  r, _ := http.NewRequest(http.MethodPost, signupUrl, bytes.NewBufferString(jsonBody))
  w := httptest.NewRecorder()
  server.ServeHTTP(w, r)
  if w.Result().StatusCode != http.StatusOK {
    t.Errorf("Expect 200 OK code, got %v", w.Result().StatusCode)
  }
  tokenResponse := struct{
    Data string
  }{}
  err := json.NewDecoder(w.Result().Body).Decode(&tokenResponse)
  if err != nil {
    t.Errorf("Expect no error, got %v", err)
  }

  // get tasks
  token := tokenResponse.Data

  for i := 0; i < maxTodo; i++ {
    t.Run(fmt.Sprintf("Add task %d", i+1), func(t *testing.T) {
      jsonBody = `{
      "content": "another task"
    }`
      r, _ = http.NewRequest(http.MethodPost, taskUrl, bytes.NewBufferString(jsonBody))
      r.AddCookie(&http.Cookie{Name: "Authorization", Value: token, HttpOnly: true})
      w = httptest.NewRecorder()
      server.ServeHTTP(w, r)
      if w.Result().StatusCode != http.StatusOK {
        t.Errorf("Expect 200 OK code, got %v", w.Result().StatusCode)
      }
      addTaskResponse := struct{
        Data *jsonTask
      }{}
      err = json.NewDecoder(w.Result().Body).Decode(&addTaskResponse)
      if err != nil {
        t.Errorf("Expect no error, got %v", err)
      }
    })
  }

  t.Run("Add another task, maximum reached", func(t *testing.T) {
    jsonBody := `{
      "content": "another task"
    }`
    r, _ := http.NewRequest(http.MethodPost, taskUrl, bytes.NewBufferString(jsonBody))
    r.AddCookie(&http.Cookie{Name: "Authorization", Value: token, HttpOnly: true})
    w := httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusBadRequest {
      t.Errorf("Expect 400 Bad Request code, got %v", w.Result().StatusCode)
    }
    errorResponse := struct{
      Error string
    }{}
    err := json.NewDecoder(w.Result().Body).Decode(&errorResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }
  })

  t.Run("Get all tasks", func(t *testing.T) {
    r, _ := http.NewRequest(http.MethodGet, taskUrl, bytes.NewBufferString(""))
    r.AddCookie(&http.Cookie{Name: "Authorization", Value: token, HttpOnly: true})
    w := httptest.NewRecorder()
    server.ServeHTTP(w, r)
    if w.Result().StatusCode != http.StatusOK {
      t.Errorf("Expect 200 OK code, got %v", w.Result().StatusCode)
    }
    taskListResponse := struct{
      Data []*jsonTask
    }{}
    err := json.NewDecoder(w.Result().Body).Decode(&taskListResponse)
    if err != nil {
      t.Errorf("Expect no error, got %v", err)
    }
    if len(taskListResponse.Data) != maxTodo {
      t.Errorf("Expect %v tasks yet, got %v", maxTodo, len(taskListResponse.Data))
    }
  })
}
