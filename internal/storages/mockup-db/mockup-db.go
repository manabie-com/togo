package mockup_db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/manabie-com/togo/internal/storages"
	"sync"
)

type MockupDB struct{
	mux sync.Mutex
}

var mockupData = map[string](map[string]interface{}){
	"firstUser": {
		"password": "example",
		"max_todo": 5,
		"tasks": []*storages.Task{
			{
				ID:          "e1da0b9b-7ecc-44f9-82ff-4623cc50446a",
				Content:     "first content",
				UserID:      "firstUser",
				CreatedDate: "2020-06-29",
			},
			{
				ID:          "055261ab-8ba8-49e1-a9e8-e9f725ba9104",
				Content:     "second content",
				UserID:      "firstUser",
				CreatedDate: "2020-06-29",
			},
			{
				ID:          "2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a",
				Content:     "another content",
				UserID:      "firstUser",
				CreatedDate: "2020-06-29",
			},
		},
	},
	"secondUser": {
		"password": "example",
		"max_todo": 5,
		"tasks": []*storages.Task{
			{
				ID:          "id_1",
				Content:     "content_1",
				UserID:      "secondUser",
				CreatedDate: "2020-06-29",
			},
			{
				ID:          "id_2",
				Content:     "content_2",
				UserID:      "secondUser",
				CreatedDate: "2020-06-29",
			},
			{
				ID:          "id_3",
				Content:     "content_3",
				UserID:      "secondUser",
				CreatedDate: "2020-06-29",
			},
			{
				ID:          "id_4",
				Content:     "content_4",
				UserID:      "secondUser",
				CreatedDate: "2020-06-29",
			},
			{
				ID:          "id_5",
				Content:     "content_5",
				UserID:      "secondUser",
				CreatedDate: "2020-06-29",
			},
		},
	},
	"thirdUser": {
		"password": "example",
		"max_todo": 5,
		"tasks":    []*storages.Task{},
	},
}

func (m *MockupDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	data, ok := mockupData[userID.String]
	if !ok {
		return nil, nil
	}
	return data["tasks"].([]*storages.Task), nil
}

func (m *MockupDB) AddTask(ctx context.Context, t *storages.Task) error {
	if t == nil {
		return errors.New("nil pointer")
	}

	m.mux.Lock()
	defer m.mux.Unlock()

	data, ok := mockupData[t.UserID]
	if !ok {
		return nil
	}
	tasks := data["tasks"].([]*storages.Task)
	tasks = append(tasks, t)
	data["tasks"] = tasks
	return nil
}

func (m *MockupDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	m.mux.Lock()
	defer m.mux.Unlock()

	data, ok := mockupData[userID.String]
	if !ok {
		return false
	}
	return data["password"].(string) == pwd.String
}

func (m *MockupDB) GetUserInfo(ctx context.Context, userID sql.NullString) *storages.User {
	m.mux.Lock()
	defer m.mux.Unlock()

	data, ok := mockupData[userID.String]
	if !ok {
		return nil
	}
	return &storages.User{
		ID:       userID.String,
		Password: data["password"].(string),
		MaxTodo:  data["max_todo"].(int),
	}
}

func (m *MockupDB) CountTasks(ctx context.Context, userID, createdDate sql.NullString) (int, error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	data, ok := mockupData[userID.String]
	if !ok {
		return -1, nil
	}
	tasks := data["tasks"].([]*storages.Task)
	if tasks == nil {
		return 0, nil
	}
	return len(tasks), nil
}
