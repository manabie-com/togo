package task_test

import (
	"os"
	"testing"

	taskRepo "github.com/manabie-com/togo/app/repo/mongo/task"
)

var (
	// taskRepoInstance support unit test function in mongo
	taskRepoInstance taskRepo.Repository
)

func TestMain(m *testing.M) {
	taskRepoCollection := taskRepo.InitColletion()
	taskRepoInstance = taskRepo.NewRepoManager(taskRepoCollection)

	os.Exit(m.Run())
}
