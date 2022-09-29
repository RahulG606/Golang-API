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
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongo()))

	if err!=nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err!=nil{
		log.Fatal(err)
	}

	// ping the database

	err = client.Ping(ctx,nil)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Connected to database")

	return client
}

var DB *mongo.Client = ConnectDB()

// getting DB collection

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection{

	collection := client.Database("Appdata").Collection(collectionName)
	return collection
}




