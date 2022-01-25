package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"todo_service/database"
	"todo_service/proto"
	"todo_service/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {

	// Load environment
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	// Open port
	listener, err := net.Listen("tcp", ":"+os.Getenv("TODO_SERVICE_PORT"))
	if err != nil {
		panic(err)
	}
	log.Print("Todo Service running on port :" + os.Getenv("TODO_SERVICE_PORT"))

	// Connect to DB
	var newdb database.DBInfo
	db, err := newdb.GetDB()
	if err != nil {
		fmt.Printf("failed to open database: %v", err)
		return
	}

	// Register service
	srv := grpc.NewServer()
	service := service.NewTodoServiceServer(db)
	proto.RegisterTodoServiceServer(srv, service)

	// Gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx := context.Background()
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("Shutting down gRPC Todo service server...")

			srv.GracefulStop()

			<-ctx.Done()
		}
	}()

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
