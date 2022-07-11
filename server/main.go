package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "pt.example/grcp-test/grcptest"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedTaskServer
}

func (s *server) SaveTodoTask(ctx context.Context, in *pb.TaskReq) (rs *pb.TaskRes, e error) {
	rs = &pb.TaskRes{Message: fmt.Sprintf("Added %s to do successfully", in.GetTitle())}
	return
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
