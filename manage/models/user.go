package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	 UserID primitive.ObjectID `json:"user_id" bson:"_id"`
	 Username string `json:"username" bson:"username"`
	 Email string `json:"email" bson:"email"`
	 Password string `json:"password" bson:"password"`
	 CreatedAt time.Time `json:"created_at" bson:"created_at"`
	 UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`

}

