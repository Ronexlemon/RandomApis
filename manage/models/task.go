package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStatus string
const (Pending TaskStatus = "Pending"
Completed TaskStatus = "Completed"
Failed TaskStatus = "In-progress"

)
type Task struct{
	Task_id  primitive.ObjectID `json:"task_id" bson:"task_id"`
	Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      TaskStatus    `json:"status"` // pending, in-progress, completed
    DueDate     time.Time `json:"due_date"`
    Priority    string    `json:"priority"` // low, medium, high
    User  primitive.ObjectID  `json:"user_id"` // User ID
    
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`

}