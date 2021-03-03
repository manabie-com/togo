package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

const (
	UserId      = "user_id"
	Password    = "password"
	CreatedDate = "2021-03-01"
)

var (
	s     *ToDoService
	Token string
)

func TestMain(m *testing.M) {
	db, err := sqlite.CreateUnitTestDB()
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	s = &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqlite.LiteDB{
			DB: db,
		},
	}

	os.Exit(m.Run())
}

func Test_loginHandler_Happy(t *testing.T) {
	reqURL := "/mock"
	reqURL += "?user_id=" + UserId
	reqURL += "&password=" + Password

	req := httptest.NewRequest("", reqURL, nil)
	respW := httptest.NewRecorder()

	s.loginHandler(respW, req)

	resp := respW.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	b, err := ioutil.ReadAll(resp.Body)
	require.Nil(t, err)
	t.Log(string(b))

	type Data struct {
		Data string
	}
	data := &Data{}
	err = json.Unmarshal(b, data)
	require.Nil(t, err)

	Token = data.Data
}

func Test_loginHandler_Edge(t *testing.T) {
	reqURL := "/mock"
	reqURL += "?user_id=" + UserId
	reqURL += "&password=" + "qweasdd"

	req := httptest.NewRequest("", reqURL, nil)
	respW := httptest.NewRecorder()

	s.loginHandler(respW, req)

	resp := respW.Result()
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_getTaskHandler(t *testing.T) {
	reqURL := "/mock"
	reqURL += "?created_date=" + CreatedDate

	req := httptest.NewRequest("", reqURL, nil)
	req = req.WithContext(context.WithValue(context.Background(), userAuthKey(0), UserId))

	respW := httptest.NewRecorder()

	s.getTasksHandler(respW, req)

	resp := respW.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	b, err := ioutil.ReadAll(resp.Body)
	require.Nil(t, err)
	t.Log(string(b))

	type Data struct {
		Data []storages.Task
	}
	data := &Data{}
	err = json.Unmarshal(b, data)
	require.Nil(t, err)
	require.Equal(t, []storages.Task{{
		ID:          "task_id",
		Content:     "example_content",
		UserID:      UserId,
		CreatedDate: CreatedDate,
	}}, data.Data)
}

func Test_addTaskHandler(t *testing.T){
	reqURL := "/mock"

	random := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	content := fmt.Sprintf("random_content_%d", random)
	task := &storages.Task{
		Content:     content,
	}
	b, err := json.Marshal(task)
	require.Nil(t, err)

	req := httptest.NewRequest("", reqURL, nil)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	respW := httptest.NewRecorder()

	s.addTaskHandler(respW, req)

	resp := respW.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	b, err = ioutil.ReadAll(resp.Body)
	require.Nil(t, err)
	t.Log(string(b))

	type Data struct {
		Data storages.Task
	}
	data := &Data{}
	err = json.Unmarshal(b, data)
	require.Nil(t, err)
	require.NotNil(t, data.Data)
	require.NotEmpty(t, data.Data.ID)
}
