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
	item := m.table[id]
	return &dynamodb.GetItemOutput{Item: item}, nil
}

func TestCreateUser(t *testing.T) {
	// cmpCompare := cmp.Comparer(func(x, y storage.User) bool {
	// 	return x.ID == y.ID &&
	// 		x.DailyLimit == y.DailyLimit &&
	// 		x.Created.Sub(y.Created) < time.Duration(1*time.Millisecond)
	// })

	table := map[string]map[string]types.AttributeValue{}
	client := userMockDynamoDBAPI{table}
	store := ddb.NewStorage(client)

	user, err := store.CreateUser(storage.User{DailyLimit: 10})
	if diff := cmp.Diff(err, nil); diff != "" {
		t.Errorf("error mismatch %s", diff)
	}
	v, ok := table[user.ID]
	if diff := cmp.Diff(ok, true); diff != "" {
		t.Error("user not found in table")
	}
	var _user storage.User
	attributevalue.UnmarshalMap(v, &_user)
	if diff := cmp.Diff(_user, user); diff != "" {
		t.Errorf("result mismatch %s", diff)
	}
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

	tests := []struct {
		name    string
		req     string
		want    *storage.User
		wantErr error
	}{
		{
			name:    "Found",
			req:     id,
			want:    &user,
			wantErr: nil,
		},
		{
			name:    "Not Found",
			req:     "notFound",
			want:    nil,
			wantErr: nil,
		},
	}

	cmpCompare := cmp.Comparer(func(actual, want *storage.User) bool {
		if want == nil {
			if actual != nil {
				return false
			}
			return true
		}
		return cmp.Equal(*actual, *want)
	})

	for _, test := range tests {
		retrievedUser, err := store.GetUser(test.req)
		if diff := cmp.Diff(err, test.wantErr); diff != "" {
			t.Errorf("error mismatch %s", diff)
		}
		if diff := cmp.Diff(retrievedUser, test.want, cmpCompare); diff != "" {
			t.Errorf("result mismatch %s %s", test.name, diff)
		}
	}
}
