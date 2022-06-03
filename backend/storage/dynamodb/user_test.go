package dynamodb_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/go-cmp/cmp"
	"github.com/jssoriao/todo-go/storage"
	ddb "github.com/jssoriao/todo-go/storage/dynamodb"
)

type userMockDynamoDBAPI struct {
	table map[string]map[string]types.AttributeValue
}

func (m userMockDynamoDBAPI) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	// TODO: Mock the errors that can be encountered with PutItem
	partitionKey := "id"
	var id string
	err := attributevalue.Unmarshal(params.Item[partitionKey], &id)
	if err != nil {
		log.Fatal(err)
	}
	m.table[id] = params.Item
	return &dynamodb.PutItemOutput{}, nil
}

func (m userMockDynamoDBAPI) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return &dynamodb.QueryOutput{}, nil
}

func (m userMockDynamoDBAPI) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	partitionKey := "id"
	var id string
	err := attributevalue.Unmarshal(params.Key[partitionKey], &id)
	if err != nil {
		log.Fatal(err)
	}
	var item map[string]types.AttributeValue
	var ok bool
	if item, ok = m.table[id]; !ok {
		item = nil
	}
	return &dynamodb.GetItemOutput{Item: item}, nil
}

func TestCreateUser(t *testing.T) {
	table := map[string]map[string]types.AttributeValue{}
	client := userMockDynamoDBAPI{table}
	store := ddb.NewStorage(client)
	t.Run("UserCreated", func(t *testing.T) {
		user, err := store.CreateUser(storage.User{DailyLimit: 10})
		if diff := cmp.Diff(err, nil); diff != "" {
			t.Errorf("error mismatch %s", diff)
		}
		_, ok := table[user.ID]
		if diff := cmp.Diff(ok, true); diff != "" {
			t.Errorf("contains mismatch %s", diff)
		}
	})
}

func TestGetUser(t *testing.T) {
	timestamp := time.Now()
	id := "id1"
	user := storage.User{
		ID:         id,
		DailyLimit: 10,
		Created:    timestamp,
		Updated:    timestamp,
	}
	item, _ := attributevalue.MarshalMap(user)
	table := map[string]map[string]types.AttributeValue{
		id: item,
	}
	client := userMockDynamoDBAPI{table}
	store := ddb.NewStorage(client)
	t.Run("UserRetrieved", func(t *testing.T) {
		retrievedUser, err := store.GetUser(id)
		if diff := cmp.Diff(err, nil); diff != "" {
			t.Errorf("error mismatch %s", diff)
		}
		if diff := cmp.Diff(retrievedUser.ID, user.ID); diff != "" {
			t.Errorf("id mismatch %s", diff)
		}
	})
}
