package database

import (
	"context"
	"log"

	"github.com/KarlMathuthu/taskify-go/models"
	"go.mongodb.org/mongo-driver/bson"
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
		return
	}

	MongoCollection = client.Database("taskify").Collection("tasks")
}

// Get Tasks
func GetAllTasksDB() (tasks []models.Task) {
	// Leave the bson.M{} empty since we don't want to filter the docs
	cur, err := MongoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var task models.Task
		if err := cur.Decode(&task); err != nil {
			log.Fatal(err)
			return
		}
		tasks = append(tasks, task)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return
	}
	return tasks
}

// Add Task
func AddTaskToDB(task models.Task) (insertResult interface{}) {
	result, err := MongoCollection.InsertOne(context.TODO(), task)

	if err != nil {
		log.Fatal(err)
		return
	}
	return result.InsertedID
}

// Find Task
func FindTaskDB(id string) (*models.Task, error) {
	myTask := MongoCollection.FindOne(context.TODO(), bson.M{"id": id})
	var taskModel models.Task

	err := myTask.Decode(&taskModel)

	if err != nil {
		return nil, err
	}
	if err := myTask.Err(); err != nil {
		return nil, err
	}
	return &taskModel, err
}

// Delete Task
func DeleteTaskDB(id string) (*mongo.DeleteResult, error) {
	deleteResult, err := MongoCollection.DeleteOne(context.TODO(), bson.M{"id": id})
	return deleteResult, err
}

// Update Task
func UpdateTaskDB(id string, updatedTask models.Task) (*mongo.UpdateResult, error) {
	update := bson.M{"$set": bson.M{
		"title":       updatedTask.Title,
		"description": updatedTask.Description,
	}}
	updateResult, err := MongoCollection.UpdateOne(context.TODO(), bson.M{"id": id}, update)
	return updateResult, err
}

// Toggle the isDone field
func UpdateTaskDone(id string, updatedTask models.Task) (*mongo.UpdateResult, error) {
	update := bson.M{"$set": bson.M{"isDone": updatedTask.IsDone}}
	updateResult, err := MongoCollection.UpdateOne(context.TODO(), bson.M{"id": id}, update)
	return updateResult, err
}
