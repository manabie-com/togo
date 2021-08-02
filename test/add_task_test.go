package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type addTaskRequest struct {
	Content string `json:"content"`
}

type addTaskResponse struct {
	ID           string `json:"id"`
	Content      string `json:"content"`
	UserID       string `json:"user_id"`
	CreatedDate  string `json:"created_date"`
	NumberInDate int    `json:"number_in_date"`
}

func sendAddTaskRequest(ctx context.Context, token string, req *addTaskRequest) (*addTaskResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", "http://localhost:5050/tasks", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := &addTaskResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return nil, err
	}
	return res, nil
}
func Test_AddTask(t *testing.T) {
	ctx := context.Background()
	u := &loginRequest{
		UserID:   "firstUser",
		Password: "example",
	}

	task := &addTaskRequest{
		Content: "something content",
	}

	res, err := sendLoginRequest(ctx, u)
	assert.NoError(t, err)

	resp, err := sendAddTaskRequest(ctx, res.Data, task)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_AddTask_5Task(t *testing.T) {
	ctx := context.Background()
	u := &loginRequest{
		UserID:   "firstUser",
		Password: "example",
	}

	task := &addTaskRequest{
		Content: "something content",
	}

	res, err := sendLoginRequest(ctx, u)
	assert.NoError(t, err)

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			resp, err := sendAddTaskRequest(ctx, res.Data, task)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
			wg.Done()
		}()
	}
	wg.Wait()
}

func Test_AddTask_LargerThan5Task(t *testing.T) {
	ctx := context.Background()
	u := &loginRequest{
		UserID:   "firstUser",
		Password: "example",
	}

	task := &addTaskRequest{
		Content: "something content",
	}

	res, err := sendLoginRequest(ctx, u)
	assert.NoError(t, err)

	flag := false
	var wg sync.WaitGroup
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go func() {
			_, err := sendAddTaskRequest(ctx, res.Data, task)
			if err != nil {
				flag = true
			}
			wg.Done()
		}()
	}
	wg.Wait()

	assert.True(t, flag)
}
