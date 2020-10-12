package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/assert"
)

var (
	conf     = config.GetConfig()
	userID   = "firstUser"
	password = "example"
	token    string
)

// Resp : http resp
type Resp struct {
	Data interface{} `json: "data"`
}

func TestLogin(t *testing.T) {
	assert := assert.New(t)

	endpoint := fmt.Sprintf("http://localhost:%s/login?user_id=%s&password=%s", conf.ServerPort, userID, password)
	req, err := http.NewRequest("POST", endpoint, nil)
	assert.Nil(err)

	client := &http.Client{}
	res, err := client.Do(req)
	assert.Nil(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(err)

	var resultAPI Resp
	err = json.Unmarshal(body, &resultAPI)
	assert.Nil(err)
	assert.NotNil(resultAPI.Data)

	token = resultAPI.Data.(string)
	fmt.Println(token)
}

func TestAddTask(t *testing.T) {
	assert := assert.New(t)

	endpoint := fmt.Sprintf("http://localhost:%s/tasks", conf.ServerPort)
	jsonValue, _ := json.Marshal(map[string]interface{}{
		"content": "task",
	})
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonValue))
	assert.Nil(err)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	assert.Nil(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(err)

	var resultAPI Resp
	err = json.Unmarshal(body, &resultAPI)
	assert.Nil(err)
	assert.NotNil(resultAPI.Data)

	var task storages.Task
	bs, _ := json.Marshal(resultAPI.Data)
	err = json.Unmarshal(bs, &task)
	assert.Nil(err)
	assert.NotNil(task)

	fmt.Println(task)
}

func TestListTask(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Format("2006-01-02")
	endpoint := fmt.Sprintf("http://localhost:%s/tasks?created_date=%s", conf.ServerPort, now)
	req, err := http.NewRequest("GET", endpoint, nil)
	assert.Nil(err)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	assert.Nil(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(err)

	var resultAPI Resp
	err = json.Unmarshal(body, &resultAPI)
	assert.Nil(err)
	assert.NotNil(resultAPI.Data)

	var tasks []*storages.Task
	bs, _ := json.Marshal(resultAPI.Data)
	err = json.Unmarshal(bs, &tasks)
	assert.Nil(err)
	assert.NotNil(tasks)

	fmt.Printf("%+v\n", tasks)
}
