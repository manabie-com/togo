package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jssoriao/todo-go/services/users"
	ddb "github.com/jssoriao/todo-go/storage/dynamodb"
)

type RequestBody struct {
	DailyLimit int `json:"daily_limit"`
}

type ResponseBody struct {
	ID         string `json:"id"`
	DailyLimit int    `json:"daily_limit"`
}

var ddbClient *dynamodb.Client

func Handler(request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	var (
		headers = map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		}
	)

	// TODO: Add validation

	var payload RequestBody
	if err := json.Unmarshal([]byte(request.Body), &payload); err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers:    headers,
			Body:       "\"message\": \"Invalid request body\"",
		}, nil
	}

	store := ddb.NewStorage(ddbClient)
	usersSvc, err := users.NewHandler(store)
	if err != nil {
		return nil, err
	}

	user, err := usersSvc.CreateUser(context.Background(), &users.User{DailyLimit: payload.DailyLimit})
	if err != nil {
		return nil, err
	}

	userBytes, err := json.Marshal(ResponseBody{ID: user.ID, DailyLimit: user.DailyLimit})
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headers,
		Body:       string(userBytes),
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
