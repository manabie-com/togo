package intergrationTest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/middlewares"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestGetAllTasksHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	var tasks []models.NewTask
	tasks = append(tasks, models.RandomNewTask(), models.RandomNewTask())
	for _, task := range tasks {
		task.UserId = AdminID
		err := bh.BaseCtrl.InsertTask(task)
		if err != nil {
			t.Fatal("insert task failed")
		}
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)
	context.Set(req, "userid", AdminID)

	r := mux.NewRouter()
	r.Use(middlewares.LoggingVerified)
	r.HandleFunc("/tasks", bh.ResponseAllTask).Methods("GET")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}

	// Read the response body.
	var returnTasks []models.Task
	if err := json.NewDecoder(w.Body).Decode(&returnTasks); err != nil {
		t.Errorf("Returned user list is invalid JSON.")
	}
	// delete after insert to check
	bh.BaseCtrl.DeleteTask(int(returnTasks[len(returnTasks)-1].Id), AdminID)
	bh.BaseCtrl.DeleteTask(int(returnTasks[len(returnTasks)-2].Id), AdminID)

	if len(returnTasks) < len(tasks) {
		t.Errorf("Returned user list is an invalid length. Expected %d. Got %d instead", len(tasks), len(returnTasks))
	}
	count := 0
	for _, returnTask := range returnTasks {
		if returnTask.Content == tasks[0].Content || returnTask.Content == tasks[1].Content {
			count++
		}
	}
	if count != len(tasks) {
		t.Fatal("Returned task list is different")
	}
}

func TestGetOneTaskHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	task := models.RandomNewTask()
	task.UserId = AdminID
	bh.BaseCtrl.InsertTask(task)
	tasks, _ := bh.BaseCtrl.GetAllTasks(AdminID)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/tasks/%v", tasks[len(tasks)-1].Id), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)
	context.Set(req, "userid", AdminID)
	context.Set(req, "id", tasks[len(tasks)-1].Id)

	r := mux.NewRouter()
	r.Use(middlewares.LoggingVerified, middlewares.MiddlewareID)
	r.HandleFunc("/tasks/{id}", bh.ResponseOneTask).Methods("GET")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}

	// Read the response body.
	var returnTask models.Task
	if err := json.NewDecoder(w.Body).Decode(&returnTask); err != nil {
		t.Errorf("Returned task is invalid JSON.")
	}

	bh.BaseCtrl.DeleteTask(tasks[len(tasks)-1].Id, AdminID)
	if strings.Compare(returnTask.Content, task.Content) != 0{
		t.Fatal("wrong task return")
	}
}

func TestDeleteTaskHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	task := models.RandomNewTask()
	task.UserId = AdminID
	bh.BaseCtrl.InsertTask(task)

	tasks, err := bh.BaseCtrl.GetAllTasks(AdminID)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE",fmt.Sprintf("/tasks/%v", tasks[len(tasks)-1].Id), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)
	context.Set(req, "id", tasks[len(tasks)-1].Id)
	context.Set(req, "userid", AdminID)

	r := mux.NewRouter()
	r.Use(middlewares.AdminVerified, middlewares.MiddlewareID)
	r.HandleFunc("/tasks/{id}", bh.DeleteFromTask).Methods("DELETE")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}

	// Read the response body.
	data, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Compare(string(data), "message: delete success") != 0 {
		t.Error("Expected message: delete success, got " + string(data))
	}
}


func TestUpdateTaskHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	task := models.RandomNewTask()
	tasks, err := bh.BaseCtrl.GetAllTasks(AdminID)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("PUT",fmt.Sprintf("/tasks/%v", tasks[len(tasks)-1].Id), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)
	context.Set(req, "id", tasks[len(tasks)-1].Id)
	context.Set(req, "userid", AdminID)

	newRequestBody, err := json.Marshal(task)
		if err != nil {
			http.Error(w, "marshal request body failed, err: "+err.Error(), http.StatusBadRequest)
			return
		}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(newRequestBody))

	r := mux.NewRouter()
	r.Use(middlewares.AdminVerified, middlewares.MiddlewareID)
	r.HandleFunc("/tasks/{id}", bh.UpdateToTask).Methods("PUT")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}
	tasks, _ = bh.BaseCtrl.GetAllTasks(AdminID)
	// Read the response body.
	var returnTasks models.NewTask
	if err := json.NewDecoder(w.Body).Decode(&returnTasks); err != nil {
		t.Errorf("Returned task is invalid JSON. Got: %s", returnTasks.Content)
	}

	if strings.Compare(returnTasks.Content, tasks[len(tasks)-1].Content) != 0{
		t.Fatal("Get return wrong task expected: "+tasks[len(tasks)-1].Content+", got: " + returnTasks.Content)
	}
}

func TestCreateTaskHandle(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal("can't load env variable")
	}
	dbConn := models.Connect(os.Getenv("DB_URI"))
	bh := controllers.NewBaseHandler(dbConn)
	defer dbConn.DB.Close()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("token", AdminToken)
	context.Set(req, "userid", AdminID)

	task := models.RandomNewTask()
	task.UserId =AdminID
	newRequestBody, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "marshal request body failed, err: "+err.Error(), http.StatusBadRequest)
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewBuffer(newRequestBody))

	r := mux.NewRouter()
	r.Use(middlewares.LoggingVerified)
	r.HandleFunc("/tasks", bh.CreateTask).Methods("POST")
	r.ServeHTTP(w, req)

	// Test that the status code is correct.
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code is invalid. Expected %d. Got %d instead", http.StatusOK, status)
	}
	users, _ := bh.BaseCtrl.GetAllTasks(AdminID)
	// Read the response body.
	var returnedTasks models.NewTask
	if err := json.NewDecoder(w.Body).Decode(&returnedTasks); err != nil {
		t.Errorf("Returned user is invalid JSON. Got: %s", returnedTasks.Content)
	}

	bh.BaseCtrl.DeleteTask(int(users[len(users)-1].Id), AdminID)

	if strings.Compare(returnedTasks.Content, users[len(users)-1].Content) != 0 {
		t.Fatal("Get return wrong user expected: " + users[len(users)-1].Content + ", got: " + returnedTasks.Content)
	}
}
