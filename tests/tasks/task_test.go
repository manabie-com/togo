package tasks_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
	"todo/database"
	"todo/modules/tasks"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	jwtware "github.com/gofiber/jwt/v3"
)

type mockresponstory struct {
	tasks []tasks.Tasks
}

func InitMockresponstory(tasks []tasks.Tasks) database.Responstory {
	return mockresponstory{tasks: tasks}
}

func (mock mockresponstory) Get(data interface{}, id string) error {
	for _, value := range mock.tasks {
		if fmt.Sprint(value.Id) == id {
			strvalue, _ := json.Marshal(value)
			json.Unmarshal(strvalue, data)
			break
		}
	}
	return nil
}

func (mock mockresponstory) Insert(data interface{}) error {
	newTask := reflect.ValueOf(data).Interface().(*tasks.Tasks)
	newTaskData := *newTask
	mock.tasks = append(mock.tasks, newTaskData)
	return nil
}

func (mock mockresponstory) GetAll(data interface{}) error {
	data = mock.tasks
	return nil
}

func (mock mockresponstory) Find(data interface{}, query string, args ...interface{}) error {
	strvalue, _ := json.Marshal(mock.tasks)
	json.Unmarshal(strvalue, data)
	return nil
}

func TestGet(t *testing.T) {
	tasksdata := []tasks.Tasks{{
		Id:         1,
		Title:      "test get",
		Desciption: "test get",
	}}
	mock := InitMockresponstory(tasksdata)
	controller := tasks.InitTaskController(mock)

	app := fiber.New()
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Get("/:id", controller.Get)

	req := httptest.NewRequest(http.MethodGet, "/1", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibGltaXQiOjEsIm5hbWUiOiJ1c2VyNSJ9.6wCtOxIddQB1uLMGRtFD-JT6Ft-xy23Yyn6xCVXENq8")

	resp, err := app.Test(req, 1)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, `{"data":{"task":{"id":1,"isActive":false,"title":"test get","description":"test get","createdAt":0,"createdBy":0,"updatedAt":0,"updatedBy":0}},"success":true}`, string(body))
}

func TestInsert(t *testing.T) {
	tasksdata := []tasks.Tasks{}
	newtask := tasks.TasksCreate{
		Title:       "test insert",
		Discription: "test insert",
	}
	mock := InitMockresponstory(tasksdata)
	controller := tasks.InitTaskController(mock)

	app := fiber.New()
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Post("/", controller.Create)
	body, _ := json.Marshal(newtask)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibGltaXQiOjEsIm5hbWUiOiJ1c2VyNSJ9.6wCtOxIddQB1uLMGRtFD-JT6Ft-xy23Yyn6xCVXENq8")

	resp, err := app.Test(req, 3)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	if status := resp.StatusCode; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	bodyRes, err := ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, `{"data":{"task":{"id":0,"isActive":false,"title":"test insert","description":"test insert","createdAt":0,"createdBy":6,"updatedAt":0,"updatedBy":0}},"success":true}`, string(bodyRes))
}

func TestInsertMore1(t *testing.T) {
	tasksdata := []tasks.Tasks{{
		Id:         1,
		Title:      "test insert",
		Desciption: "test insert",
		CreatedBy:  6,
		CreatedAt:  time.Now().Unix(),
	}}
	newtask := tasks.TasksCreate{
		Title:       "test insert",
		Discription: "test insert",
	}

	mock := InitMockresponstory(tasksdata)
	controller := tasks.InitTaskController(mock)

	app := fiber.New()
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Post("/", controller.Create)
	body, _ := json.Marshal(newtask)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibGltaXQiOjEsIm5hbWUiOiJ1c2VyNSJ9.6wCtOxIddQB1uLMGRtFD-JT6Ft-xy23Yyn6xCVXENq8")
	resp, err := app.Test(req, 3)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	bodyRes, err := ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, `{"message":"User had task more then limmit perday","success":false}`, string(bodyRes))
}
