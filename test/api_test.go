package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/manabie-com/togo/validate"
)

func Test_API_Login(t *testing.T) {

	var (
		uname  = "francis"
		pwd    = "1234567"
		method = http.MethodGet
	)

	req, err := http.NewRequest(method, path["login"], nil)
	if err != nil {
		panic(err)
	}

	params := req.URL.Query()
	params.Add("user_id", uname)
	params.Add("password", pwd)
	req.URL.RawQuery = params.Encode()

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	client.CloseIdleConnections()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---response: ", string(body))

	if res.StatusCode != 200 {
		t.Errorf("Test Login Failed")
	}

	defer res.Body.Close()
}

func Test_API_GetTasks(t *testing.T) {

	var (
		uname      = "notail"
		method     = http.MethodGet
		headerName = "Authorization"
	)

	req, err := http.NewRequest(method, path["tasks"], nil)
	if err != nil {
		fmt.Println("err:", err)
	}

	token, err := validate.CreateToken(uname)
	if err != nil {
		panic(err)
	}

	req.Header.Add(headerName, token)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	client.CloseIdleConnections()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---response: ", string(body))

	if res.StatusCode != 200 {
		t.Errorf("Test Login failed")
	}

	defer res.Body.Close()
}

func Test_API_AddTask(t *testing.T) {

	var (
		method     = http.MethodPost
		headerName = "Authorization"
		content    = "hello world"
		uname      = "notail"
	)

	payload := make(map[string]string)
	payload["content"] = content
	data, _ := json.Marshal(payload)

	req, err := http.NewRequest(method, path["tasks"], bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("err:", err)
	}

	token, err := validate.CreateToken(uname)
	if err != nil {
		panic(err)
	}

	req.Header.Add(headerName, token)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	client.CloseIdleConnections()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---response: ", string(body))

	if res.StatusCode != 200 {
		t.Errorf("Test Login failed")
	}

	defer res.Body.Close()
}
