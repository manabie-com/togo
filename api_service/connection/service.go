package connection

import (
	"api_service/proto"
	"os"

	"google.golang.org/grpc"
)

//ServiceConnection ...
type ServiceConnection struct {
	ClientSessionService proto.SessionServiceClient
	ClientAccountService proto.AccountServiceClient
	ClientTodoService    proto.TodoServiceClient
}

//DialToSessionServiceServer ...
func DialToSessionServiceServer() *ServiceConnection {

	port := os.Getenv("SESSION_SERVICE_PORT")

	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &ServiceConnection{
		ClientSessionService: proto.NewSessionServiceClient(conn),
	}
}

//DialToAccountServiceServer ...
func DialToAccountServiceServer() *ServiceConnection {

	port := os.Getenv("ACCOUNT_SERVICE_PORT")

	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &ServiceConnection{
		ClientAccountService: proto.NewAccountServiceClient(conn),
	}
}

//DialToTodoServiceServer ...
func DialToTodoServiceServer() *ServiceConnection {

	port := os.Getenv("TODO_SERVICE_PORT")

	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &ServiceConnection{
		ClientTodoService: proto.NewTodoServiceClient(conn),
	}
}
