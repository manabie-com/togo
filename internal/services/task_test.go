package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"reflect"

	"time"

	// "net/url"
	"testing"

	//
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/manabie-com/togo/internal/storages"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
)

func InitialTestEnv() *ToDoService {
	// Initial Test environment - note that the database name -> "manabie_test"
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=admin password=admin123 dbname=manabie_test sslmode=disable")

	fmt.Println("Initial Server!")

	if err != nil {
		fmt.Println("error opening db", err.Error())
	}

	// Initial table for database
	db.AutoMigrate(&storages.Task{})
	db.AutoMigrate(&storages.User{})
	db.AutoMigrate(&storages.ConfigServer{})

	// Create user for testing
	var user = storages.User{
		ID:       "userTest",
		Password: "testExample",
		CurrentNumberTask: 3,
	}

	if db.Where("id=?", user.ID).First(&user).RowsAffected <= 0 {
		rs := db.Create(&user)

		if rs.Error != nil {
			fmt.Println(rs.Error.Error())
		}
	}

	// Initial config server for processing max tasks / day of user for testing

	var configServers = []storages.ConfigServer{{
		Name:  "config_max_tasks_per_day",
		Value: 5,
	}}

	for index := range configServers {
		var config = configServers[index]
		if db.Where("name=?", config.Name).First(&config).RowsAffected <= 0 {
			db.Create(&config)
		} else {
			db.Model(&storages.ConfigServer{}).Where("name=?", configServers[index].Name).Update("value", configServers[index].Value)
		}
	}

	now := time.Now()

	// Initial some tasks data of user
	var tasks = []storages.Task{
		{
			ID:          uuid.New().String(),
			CreatedDate: now.Format("2006-01-02"),
			UserID:      "userTest",
			Content:     "Some Content",
		},
		{
			ID:          uuid.New().String(),
			CreatedDate: now.Format("2006-01-02"),
			UserID:      "userTest",
			Content:     "Another Content",
		},
		{
			ID:          uuid.New().String(),
			CreatedDate: now.Format("2006-01-02"),
			UserID:      "userTest",
			Content:     "Add Content",
		},
	}

	for index := range tasks {
		db.Create(&tasks[index])
	}

	var configTask storages.ConfigServer
	rsConfig := db.Where("name=?", "config_max_tasks_per_day").First(&configTask)
	if rsConfig.Error != nil {
		fmt.Println(rsConfig.Error.Error())
	}

	var serviceTest = ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PostgresDB{
			DB: db,
		},
		MaxTasksConfig: configTask.Value,
	}

	return &serviceTest
}

func ClearTestData(testEnv *ToDoService) {
	defer testEnv.Store.DB.Close()
	// Clear db
	defer testEnv.Store.DB.Model(&storages.User{}).Delete(&storages.User{})
	defer testEnv.Store.DB.Model(&storages.Task{}).Delete(&storages.Task{})
	defer testEnv.Store.DB.Model(&storages.ConfigServer{}).Delete(&storages.ConfigServer{})
}

type TokenData struct {
	data string
}

// TestLogin for check the authentication
func TestLogin(t *testing.T) {

	var testEnv = InitialTestEnv()

	defer ClearTestData(testEnv)

	// Create new request to "/login" endpoint
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	//
	// Data Input for test case
	user := make(map[string]string)
	user["user_id"] = "userTest"
	user["password"] = "testExample"

	q := req.URL.Query()
	q.Add("user_id", user["user_id"])
	q.Add("password", user["password"])

	//
	//
	req.URL.RawQuery = q.Encode()

	// Creates a new recorder to record the response received by the "/login" endpoint.
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testEnv.ServeHTTP)
	handler.ServeHTTP(rr, req)

	// Mapping for any type object -> use json.Unmarshal for get data.
	var result map[string]interface{}
	err = json.Unmarshal([]byte(strings.TrimSpace(rr.Body.String())), &result)
	if err != nil {
		t.Error(err.Error())
	}
	status := rr.Code

	// If want to get any data => just `token["data"]`
	// Unauthorization case
	if (status == http.StatusUnauthorized) && (result["error"] == "incorrect user_id/pwd") {
		t.Errorf("Case Unauthorization wrong UserID: %v or Password: %v - Error: %v", user["user_id"], user["password"], result["error"])
	}

	// Can not sign token case
	if (status == http.StatusInternalServerError) && (reflect.TypeOf(result["error"]).String() == "string") {
		t.Errorf("Case Internal server error because can not sign Token: %v", result["error"])
	}

	// Success Case
	if (status == http.StatusOK) && (reflect.TypeOf(result["data"]).String() == "string") {
		t.Logf("Case success login, result Token: %v", result["data"])
	} else {
		t.Errorf("Internal server error, You must find the bug!")
	}
}

