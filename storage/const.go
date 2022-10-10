package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient *mongo.Client

)

const (
	TogoDbName          = "togo"
	UserTaskTableName   = "user_task"
	UserConfigTableName = "user_config"
)

const (
	MongoURI = "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	MongoClientTimeout = 10*time.Second
)
