package models

import (
	"context"
	"gomanage/config"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskStatus string

const (
	Pending   TaskStatus = "Pending"
	Completed TaskStatus = "Completed"
	Failed    TaskStatus = "In-progress"
)

type Task struct {
	Task_id     primitive.ObjectID `json:"task_id" bson:"task_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      TaskStatus         `json:"status"` // pending, in-progress, completed
	DueDate     time.Time          `json:"due_date"`
	Priority    string             `json:"priority"` // low, medium, high
	User        primitive.ObjectID `json:"user_id"`  // User ID

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var TaskCollection *mongo.Collection = config.Collection(config.NewClient(), "task")

// create a task
func (t *Task) Create(ctx context.Context) (*mongo.InsertOneResult, error) {
	result, err := TaskCollection.InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// get userTasks
func (t *Task) UserTasks(ctx context.Context, user_id primitive.ObjectID) ([]Task,error) {
	var tasks []Task
	result, err := TaskCollection.Find(ctx, bson.M{"user_id": user_id}, options.Find().SetSort(bson.D{{Key: "created_at", Value: 1}}))
	if err != nil {
		return nil,err
	}
	err = result.All(ctx, &tasks)
	if err != nil {
		return nil,err
	}
	return tasks,nil
}

// update a task
func (t *Task) Update(ctx context.Context, user_id primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{"task_id": t.Task_id, "user_id": user_id}
	update := bson.M{"$set": bson.M{"status": t.Status, "due_date": t.DueDate}}
	result, err := TaskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// delete a task
func (t *Task) Delete(ctx context.Context, user_id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"task_id": t.Task_id, "user_id": user_id}
	result, err := TaskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil

}
