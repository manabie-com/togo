package task

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_NewRepository(t *testing.T) {
	repo := NewRepository()
	assert.IsType(t, &RecordRepository{}, repo)
}

func Test_InsertUserTask(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		RecordCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		repo := NewRepository()
		err := repo.InsertUserTask("1", "todo", time.Now())
		assert.Nil(t, err)
	})
	mt.Run("failed with internal mongo error", func(mt *mtest.T) {
		RecordCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "internal mongo error",
		}))
		repo := NewRepository()
		err := repo.InsertUserTask("1", "todo", time.Now())
		assert.ErrorContains(t, err, "insert user task error")
	})
}

func Test_GetUserConfig(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		UserCollection = mt.Coll
		expectedUserConfig := &UserConfig{
			UserId:  "1",
			Limit: 3,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{"user_id", expectedUserConfig.UserId},
			{"limit", expectedUserConfig.Limit},
		}))
		repo := NewRepository()
		user, err := repo.GetUserConfig("1")
		assert.Nil(t, err)
		assert.Equal(t, expectedUserConfig, user)
	})
	mt.Run("failed with mongo error", func(mt *mtest.T) {
		UserCollection = mt.Coll
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
			Index:   1,
			Code:    11000,
			Message: "internal mongo error",
		}))
		repo := NewRepository()
		user, err := repo.GetUserConfig("1")
		assert.NotNil(t, err)
		assert.Nil(t, user)
	})
} 