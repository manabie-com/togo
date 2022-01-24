package connection

import (
	"account_service/proto"
	"os"

	"google.golang.org/grpc"
)

//ServiceConnection ...
type ServiceConnection struct {
	ClientSessionService proto.SessionServiceClient
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
