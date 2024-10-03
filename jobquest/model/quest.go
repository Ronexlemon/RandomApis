package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "Approved"
	StatusPending  Status = "pending"
	StatusRejected Status = "rejected"
)

type Quest struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	Level      int64              `json:"level" bson:"level"`
	Experience int64              `json:"exp" bson:"experience"`
	ReqItem    []string           `json:"req_item" bson:"req_item"`
	ReqSkill   []string           `json:"req_skill" bson:"req_skill"`
	Status     Status             `json:"status" bson:"status"`
	ReqMoney   int64              `json:"req_money" bson:"req_money"`
	ReqTime    int64              `json:"req_time" bson:"req_time"`
}
