package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"session_service/auth"
	"session_service/proto"
	"session_service/service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {

	var err error
	// Load environment
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	auth.InitRedis()
}

func main() {

	// Open port
	listener, err := net.Listen("tcp", ":"+os.Getenv("SESSION_SERVICE_PORT"))
	if err != nil {
		panic(err)
	}
	log.Print("Session Service running on port :" + os.Getenv("SESSION_SERVICE_PORT"))

	// Register session service
	srv := grpc.NewServer()
	service := service.NewSessionServiceServer()
	proto.RegisterSessionServiceServer(srv, service)

	// Gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx := context.Background()
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("Shutting down gRPC Session service server...")

			srv.GracefulStop()

			<-ctx.Done()
		}
	}()

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
