package models

import (
	"context"
	"errors"
	"fmt"
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
func (u *User) CreateUser(ctx context.Context) (*mongo.InsertOneResult, error) {
    // Check if the email already exists
    var existingUser User
    err := collection.FindOne(ctx, bson.M{"email": u.Email}).Decode(&existingUser)
    
    if err == nil {
        // If no error, it means a user with this email already exists
        return nil, fmt.Errorf("user already exists")
    } else if err != mongo.ErrNoDocuments {
        // If it's not the "no documents" error, return the actual error
        return nil, err
    }

    // If no existing user found, proceed with the insertion
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    result, err := collection.InsertOne(ctx, u)
    if err != nil {
        return nil, err
    }
    return result, nil
}


func (u *User) Login(ctx context.Context,email string,password string)(*User,error){
	var user User
	
    err := collection.FindOne(ctx,bson.M{"email":email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
            return nil, errors.New("user not found")
        }
		return nil,err}
		return &user,nil

}
//getUserProfile
func (u *User) Profile(ctx  context.Context,user_id primitive.ObjectID)(*User,error){
	var user User
	
	err:= collection.FindOne(ctx,bson.M{"_id":user_id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user Profile found")
			}
			return nil,err
			}
			return &user,nil
}