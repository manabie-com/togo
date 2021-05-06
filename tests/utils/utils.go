package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

const (
	GET        = "GET"
	POST       = "POST"
	PUT        = "PUT"
	DELETED    = "DELETED"
	NONE_TOKEN = ""
	OLD_JWT    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjAyMDE2MTcsInVzZXJfaWQiOiJmaXJzdFVzZXIifQ.BYER5lzX8FEKaga98ph2i0kWnV-UuIOQB8mzMF_Jz3Y"
)

const USERID = "firstUser"
const PASSWORD = "example"

func Install() (*httptest.Server, *services.ToDoService, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	serv := &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}
	serv.Store.CreateDatabase()
	httpServer := httptest.NewServer(serv)
	return httpServer, serv, err
}

func CreateRequest(method string, url string, token string, jsonData []byte, params ...map[string]string) (int, map[string]interface{}, error) {
	var body map[string]interface{}
	client := &http.Client{}
	var req *http.Request
	var err error

	if nil != jsonData {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	switch method {
	case GET:

		if nil != params && len(params) > 0 {
			query := req.URL.Query()
			for key, val := range params[0] {
				query.Add(key, val)
			}
			req.URL.RawQuery = query.Encode()
		}
		break
	}
	if "" != token {
		req.Header.Set("Authorization", token)
	}
	res, err := client.Do(req)
	json.NewDecoder(res.Body).Decode(&body)

	return res.StatusCode, body, err
}
