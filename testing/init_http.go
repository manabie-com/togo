package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

var (
	client = &http.Client{}
	path   = initPath()
	token  = initToken()
)

func initPath() string {
	path := fmt.Sprintf("http://localhost:7171/local/api")
	return path
}
func initToken() string {
	path := initPath()
	url := fmt.Sprintf("%s/auth/login", path)
	var jsonStr = bytes.NewBufferString(`{"email":"admin@gmail.com", "password":"123456"}`)
	req, _ := http.NewRequest("POST", url, jsonStr)
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var user struct {
		Data struct {
			Token string `json:"token"`
		}
	}
	json.Unmarshal(body, &user)
	return user.Data.Token
}

func InitHttp(path string, h func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(path, h)

	return server
}
