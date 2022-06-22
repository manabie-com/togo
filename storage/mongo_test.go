package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)



func Test_StartMongo(t *testing.T){
	client := StartMongo()
	assert.IsType(t, &mongo.Client{}, client)
}