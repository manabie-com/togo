package database

import "go.mongodb.org/mongo-driver/mongo"

type Repository interface {
	GetCollection() *mongo.Collection
}
