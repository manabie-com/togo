package connection

import (
	"os"
	"todo_service/proto"

	"google.golang.org/grpc"
)

//ServiceConnection ...
type ServiceConnection struct {
	ClientSessionService proto.SessionServiceClient
	ClientAccountService proto.AccountServiceClient
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
