package persistent_test

import (
	"context"
	"testing"
	"togo/domain/model"
	"togo/infrastructure/persistent"
)

func TestTaskMySQLRepository_Create(t *testing.T) {
	repo := persistent.NewTaskMySQLRepository(db)
	err := repo.Create(context.Background(), model.Task{
		Title:       "task1",
		Description: "task1 description",
		CreatedBy:   1,
	})
	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}
	t.Log("Create task success")
}
