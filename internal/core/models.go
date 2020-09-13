package core

import "time"

type Task struct {
  ID          string
  Content     string
  UserID      string
  CreatedDate time.Time
}

type User struct {
  ID      string
  Hash    string
  MaxTodo int
}
