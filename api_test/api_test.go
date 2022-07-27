package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"togo-thdung002/entities"
)

var (
	baseURL string = "http://localhost:8800/api/togo/v1"
)

func TestCreateUser(t *testing.T) {
	user := entities.User{
		Username: "dung",
		Password: "dung123",
		Limit:    15,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(user)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodPost, baseURL+"/user/register", &buf)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	response := new(interface{})
	err = json.Unmarshal(body, response)
	if err != nil {
		t.Fail()
	}
	t.Log(*response)

	assert.Equal(t, 200, resp.StatusCode)

}

func TestCreateTask(t *testing.T) {
	task := entities.Task{
		Content: "Task number 1",
		UserID:  1,
		Date:    "25-06-2022",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(task)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	request, _ := http.NewRequest(http.MethodPost, baseURL+"/task", &buf)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	response := new(interface{})
	err = json.Unmarshal(body, response)
	if err != nil {
		t.Fail()
	}
	t.Log(*response)

	assert.Equal(t, 200, resp.StatusCode)

}

func TestCreateBatchTask(t *testing.T) {
	//create batch task to get error
	for i := 1; i <= 10; i++ {
		task := entities.Task{
			Content: "Task number " + strconv.Itoa(i),
			UserID:  1,
			Date:    "15-07-2022",
		}
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(task)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		request, _ := http.NewRequest(http.MethodPost, baseURL+"/task", &buf)
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
		response := new(interface{})
		err = json.Unmarshal(body, response)
		if err != nil {
			t.Fail()
		}
		t.Log(*response)
		assert.Equal(t, 200, resp.StatusCode)

	}

}
