package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jssoriao/todo-go/services/tasks"
	ddb "github.com/jssoriao/todo-go/storage/dynamodb"
)

type RequestBody struct {
	Title   string    `json:"title"`
	DueDate time.Time `json:"due_date"`
}

type ResponseBody struct {
	ID      string    `json:"id"`
	UserID  string    `json:"user_id"`
	Title   string    `json:"title"`
	Done    bool      `json:"done"`
	DueDate time.Time `json:"due_date"`
}

var ddbClient *dynamodb.Client

func Handler(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	var (
		headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}
	)

	var payload RequestBody
	if err := json.Unmarshal([]byte(request.Body), &payload); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    headers,
			Body:       "\"message\": \"Invalid request body\"",
		}, nil
	}

	// Payload title must be minimum of 1 character and maximum of 200 characters
	lenTitle := len(payload.Title)
	if lenTitle < 1 || lenTitle > 200 {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    headers,
			Body:       "\"message\": \"Title character length out of range [1, 200]\"",
		}, nil
	}
	// Due Date must be today or a future date
	y, m, d := time.Now().Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())
	if payload.DueDate.Before(today) {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    headers,
			Body:       "\"message\": \"Due date cannot be set to a past date\"",
		}, nil
	}

	var userId string
	var ok bool
	if userId, ok = request.PathParameters["userId"]; !ok {
		return nil, errors.New("missing pathParameters userId")
	}

	store := ddb.NewStorage(ddbClient)
	tasksSvc, err := tasks.NewHandler(store)
	if err != nil {
		return nil, err
	}

	task, err := tasksSvc.CreateTask(context.Background(), &tasks.Task{UserID: userId, Title: payload.Title, DueDate: payload.DueDate})
	if err != nil {
		if errors.Is(err, tasks.ErrUserNotFound) {
			return &events.APIGatewayProxyResponse{
				StatusCode: 404,
				Headers:    headers,
				Body:       "\"message\": \"User not found\"",
			}, nil
		} else if errors.Is(err, tasks.ErrExceededTasksLimit) {
			return &events.APIGatewayProxyResponse{
				StatusCode: 400,
				Headers:    headers,
				Body:       "\"message\": \"Exceeded maximum tasks for user\"",
			}, nil
		}
		return nil, err
	}

	taskBytes, err := json.Marshal(ResponseBody{
		ID:      task.ID,
		UserID:  task.UserID,
		Title:   task.Title,
		Done:    task.Done,
		DueDate: task.DueDate,
	})
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headers,
		Body:       string(taskBytes),
	}, nil
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("unable to load AWS config, %v", err)
	}
	ddbClient = dynamodb.NewFromConfig(cfg)
}

func main() {
	lambda.Start(Handler)
}
