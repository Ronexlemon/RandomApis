package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type Quest struct{
	ID primitive.ObjectID    `json:"id" bson:"_id"`
	Title string        `json:"title" bson:"title"`
	Level int64          `json:"level" bson:"level"`
	Experience int64            `json:"exp" bson:"experience"`
	ReqItem []string      `json:"req_item" bson:"req_item"`
	ReqSkill []string      `json:"req_skill" bson:"req_skill"`
	
	ReqMoney int64         `json:"req_money" bson:"req_money"`
	ReqTime int64          `json:"req_time" bson:"req_time"`


}