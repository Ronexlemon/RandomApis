package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServerClient struct {
	MongoDB *mongo.Client
}

func NewClient() *mongo.Client {
	//create context for only 10 seconds
	
	client,err := mongo.NewClient(options.Client().ApplyURI(EnvLoad()))
	if err !=nil{
		log.Fatal(err)
	}
	ctx,_:=context.WithTimeout(context.Background(),time.Second *10) // give it 10 second to connect
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)}
		//ping the database
		err = client.Ping(ctx,nil)
		if err != nil {
			log.Fatal(err)}
			fmt.Println("Connected to db")

			return client
}

func Collection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Manage").Collection(collectionName)
	return collection
}
