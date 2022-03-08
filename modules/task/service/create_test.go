package service

import (
	"github.com/khoale193/togo/models"
	"testing"

	"github.com/khoale193/togo/models/dbcon"
	"github.com/khoale193/togo/models/migration"
)

func init() {
	dbcon.SetupTest()
	migration.Migrate()
}

func TestCreateTask(t *testing.T) {
	(&Task{MemberID: 1, Name: "Task Name"}).CreateTask()
	task, _ := (models.Task{}).GetLatestInserted()
	if task.MemberID != 1 || task.Name != "Task Name" {
		t.Logf("Task inserted wrong.")
		t.Fail()
	}
}
