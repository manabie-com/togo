package task

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/tests/utils"
	_ "github.com/mattn/go-sqlite3"
)

const USERID = "firstUser"
const PASSWORD = "example"

func GetToken(userId string, password string) (*httptest.Server, *services.ToDoService, int, string, error) {
	httpServer, serv, err := utils.Install()

	if nil != err {
		return httpServer, serv, 0, "", err
	}

	status, body, err := utils.CreateRequest(utils.GET, httpServer.URL+"/login", "", nil, map[string]string{"user_id": userId, "password": password})
	token := ""
	if nil != body["data"] {
		token = fmt.Sprintf("%s", body["data"])
	}
	return httpServer, serv, status, token, err
}

func PostTask(url string, token string, content string) (int, storages.Task, error) {
	var body map[string]interface{}

	values := map[string]string{"content": content}
	jsonData, _ := json.Marshal(values)
	status, body, err := utils.CreateRequest(utils.POST, url, token, jsonData)
	task := storages.Task{}
	dataByte, _ := json.Marshal(body["data"])
	json.Unmarshal(dataByte, &task)

	return status, task, err
}

func GetTask(url string, token string) (int, map[string]interface{}, error) {
	var body map[string]interface{}
	createdDate := time.Now().Format("2006-01-02")
	status, body, err := utils.CreateRequest(utils.GET, url, token, nil, map[string]string{"created_date": createdDate})
	return status, body, err
}

func TestLogin(t *testing.T) {
	httpServer, serv, status, token, err := GetToken(USERID, PASSWORD)
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		httpServer.Close()
		serv.Store.DB.Close()
	}(httpServer, serv)
	if 200 != status {
		t.Error("Request failed")
	}
	if "" == token {
		t.Error("Token is empty")
	}
	if nil != err {
		t.Error(err.Error())
	}
}

func TestPostTask(t *testing.T) {
	httpServer, serv, status, token, err := GetToken(USERID, PASSWORD)
	if 200 != status {
		t.Error("Request failed")
	}
	if "" == token {
		t.Error("Token is empty")
	}
	if nil != err {
		t.Error(err.Error())
	}
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		serv.Store.DB.Close()
		httpServer.Close()
	}(httpServer, serv)

	content := "OcSen Hoc Code"
	status, task, err := PostTask(httpServer.URL+"/tasks", token, content)

	if 200 != status {
		t.Error("Request failed")
	}
	if USERID != task.UserID {
		t.Error("UserID is not correctly")
	}

	if content != task.Content {
		t.Error("Content is not correctly")
	}
	if nil != err {
		t.Error(err.Error())
	}
}

func TestPostTaskLimit(t *testing.T) {
	httpServer, serv, status, token, err := GetToken(USERID, PASSWORD)
	if 200 != status {
		t.Error("Request failed")
	}
	if "" == token {
		t.Error("Token is empty")
	}
	if nil != err {
		t.Error(err.Error())
	}
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		serv.Store.DB.Close()
		httpServer.Close()
	}(httpServer, serv)
	content := "OcSen Hoc Code"
	for i := 1; i <= 5; i++ {
		status, task, err := PostTask(httpServer.URL+"/tasks", token, content)
		if 200 != status {
			t.Error("Request failed")
		}
		if USERID != task.UserID {
			t.Error("UserID is not correctly")
		}

		if content != task.Content {
			t.Error("Content is not correctly")
		}
		if nil != err {
			t.Error(err.Error())
		}
	}
	status, _, err = PostTask(httpServer.URL+"/tasks", token, content)
	if 500 != status {
		t.Error("Request failed")
	}
}

func TestPostTaskWithExpireToken(t *testing.T) {
	httpServer, serv, status, token, err := GetToken(USERID, PASSWORD)
	if 200 != status {
		t.Error("Request failed")
	}
	if "" == token {
		t.Error("Token is empty")
	}
	if nil != err {
		t.Error(err.Error())
	}
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		serv.Store.DB.Close()
		httpServer.Close()
	}(httpServer, serv)

	if nil != err {
		t.Error(err.Error())
	}
	content := "OcSen Hoc Code"
	status, _, _ = PostTask(httpServer.URL+"/tasks", utils.OLD_JWT, content)

	if 401 != status || 200 == status {
		t.Error("Expired token can create task")
	}
}

func TestGetTask(t *testing.T) {
	httpServer, serv, status, token, err := GetToken(USERID, PASSWORD)
	if 200 != status {
		t.Error("Request failed")
	}
	if "" == token {
		t.Error("Token is empty")
	}
	if nil != err {
		t.Error(err.Error())
	}
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		serv.Store.DB.Close()
		httpServer.Close()
	}(httpServer, serv)

	if nil != err {
		t.Error(err.Error())
	}
	content := "OcSen Hoc Code"
	PostTask(httpServer.URL+"/tasks", token, content)
	GetTask(httpServer.URL+"/tasks", token)
}

func TestGetTaskWithExpireToken(t *testing.T) {
	httpServer, serv, status, token, err := GetToken(USERID, PASSWORD)
	if 200 != status {
		t.Error("Request failed")
	}
	if "" == token {
		t.Error("Token is empty")
	}
	if nil != err {
		t.Error(err.Error())
	}
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		serv.Store.DB.Close()
		httpServer.Close()
	}(httpServer, serv)

	if nil != err {
		t.Error(err.Error())
	}
	content := "OcSen Hoc Code"
	PostTask(httpServer.URL+"/tasks", token, content)
	status, _, _ = GetTask(httpServer.URL+"/tasks", utils.OLD_JWT)
	if 401 != status || 200 == status {
		t.Error("Expired token can get task")
	}
}
