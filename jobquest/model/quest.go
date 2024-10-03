package model

import (
	"context"
	"gojobquest/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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

var questCollection = config.QuestCollection(config.DbConnect(), "quests")

func (q *Quest) Create(ctx context.Context) (*mongo.InsertOneResult, error) {
	
	result, err := questCollection.InsertOne(ctx, q)
	if err != nil {
		return nil, err
	}

	return result, nil

}


func (q *Quest) Quests(ctx context.Context)([]Quest,error){
	var quest []Quest
	cur, err := questCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil,err}
		err = cur.All(ctx,&quest)
		if err != nil {
			return nil,err
			}
			return quest,nil

}