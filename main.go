package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/kier1021/togo/db"
	"github.com/kier1021/togo/server"
)

func main() {
	godotenv.Load(".env")

	// Connect to MongoDB
	mongoDB := db.NewMongoDB()
	if err := mongoDB.Connect(); err != nil {
		log.Fatal(err)
	}

	// Defer the disconnection from MongoDB
	defer func() {
		log.Println("Disconnecting from MongoDB")
		if err := mongoDB.GetClient().Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Run the API server
	srv := server.NewAPIServer(&db.DB{MongoDB: mongoDB})
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatal("Error when starting the server:", err)
		}
	}()

	// Gracefully shutdown the server
	// Ref: https://github.com/gin-gonic/gin#graceful-shutdown-or-restart
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
