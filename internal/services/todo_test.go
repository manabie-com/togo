package services

import (
	"context"
	"fmt"
	"github.com/banhquocdanh/togo/internal/storages"
	"github.com/banhquocdanh/togo/internal/storages/mocks"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var listTasksTestInput = []struct {
	name       string
	userID     string
	createDate string

	err   error
	tasks []*storages.Task
}{
	{
		"test happy case",
		"user_1",
		"2021-03-14",
		nil,
		nil,
	},
	{
		"test empty userID",
		"",
		"2021-03-14",
		fmt.Errorf("user id invalid"),
		nil,
	},
	{
		"test empty createDate",
		"user_1",
		"",
		fmt.Errorf("created date invalid"),
		nil,
	},
}

func TestListTasks(t *testing.T) {
	storeMock := &mocks.StoreInterface{}
	storeMock.On(
		"RetrieveTasks",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
	).Return(nil, nil)

	srv := NewToDoService(WithStore(storeMock))

	for _, tt := range listTasksTestInput {
		t.Run(tt.name, func(t *testing.T) {
			tasks, err := srv.ListTasks(context.Background(), tt.userID, tt.createDate)
			if (err != tt.err) && (err == nil || tt.err == nil || err.Error() != tt.err.Error()) {
				t.Errorf("Error: got %s, want %s", err, tt.err)
			}
			if len(tasks) != len(tt.tasks) {
				t.Errorf("Taks: got %+v, want %+v", tasks, tt.tasks)
			}
		})
	}
}

var addTaskTestInput = []struct {
	name    string
	user    string
	content string
	err     error
	task    *storages.Task
}{
	{
		"test happy case",
		"user_1",
		"content",
		nil,
		&storages.Task{
			ID:          "any",
			Content:     "content",
			UserID:      "user_1",
			CreatedDate: time.Date(2021, 03, 14, 0, 0, 0, 0, time.Local).Format("2006-01-02"),
		},
	},
	{
		"test empty userID",
		"",
		"content",
		fmt.Errorf("invalid userID"),
		nil,
	},
	{
		"test empty content",
		"user_1",
		"",
		fmt.Errorf("invalid task's content"),
		nil,
	},
}

func TestAddTask(t *testing.T) {
	storeMock := &mocks.StoreInterface{}
	storeMock.On(
		"AddTask",
		mock.Anything,
		mock.AnythingOfType("*storages.Task"),
	).Return(func(ctx context.Context, t *storages.Task) error { return nil })

	srv := NewToDoService(WithStore(storeMock))
	taskCompareFunc := func(t1, t2 *storages.Task) bool {
		if t1 == t2 {
			return true
		}
		if t1 == nil {
			return false
		}
		if t1.Content != t2.Content {
			return false
		}
		if t1.CreatedDate != t2.CreatedDate {
			return false
		}
		if t1.UserID != t2.UserID {
			return false
		}
		return true
	}

	for _, tt := range addTaskTestInput {
		t.Run(tt.name, func(t *testing.T) {
			task, err := srv.AddTask(context.Background(), tt.user, tt.content)
			if (err != tt.err) && (err == nil || tt.err == nil || err.Error() != tt.err.Error()) {
				t.Errorf("Error: got %s, want %s", err, tt.err)
			}
			if taskCompareFunc(task, tt.task) == false {
				t.Errorf("Taks: got %+v, want %+v", task, tt.task)
			}
		})
	}
}

var jwtKeyTest = "wqGyEBBfPK9w3Lxw"
var loginTestInput = []struct {
	name   string
	userID string
	pw     string

	err   error
	token string
}{
	{
		"test happy case",
		"user_1",
		"pw_user_1",
		nil,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU2NTU3MDAsInVzZXJfaWQiOiJ1c2VyXzEifQ.L9ZKMYUFJHl4CnYti5MM_jkUt9zA0CGsVmbqOR-zhvU",
	},
	{
		"test empty userID",
		"",
		"pw_user_1",
		fmt.Errorf("invalid user/pw"),
		"",
	},
	{
		"test empty createDate",
		"user_1",
		"",
		fmt.Errorf("invalid user/pw"),
		"",
	},
	{
		"test empty userID and pw",
		"",
		"",
		fmt.Errorf("invalid user/pw"),
		"",
	},
	{
		"test empty createDate",
		"user_1",
		"wrong_pw",
		fmt.Errorf("user/pw is incorrect"),
		"",
	},
}

func TestLogin(t *testing.T) {
	storeMock := &mocks.StoreInterface{}
	storeMock.On(
		"ValidateUser",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
	).Return(func(ctx context.Context, user, pw string) bool {
		if user == "user_1" && pw == "pw_user_1" {
			return true
		}
		return false
	})

	Now = func() time.Time {
		return time.Date(2021, 03, 14, 0, 0, 0, 0, time.Local)
	}

	srv := NewToDoService(WithStore(storeMock))

	for _, tt := range loginTestInput {
		t.Run(tt.name, func(t *testing.T) {
			token, err := srv.Login(context.Background(), tt.userID, tt.pw, jwtKeyTest)
			if (err != tt.err) && (err == nil || tt.err == nil || err.Error() != tt.err.Error()) {
				t.Errorf("Error: got %s, want %s", err, tt.err)
			}
			if token != tt.token {
				t.Errorf("Taks: got %+v, want %+v", token, tt.token)
			}
		})
	}
}

var validTokenTestInput = []struct {
	name  string
	token string

	err    error
	userID string
}{
	{
		"test happy case",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU2NTU3MDAsInVzZXJfaWQiOiJ1c2VyXzEifQ.L9ZKMYUFJHl4CnYti5MM_jkUt9zA0CGsVmbqOR-zhvU",
		nil,
		"user_1",
	},
	{
		"test wrong token",
		"wrong token",
		fmt.Errorf("token contains an invalid number of segments"),
		"",
	},
	{
		"test token invalid",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU2NTU3MDAsInVzZXJfaWQiOiJ1c2VyXzEifQ.L9ZKMYUFJHl4CnYti5MM_jkUt9zA0CGsVmbqOR-zhvz",
		fmt.Errorf("signature is invalid"),
		"",
	},
	{
		"test not found userID",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHQiOjE2MTU3MDA4OTksImJhbmRhbiI6ImZpcnN0VXNlciJ9.9PYfkmjuGnMAPAFl3jy8KDjupMpmoRXw8fPfPSGSFlw",
		fmt.Errorf("not found userID"),
		"",
	},
	{
		"test not found expired time",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHQiOjE2MTU3MDA4OTksInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.nDW-my-rXmtyMdoM_tb0ePYWkto7nZ9sd8YRgayReww",
		fmt.Errorf("not found expired time"),
		"",
	},
	{
		"test not found expired time",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHQiOjE2MTU3MDA4OTksInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.nDW-my-rXmtyMdoM_tb0ePYWkto7nZ9sd8YRgayReww",
		fmt.Errorf("not found expired time"),
		"",
	},
	{
		"test token is expired",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjEsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.GAACfOdc2IXFv2osUm_2O2IV3XT1lm0bxAurFX1G74s",
		fmt.Errorf("Token is expired"),
		"",
	},
}

func TestValidToken(t *testing.T) {
	srv := NewToDoService(WithStore(nil))
	Now = func() time.Time {
		return time.Date(2021, 03, 14, 0, 0, 0, 0, time.Local)
	}
	jwt.TimeFunc = Now
	for _, tt := range validTokenTestInput {
		t.Run(tt.name, func(t *testing.T) {
			userID, err := srv.ValidToken(tt.token, jwtKeyTest)
			if (err != tt.err) && (err == nil || tt.err == nil || err.Error() != tt.err.Error()) {
				t.Errorf("Error: got %s, want %s", err, tt.err)
			}
			if len(userID) != len(tt.userID) {
				t.Errorf("Taks: got %+v, want %+v", userID, tt.userID)
			}
		})
	}
}

//1615654800
//1615655700
