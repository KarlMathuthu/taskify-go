package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id"`
	TaskId      string             `json:"id" bson:"id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	IsDone      bool               `json:"isDone" bson:"isDone"`
}
