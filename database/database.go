package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCollection *mongo.Collection

// Init Mongo
func MongoDBInit(dbUrl string) {
	clientOptions := options.Client().ApplyURI(dbUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	MongoCollection = client.Database("taskify").Collection("tasks")
}

// Get Tasks
// Add Task
// Remove Task
// Update Task
