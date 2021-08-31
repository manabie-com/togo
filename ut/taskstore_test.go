package ut

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/repo"
	"github.com/manabie-com/togo/utils"
)

func InitTaskStore() *repo.TaskStore {
	db, err := utils.InitDB()
	if err != nil {
		log.Fatal("error opening db", err)
	}
	return &repo.TaskStore{
		DB: db,
	}
}

func TestRetrieveTasks(t *testing.T) {
	c := InitTaskStore()
	created_date := sql.NullString{
		String: "2021-08-29",
		Valid:  true,
	}
	ctx := context.Background()
	tasks, err := c.RetrieveTasks(ctx, "testUser", created_date)

	if err != nil {
		t.Errorf("RetrieveTasks returned err %v", err)
		return
	}

	if len(tasks) != 1 {
		t.Errorf("RetrieveTasks returned wrong number of tasks: got %v want %v", len(tasks), 1)
	}
}

func TestCountTask(t *testing.T) {
	c := InitTaskStore()
	ctx := context.Background()
	cnt, err := c.CountTask(ctx, "testUser", "2021-08-29")

	if err != nil {
		t.Errorf("RetrieveTasks returned err %v", err)
		return
	}

	if cnt != 1 {
		t.Errorf("RetrieveTasks returned wrong number of tasks: got %v want %v", cnt, 1)
	}
}

func TestAddTask(t *testing.T) {
	c := InitTaskStore()
	task := &storages.Task{
		ID:          uuid.New().String(),
		Content:     "test adding task",
		UserID:      "firstUser",
		CreatedDate: time.Now().Format("2006-01-02"),
	}

	ctx := context.Background()
	err := c.AddTask(ctx, task)
	if err != nil {
		t.Errorf("AddTask returned err %v", err)
		return
	}
}
