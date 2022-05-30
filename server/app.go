package main

import (
	"context"
	"fmt"
	"os"

	lr "togo/utils/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set logging
	logger := lr.NewLogger(os.Getenv("LOG_LEVEL"))

	// Create a new client and connect to the server
	mongodb_uri := fmt.Sprintf("%v:%v", os.Getenv("MONGODB_URI"), os.Getenv("DATABASE_PORT"))

	// Set client options
	clientOptions := options.Client().ApplyURI(mongodb_uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	logger.Info().Msgf("Connecting to database %v", mongodb_uri)

	if err != nil {
		logger.Fatal().Err(err).Msg("Connection failed")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		logger.Fatal().Err(err).Msg("Ping failed")
	}

	logger.Info().Msg("Connected to database")
}