// TestListTask for get the list of tasks of user
func TestListTask(t *testing.T) {
	var testEnv = InitialTestEnv()

	defer ClearTestData(testEnv)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU3ODgwNzEsInVzZXJfaWQiOiJ1c2VyVGVzdCJ9.JIyIXpdIjQQMLnEr5WCh_E0rhNt0_B5yMhrFDAsgemw"

	// Create new request to "/tasks" endpoint
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("authorization", token)

	now := time.Now()
	q := req.URL.Query()
	q.Add("created_date", now.Format("2006-01-02"))
	//
	//
	req.URL.RawQuery = q.Encode()

	// Creates a new recorder to record the response received by the "/login" endpoint.
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testEnv.ServeHTTP)
	handler.ServeHTTP(rr, req)

	// Mapping for any type object -> use json.Unmarshal for get data.
	var result map[string]interface{}
	err = json.Unmarshal([]byte(strings.TrimSpace(rr.Body.String())), &result)
	status := rr.Code
	if err != nil {
		fmt.Println(err.Error())
	}

	// Case not input or Invalid Token
	if status == http.StatusUnauthorized {
		t.Error("Case Invalid Token - or - Expired Token")
	}

	// Case internal server error
	if (status == http.StatusInternalServerError) && (reflect.TypeOf(result["error"]).String() == "string") {
		t.Errorf("Internal Server Error, When get list tasks: %v", result["error"])
	}

	if (status == http.StatusOK) && ((result["data"] == nil) || (reflect.TypeOf(result["data"]).String() == "[]interface {}")) {
		t.Logf("Success get the data: %v", result["data"])
	} else {
		t.Errorf("Internal server error, you must find the bug")
	}
}

// TestCreateTask test create new task
func TestCreateTask(t *testing.T) {
	var testEnv = InitialTestEnv()

	defer ClearTestData(testEnv)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU3ODgwNzEsInVzZXJfaWQiOiJ1c2VyVGVzdCJ9.JIyIXpdIjQQMLnEr5WCh_E0rhNt0_B5yMhrFDAsgemw"

	var jsonTaskStr = []byte(`{"content":"manabie content"}`)
	// Create new request to "/tasks" endpoint
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonTaskStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("authorization", token)
	req.Header.Set("Content-Type", "application/json")

	// Creates a new recorder to record the response received by the "/login" endpoint.
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testEnv.ServeHTTP)
	handler.ServeHTTP(rr, req)

	// Mapping for any type object -> use json.Unmarshal for get data.
	var result map[string]interface{}
	err = json.Unmarshal([]byte(strings.TrimSpace(rr.Body.String())), &result)
	status := rr.Code
	if err != nil {
		fmt.Println(err.Error())
	}
	// Case not input or Invalid Token
	if status == http.StatusUnauthorized {
		t.Error("Case Invalid Token - or - Expired Token")
	}

	if status == http.StatusInternalServerError {
		t.Error("StatusInternalServerError Cannot Decode Body")
	}

	if status == http.StatusNotAcceptable && (result["error"] == "User tasks are limited, cant add more, comeback tomorrow") {
		t.Errorf("User has limit to add new task: %v", result["error"])
	}

	if status == http.StatusInternalServerError && (reflect.TypeOf(result["error"]).String() == "string") {
		t.Errorf("User add task get error: %v", result["error"])
	}

	if (status == http.StatusOK) && ((result["data"] == nil) || (reflect.TypeOf(result["data"]).String() == "map[string]interface {}")) {
		t.Logf("Success get the data: %v", result["data"])
	} else {
		t.Errorf("Internal server error, you must find the bug")
	}
}
