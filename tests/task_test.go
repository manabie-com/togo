package tests

import (
	"context"
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_AddTask_WithOutAuth(t *testing.T) {
	task, err := sendAddTask(context.Background(), "", &addTaskReq{
		Content: "text",
	})
	assert.Nil(t, task)
	assert.Error(t, err)
}

func Test_AddTask_Successful(t *testing.T) {
	ctx := context.Background()
	username := faker.Username()
	password := "a123456"
	err := sendRegister(ctx, &registerReq{
		FullName:    faker.Name(),
		Username:    username,
		Password:    password,
		TasksPerDay: 2,
	})
	assert.NoError(t, err)
	loginRes, err := sendLogin(ctx, &loginReq{
		Username: username,
		Password: password,
	})
	assert.NoError(t, err)
	task, err := sendAddTask(ctx, loginRes.Token, &addTaskReq{
		Content: "text",
	})
	assert.NoError(t, err)
	assert.NotNil(t, task)
}

func Test_AddTask_Async5Tasks(t *testing.T) {
	ctx := context.Background()
	username := faker.Username()
	password := "a123456"
	err := sendRegister(ctx, &registerReq{
		FullName:    faker.Name(),
		Username:    username,
		Password:    password,
		TasksPerDay: 1000,
	})
	assert.NoError(t, err)
	loginRes, err := sendLogin(ctx, &loginReq{
		Username: username,
		Password: password,
	})
	assert.NoError(t, err)
	haveError := false
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			if _, err := sendAddTask(ctx, loginRes.Token, &addTaskReq{Content: "text"}); err != nil {
				haveError = true
			}
			wg.Done()
		}()
	}
	wg.Wait()
	assert.False(t, haveError)
}

func Test_AddTask_AsyncMoreThan5Tasks(t *testing.T) {
	ctx := context.Background()
	username := faker.Username()
	password := "a123456"
	err := sendRegister(ctx, &registerReq{
		FullName:    faker.Name(),
		Username:    username,
		Password:    password,
		TasksPerDay: 1000,
	})
	assert.NoError(t, err)
	loginRes, err := sendLogin(ctx, &loginReq{
		Username: username,
		Password: password,
	})
	assert.NoError(t, err)
	haveError := false
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if _, err := sendAddTask(ctx, loginRes.Token, &addTaskReq{Content: "text"}); err != nil {
				haveError = true
			}
			wg.Done()
		}()
	}
	wg.Wait()
	assert.False(t, haveError)
}
