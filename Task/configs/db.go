package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client{
	client,err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURL()))
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

//Client Instance
var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client,collectionName string) *mongo.Collection{
	collection := client.Database("TaskManage").Collection(collectionName)
	return collection

}