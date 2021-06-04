package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

//Login by existing user using POST
func TestServerHTTP_LoginSuccess_POST(t *testing.T) {
	//Setup DB Connection
	store := DBConn()

	//Todo struct
	todo := ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  &store,
	}

	//Create http test mock
	mux := http.NewServeMux()
	mux.HandleFunc("/", todo.ServeHTTP)

	//login user
	data := url.Values{}
	data.Set("user_id", "firstUser")
	data.Set("password", "example")

	//Create response and resquest
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	mux.ServeHTTP(resp, req)

	//Checking response
	assert := assert.New(t)
	assert.Equal(resp.Code, 200, "Login Success")
}

//Login by existing user using GET
func TestServerHTTP_LoginSuccess_GET(t *testing.T) {
	//Setup DB Connection
	store := DBConn()

	//Todo struct
	todo := ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  &store,
	}

	//Create http test mock
	mux := http.NewServeMux()
	mux.HandleFunc("/", todo.ServeHTTP)

	//Create response and resquest
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)

	//login user
	q := req.URL.Query()
	q.Add("user_id", "firstUser")
	q.Add("password", "example")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	mux.ServeHTTP(resp, req)

	//Checking response
	assert := assert.New(t)
	assert.Equal(resp.Code, 200, "Login Success")
}

//Login by non-existing user
func TestServerHTTP_LoginFail(t *testing.T) {
	//Setup DB Connection
	store := DBConn()

	//Todo struct
	todo := ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  &store,
	}

	//Create http test mock
	mux := http.NewServeMux()
	mux.HandleFunc("/", todo.ServeHTTP)

	//login user
	data := url.Values{}
	data.Set("user_id", "NonExistingUser")
	data.Set("password", "example")

	//Create response and resquest
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	mux.ServeHTTP(resp, req)

	//Checking response
	assert := assert.New(t)
	assert.NotEqual(resp.Code, 200, "Login Fail")
}

//List logged in user's task
func TestServerHTTP_ListTasks(t *testing.T) {
	//Setup DB Connection
	store := DBConn()

	//Todo struct
	todo := ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  &store,
	}

	//Create http test mock
	mux := http.NewServeMux()
	mux.HandleFunc("/", todo.ServeHTTP)

	//login user
	data := url.Values{}
	data.Set("user_id", "firstUser")
	data.Set("password", "example")

	//Create response and resquest
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	mux.ServeHTTP(resp, req)

	//Checking log in response
	assert := assert.New(t)
	assert.Equal(resp.Code, 200, "Login Success")

	body := map[string]interface{}{}
	json.NewDecoder(resp.Result().Body).Decode(&body)

	//Get user token
	token := fmt.Sprintf("%v", body["data"])

	//Create response and request for list task
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/tasks", nil)

	//Add parameter
	q := req.URL.Query()
	q.Add("created_date", "2020-06-29")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	mux.ServeHTTP(resp, req)

	//Checking if tasks listed
	assert.Equal(resp.Code, 200, "Tasks listed")

}

//Add new task to logged in user
func TestServerHTTP_AddTasks(t *testing.T) {
	//Setup DB Connection
	store := DBConn()

	//Todo struct
	todo := ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  &store,
	}

	//Create http test mock
	mux := http.NewServeMux()
	mux.HandleFunc("/", todo.ServeHTTP)

	//login user
	data := url.Values{}
	data.Set("user_id", "firstUser")
	data.Set("password", "example")

	//Create response and resquest
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	mux.ServeHTTP(resp, req)

	//Checking log in response
	assert := assert.New(t)
	assert.Equal(resp.Code, 200, "Login Success")

	//Get user token
	body := map[string]interface{}{}
	json.NewDecoder(resp.Result().Body).Decode(&body)
	token := fmt.Sprintf("%v", body["data"])

	//new task content
	jsonStr := []byte(`{"content":"content Task."}`)

	//Create response and request for add task
	resp = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonStr))

	req.Header.Set("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	mux.ServeHTTP(resp, req)

	// Checking if new task added
	assert.Equal(resp.Code, 200, "Task added")
}

//Request list task or add task using invalid token
func TestServerHTTP_TaskWithInvalidToken(t *testing.T) {
	//Setup DB Connection
	store := DBConn()

	//Todo struct
	todo := ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  &store,
	}

	// //Create http test mock
	mux := http.NewServeMux()
	mux.HandleFunc("/", todo.ServeHTTP)

	//new task content
	data := url.Values{}
	data.Set("content", "NewTask")

	//Create response and request for add task
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(data.Encode()))

	req.Header.Set("Authorization", "xxxxxx")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	mux.ServeHTTP(resp, req)

	//Checking if new task added
	assert := assert.New(t)
	assert.Equal(resp.Code, 401, "Unauthorized")
}

//Setup database connection
func DBConn() sqllite.LiteDB {
	db, err := sql.Open("sqlite3", "../../data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	store := sqllite.LiteDB{
		DB: db,
	}

	return store
}
