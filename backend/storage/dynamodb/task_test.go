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

type taskMockDynamoDBAPI struct {
	table map[string][]map[string]types.AttributeValue
	// userIdDueDateIndex []map[string]types.AttributeValue
}

func (m taskMockDynamoDBAPI) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	var task storage.Task
	err := attributevalue.UnmarshalMap(params.Item, &task)
	if err != nil {
		log.Fatal(err)
	}
	v, ok := m.table[task.UserID]
	if !ok {
		// There's no partition with the given userId yet so we can append the params.Item
		m.table[task.UserID] = append(m.table[task.UserID], params.Item)
	} else {
		// Find if an existing task with same id exists. If found, replace it. Else, append the item to the partition.
		var found bool
		for i := range v {
			var _task storage.Task
			err := attributevalue.UnmarshalMap(v[i], &_task)
			if err != nil {
				log.Fatal(err)
			}
			if task.ID == _task.ID {
				m.table[task.UserID][i] = params.Item
				found = true
				break
			}
		}
		if !found {
			m.table[task.UserID] = append(m.table[task.UserID], params.Item)
		}
	}
	return &dynamodb.PutItemOutput{}, nil
}

func (m taskMockDynamoDBAPI) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return &dynamodb.QueryOutput{}, nil
}

type TaskKey struct {
	UserID string `dynamodbav:"user_id"`
	ID     string `dynamodbav:"id"`
}

func (m taskMockDynamoDBAPI) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	var key TaskKey
	err := attributevalue.UnmarshalMap(params.Key, &key)
	if err != nil {
		log.Fatal(err)
	}
	v, ok := m.table[key.UserID]
	if !ok {
		return &dynamodb.GetItemOutput{}, nil
	}
	for i := range v {
		var _task storage.Task
		err := attributevalue.UnmarshalMap(v[i], &_task)
		if err != nil {
			log.Fatal(err)
		}
		if key.ID == _task.ID {
			item := make(map[string]types.AttributeValue)
			for key, value := range v[i] {
				item[key] = value
			}
			return &dynamodb.GetItemOutput{Item: item}, nil
		}
	}
	return &dynamodb.GetItemOutput{}, nil
}

func TestCreateTask(t *testing.T) {
	table := map[string][]map[string]types.AttributeValue{}
	client := taskMockDynamoDBAPI{table}
	store := ddb.NewStorage(client)

	task, err := store.CreateTask(storage.Task{UserID: "userId1", Title: "title1", Done: false, DueDate: time.Now().Unix()})
	if diff := cmp.Diff(err, nil); diff != "" {
		t.Errorf("error mismatch %s", diff)
	}
	v, ok := table[task.UserID]
	if diff := cmp.Diff(ok, true); diff != "" {
		t.Error("partitionKey not found in table")
	}
	var found bool
	var _task storage.Task
	for i := range v {
		err := attributevalue.UnmarshalMap(v[i], &_task)
		if err != nil {
			log.Fatal(err)
		}
		if _task.ID == task.ID {
			found = true
			break
		}
	}
	if !found {
		t.Error("task not found in table")
	}
	if diff := cmp.Diff(_task, task); diff != "" {
		t.Errorf("result mismatch %s", diff)
	}
}

func TestGetTask(t *testing.T) {
	timestamp := time.Now()
	userId := "userId1"
	id := "id1"
	task := storage.Task{
		UserID:  userId,
		ID:      id,
		Title:   "title1",
		Done:    false,
		DueDate: timestamp.Unix(),
		Created: timestamp,
		Updated: timestamp,
	}
	item, _ := attributevalue.MarshalMap(task)
	table := make(map[string][]map[string]types.AttributeValue)
	table[userId] = []map[string]types.AttributeValue{item}
	client := taskMockDynamoDBAPI{table}
	store := ddb.NewStorage(client)

	tests := []struct {
		name string
		req  struct {
			userId string
			id     string
		}
		want    *storage.Task
		wantErr error
	}{
		{
			name: "Found",
			req: struct {
				userId string
				id     string
			}{userId, id},
			want:    &task,
			wantErr: nil,
		},
		{
			name: "Not Found Partition Key",
			req: struct {
				userId string
				id     string
			}{"notFound", "id"},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "Not Found Sort Key",
			req: struct {
				userId string
				id     string
			}{userId, "notFound"},
			want:    nil,
			wantErr: nil,
		},
	}

	cmpCompare := cmp.Comparer(func(actual, want *storage.Task) bool {
		if want == nil {
			if actual != nil {
				return false
			}
			return true
		}
		return cmp.Equal(*actual, *want)
	})

	for _, test := range tests {
		retrievedTask, err := store.GetTask(test.req.userId, test.req.id)
		if diff := cmp.Diff(err, test.wantErr); diff != "" {
			t.Errorf("error mismatch %s", diff)
		}
		if diff := cmp.Diff(retrievedTask, test.want, cmpCompare); diff != "" {
			t.Errorf("result mismatch %s %s", test.name, diff)
		}
	}
}
