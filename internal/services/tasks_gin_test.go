package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/jericogantuangco/togo/internal/storages/mock"
	"github.com/jericogantuangco/togo/internal/storages/postgres"
	"github.com/jericogantuangco/togo/internal/util"
	"github.com/stretchr/testify/require"
)

func TestListTaskAPIOK(t *testing.T) {
	now := time.Now().Format("2006-01-02")
	numberOf := 5
	tasks := make([]postgres.Task, numberOf)
	for i := 0; i < numberOf; i++ {
		tasks[i] = postgres.Task{
			ID:          util.RandomInt(1, 50),
			Content:     util.RandomString(6),
			UserID:      util.RandomString(6),
			CreatedDate: now,
		}
	}
	//create mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	arg := postgres.RetrieveTasksParams{
		UserID:      "testUser",
		CreatedDate: now,
	}

	//create db mocking
	store := mockdb.NewMockStore(ctrl)
	// mock the db
	store.EXPECT().
		RetrieveTasks(gomock.Any(), gomock.Eq(arg)).
		Times(1).
		Return(tasks, nil)
	server, err := NewServer(store)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	url := "/tasks"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	q := req.URL.Query()

	q.Add("created_date", now)
	req.URL.RawQuery = q.Encode()
	addAuth(t, req, server.TokenMaker, authorizationTypeBearer, "testUser", 10*time.Minute)
	server.Router.ServeHTTP(recorder, req)
	checkResponse(t, recorder)
}

func TestAddTaskAPIOK(t *testing.T) {
	now := time.Now().Format("2006-01-02")
	user := postgres.User{
		Username: "testUser",
		Password: "password",
		MaxTodo: 5,
	}
	task := postgres.Task{
		ID:          util.RandomInt(1, 50),
		Content:     util.RandomString(6),
		UserID:      util.RandomString(6),
		CreatedDate: now,
	}
	numberOf := 2
	tasks := make([]postgres.Task, numberOf)
	for i := 0; i < numberOf; i++ {
		tasks[i] = postgres.Task{
			ID:          util.RandomInt(1, 50),
			Content:     util.RandomString(6),
			UserID:      util.RandomString(6),
			CreatedDate: now,
		}
	}

	//create mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	arg1 := postgres.CreateTaskParams{
		Content:     task.Content,
		UserID:      task.UserID,
		CreatedDate: task.CreatedDate,
	}

	arg2 := postgres.RetrieveTasksParams{
		UserID: task.UserID,
		CreatedDate: task.CreatedDate,
	}
	//create db mocking
	store := mockdb.NewMockStore(ctrl)
	// mock the db
	store.EXPECT().
		RetrieveUser(gomock.Any(), gomock.Eq(arg1.UserID)).
		Times(1).
		Return(user, nil)
	store.EXPECT().
		RetrieveTasks(gomock.Any(), gomock.Eq(arg2)).
		Times(1).
		Return(tasks, nil)
	store.EXPECT().
		CreateTask(gomock.Any(), gomock.Eq(arg1)).
		Times(1).
		Return(task, nil)
	server, err := NewServer(store)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	data, err := json.Marshal(gin.H{
		"content": task.Content,
	})
	require.NoError(t, err)

	url := "/tasks"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	addAuth(t, req, server.TokenMaker, authorizationTypeBearer, task.UserID, 10*time.Minute)
	server.Router.ServeHTTP(recorder, req)
	checkResponse(t, recorder)
}
