package user_test

import (
	"os"
	"testing"

	userRepo "github.com/manabie-com/togo/app/repo/mongo/user"
)

var (
	// userRepoInstance support unit test function in mongo
	userRepoInstance userRepo.Repository
)

func TestMain(m *testing.M) {
	userRepoCollection := userRepo.InitColletion()
	userRepoInstance = userRepo.NewRepoManager(userRepoCollection)

	os.Exit(m.Run())
}
