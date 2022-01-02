package user

import (
	"context"
	"sync"

	"github.com/manabie-com/togo/app/common/adapter"
	"github.com/manabie-com/togo/app/common/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	cfg          = config.GetConfig()
	Once         sync.Once
	CollInstance adapter.CollectionV2
)

const (
	dbName   string = "app"
	collName string = "User"
)

// InitColletion init db.collection, required index and return collection instance
func InitColletion() *adapter.CollectionV2 {
	// Init only one time
	Once.Do(func() {

		mongoInstance := cfg.Mongo.Get(dbName)
		mongoInstance.Init()
		CollInstance = mongoInstance.GetCollectionV2(collName)

		// init indexes
		requiredIndexes := []mongo.IndexModel{
			{
				Keys: bsonx.Doc{
					{Key: "username", Value: bsonx.Int32(1)},
				},
				Options: options.Index().
					SetBackground(true),
			},
			{
				Keys: bsonx.Doc{
					{Key: "status", Value: bsonx.Int32(1)},
				},
				Options: options.Index().
					SetBackground(true),
			},
		}

		go initIndexes(requiredIndexes)
	})

	// Get MongoV2
	return &CollInstance
}

func initIndexes(indexKeys []mongo.IndexModel) {
	ctx := context.Background()

	for _, indexKey := range indexKeys {
		_, err := CollInstance.Indexes().CreateOne(ctx, indexKey)
		if err != nil {
			panic(err)
		}
	}
}
