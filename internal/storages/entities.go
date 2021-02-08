package storages

import (
	"bytes"
	"encoding/json"
	"io"
)

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string
	Password string
}

func (t *Task) ToIOReader() io.Reader {
	byteArr, _ := json.Marshal(t)
	return bytes.NewReader(byteArr)
}
