package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnect() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(LoadEnv()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//ping

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client

}
