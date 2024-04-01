package models

type Task struct {
	ID          string `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	IsDone      bool   `json:"isDone" bson:"isDone"`
}
