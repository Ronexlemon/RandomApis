package models

import (
	"context"
	"errors"
	"gomanage/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	UserID    primitive.ObjectID `json:"user_id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

var collection *mongo.Collection = config.Collection(config.NewClient(),"users")
func (u *User) CreateUser(ctx context.Context)(*mongo.InsertOneResult,error){
	result,err := collection.InsertOne(ctx,u)
	if err !=nil{
		return nil,err
	}
	return result,nil


}

func (u *User) Login(ctx context.Context,email string)(*User,error){
	var user User
    err := collection.FindOne(ctx,bson.M{"email":email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
            return nil, errors.New("user not found")
        }
		return nil,err}
		return &user,nil

}