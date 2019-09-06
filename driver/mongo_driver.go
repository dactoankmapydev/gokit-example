package driver

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB() *MongoDB {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected ok!")
	Mongo.Client = client
	return Mongo
}
