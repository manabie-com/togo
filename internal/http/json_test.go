package http

import (
  "bytes"
  "encoding/json"
  "net/http"
  "testing"
)

func TestJsonUserParser(t *testing.T) {
  t.Run("Test parse user info password", func(t *testing.T) {
    t.Run("correct format", func(t *testing.T) {
      data := struct {
        Id       string `json:"id"`
        Password string `json:"password"`
        MaxTodo  int    `json:"max_todo"`
      }{
        Id:       "firstUser",
        Password: "password",
        MaxTodo:  5,
      }
      jsonBody, _ := json.Marshal(data)
      r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
      parser := JsonUserParser{}
      user, password, err := parser.parseUserInfoPassword(r)
      if user == nil {
        t.Error("Expect non-nil user, got nil")
      }
      if err != nil {
        t.Errorf("Expect no error, got %v", err)
      }
      if user.ID != data.Id || user.MaxTodo != data.MaxTodo || password != data.Password {
        t.Errorf("Parsed data mismatch with original, expect %v %v %v, got %v %v %v", data.Id, data.Password,
          data.MaxTodo, user.ID, password, user.MaxTodo)
      }
    })
    t.Run("missing field", func(t *testing.T) {
      data := struct {
        Id       string `json:"id"`
        Password string `json:"password"`
      }{
        Id:       "firstUser",
        Password: "password",
      }
      jsonBody, _ := json.Marshal(data)
      r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
      parser := JsonUserParser{}
      user, password, err := parser.parseUserInfoPassword(r)
      if user != nil || password != "" || err != ErrParseUserInfoPassword {
        t.Errorf("Expect (nil user, empty password, %v error), got (%v,%v,%v)", ErrParseUserInfoPassword, user, password, err)
      }
    })
    t.Run("wrong field name", func(t *testing.T) {
      data := struct {
        Id       string `json:"ids"`
        Password string `json:"password"`
        MaxTodo  int    `json:"max_todo"`
      }{
        Id:       "firstUser",
        Password: "password",
        MaxTodo:  5,
      }
      jsonBody, _ := json.Marshal(data)
      r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
      parser := JsonUserParser{}
      user, password, err := parser.parseUserInfoPassword(r)
      if user != nil || password != "" || err != ErrParseUserInfoPassword {
        t.Errorf("Expect (nil user, empty password, %v error), got (%v,%v,%v)", ErrParseUserInfoPassword, user, password, err)
      }
    })
    t.Run("wrong data type", func(t *testing.T) {
      data := struct {
        Id       string `json:"id"`
        Password string `json:"password"`
        MaxTodo  string `json:"max_todo"`
      }{
        Id:       "firstUser",
        Password: "password",
        MaxTodo:  "5",
      }
      jsonBody, _ := json.Marshal(data)
      r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
      parser := JsonUserParser{}
      user, password, err := parser.parseUserInfoPassword(r)
      if user != nil || password != "" || err != ErrParseUserInfoPassword {
        t.Errorf("Expect (nil user, empty password, %v error), got (%v,%v,%v)", ErrParseUserInfoPassword, user, password, err)
      }
    })
  })

  t.Run("Test parse user id password", func(t *testing.T) {
    t.Run("correct format", func(t *testing.T) {
      data := struct {
        Id       string `json:"id"`
        Password string `json:"password"`
      }{
        Id:       "firstUser",
        Password: "password",
      }
      jsonBody, _ := json.Marshal(data)
      r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
      parser := JsonUserParser{}
      id, password, err := parser.parseIdPassword(r)
      if err != nil {
        t.Errorf("Expect no error, got %v", err)
      }
      if id != data.Id || password != data.Password {
        t.Errorf("Parsed data mismatch with original, expect %v %v, got %v %v", data.Id, data.Password,
          id, password)
      }
    })
    t.Run("missing field", func(t *testing.T) {
      data := struct {
        Password string `json:"password"`
      }{
        Password: "password",
      }
      jsonBody, _ := json.Marshal(data)
      r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
      parser := JsonUserParser{}
      id, password, err := parser.parseIdPassword(r)
      if id != "" || password != "" || err != ErrParseIdPassword {
        t.Errorf("Expect (empty id, empty password, %v error), got (%v,%v,%v)", ErrParseIdPassword, id,
          password, err)
      }
    })
    t.Run("wrong field name", func(t *testing.T) {
      data := struct {
        Id       string `json:"ids"`
        Password string `json:"password"`
      }{
        Id:       "firstUser",
        Password: "password",
      }
      jsonBody, _ := json.Marshal(data)
      r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
      parser := JsonUserParser{}
      id, password, err := parser.parseIdPassword(r)
      if id != "" || password != "" || err != ErrParseIdPassword {
        t.Errorf("Expect (empty id, empty password, %v error), got (%v,%v,%v)", ErrParseIdPassword, id,
          password, err)
      }
    })
    t.Run("wrong data type", func(t *testing.T) {
      data := struct {
        Id       int `json:"id"`
        Password string `json:"password"`
      }{
        Id:       1,
        Password: "password",
      }
      jsonBody, _ := json.Marshal(data)
      r, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBody))
      parser := JsonUserParser{}
      id, password, err := parser.parseIdPassword(r)
      if id != "" || password != "" || err != ErrParseIdPassword {
        t.Errorf("Expect (empty id, empty password, %v error), got (%v,%v,%v)", ErrParseIdPassword, id,
          password, err)
      }
    })
  })
}
