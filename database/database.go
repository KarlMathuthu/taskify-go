package database

import (
	"context"
	"log"

	"github.com/KarlMathuthu/taskify-go/models"
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
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoCollection = client.Database("taskify").Collection("tasks")
}

// Get Tasks
// Add Task
func AddTaskToDB(task models.Task) (insertResult interface{}) {
	result, err := MongoCollection.InsertOne(context.TODO(), task)

	if err != nil {
		log.Fatal(err)
	}
	return result.InsertedID
}

// Remove Task
// Update Task
