package models

import (
	"github.com/samuelmatt1998/golang-todo-app/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

var TaskCollection *mongo.Collection = config.DB.Collection("tasks")

type Task struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Task       string             `bson:"task" json:"task"`
	TaskStatus string             `bson:"task_status" json:"task_status"` // "done" or "not done"
}