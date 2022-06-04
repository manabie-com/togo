package dynamodb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jssoriao/todo-go/storage"
)

func (s *Storage) CreateTask(task storage.Task) (storage.Task, error) {
	saveTaskFunc := func(id string) error {
		// Check if duplicate entry exists using the composite key (user_id, id)
		existingTask, err := s.GetTask(task.UserID, id)
		if err != nil {
			return err
		}
		if existingTask != nil {
			return ErrIDExists
		}

		// Assign the autogenerated id to the task
		task.ID = id

		// Add timestamps
		timestamp := time.Now()
		task.Created = timestamp
		task.Updated = timestamp

		item, err := attributevalue.MarshalMap(task)
		if err != nil {
			return fmt.Errorf("failed to marshal the task object to dynamodb format: %w", err)
		}
		if _, err := s.client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: &tasksTableName,
			Item:      item,
		}); err != nil {
			return err
		}
		return nil
	}

	_, err := generateID(saveTaskFunc, 12, 5, Alphanumeric)
	if err != nil {
		return storage.Task{}, err
	}

	return task, nil
}

func (s *Storage) GetTask(userId, id string) (*storage.Task, error) {
	key, err := attributevalue.MarshalMap(map[string]string{
		"user_id": userId,
		"id":      id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the task key to dynamodb format: %w", err)
	}
	resp, err := s.client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: &tasksTableName,
		Key:       key,
	})
	if err != nil {
		return nil, err
	}
	if len(resp.Item) == 0 {
		return nil, nil
	}

	task := storage.Task{}
	if err = attributevalue.UnmarshalMap(resp.Item, &task); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dynamodb item to task struct: %w", err)
	}
	return &task, err
}

func (s *Storage) CountTasksForTheDay(userID string, dueDate time.Time) (int, error) {
	indexName := "userIdDueDateIndex"
	keyCondExpression := "user_id=:userID AND due_date BETWEEN :dt1 AND :dt2"
	year, month, day := dueDate.Date()
	dt1 := time.Date(year, month, day, 0, 0, 0, 0, dueDate.Location()).Unix()
	dt2 := time.Date(year, month, day, 23, 59, 59, 0, dueDate.Location()).Unix()
	exprAttrValues, err := attributevalue.MarshalMap(map[string]interface{}{
		":userID": userID,
		":dt1":    dt1,
		":dt2":    dt2,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to marshal the ExpressionAttributeValues: %w", err)
	}
	resp, err := s.client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:                 &tasksTableName,
		IndexName:                 &indexName,
		KeyConditionExpression:    &keyCondExpression,
		ExpressionAttributeValues: exprAttrValues,
	})
	if err != nil {
		return 0, err
	}
	return len(resp.Items), nil
}