package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	entity "github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/repository"
	"github.com/stretchr/testify/assert"
)

var (
	getByUserID func(userID string, createdDate string) ([]entity.Task, error)
	getAll      func(createdDate string) ([]entity.Task, error)
	getByID     func(taskID string) (*entity.Task, error)
	add         func(entity *entity.Task) (*entity.Task, error)
)

type TaskRepositoryMock struct{}

func (taskMock *TaskRepositoryMock) GetByUserID(userID string, createdDate string) ([]entity.Task, error) {
	return getByUserID(userID, createdDate)
}

func (taskMock *TaskRepositoryMock) Add(entity *entity.Task) (*entity.Task, error) {
	return add(entity)
}

func (taskMock *TaskRepositoryMock) GetAll(createdDate string) ([]entity.Task, error) {
	return getAll(createdDate)
}

func (taskMock *TaskRepositoryMock) GetByID(taskID string) (*entity.Task, error) {
	return getByID(taskID)
}

func TestAddTask_Success(t *testing.T) {
	jsonBodyArray := `[{"Id": "", "Content": "content"}]`

	var taskOfUser []entity.Task

	json.Unmarshal(bytes.NewBufferString(jsonBodyArray).Bytes(), &taskOfUser)

	getByUserID = func(userID string, createdDate string) ([]entity.Task, error) {
		return taskOfUser, nil
	}

	jsonBody := `{"content": "content"}`

	var task entity.Task

	json.Unmarshal(bytes.NewBufferString(jsonBody).Bytes(), &task)

	add = func(entity *entity.Task) (*entity.Task, error) {
		return &task, nil
	}

	repository.TaskRepo = &TaskRepositoryMock{}

	var taskHandler = &TaskHandler{}

	router := mux.NewRouter()

	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(jsonBody))

	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()

	router.HandleFunc("/tasks", taskHandler.AddTask)

	router.ServeHTTP(rr, req)

	var response map[string]*entity.Task

	json.Unmarshal(rr.Body.Bytes(), &response)

	assert.NotNil(t, rr.Body.Bytes())
	assert.EqualValues(t, http.StatusOK, rr.Code)
	assert.EqualValues(t, "content", response["data"].Content)
}
