package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServerClient struct{
	MongoDB *mongo.Client
}

func NewClient()*mongo.Client{
	//create context for only 10 seconds
	ctx,cancel := context.WithTimeout(context.Background(),time.Second *10)
	defer  cancel()
	client,err := mongo.NewClient(options.Client().ApplyURI(EnvLoad()))
	if err !=nil{
		log.Fatal("Failed to create a client")
	}
	//ping client
	err = client.Ping(ctx,nil)
	if err !=nil{
		log.Fatal("Failed to ping the client")
		}

		return client
}

func Collection(client *mongo.Client,collectionName string)*mongo.Collection{
	collection := client.Database("Databse").Collection(collectionName)
	return collection
}