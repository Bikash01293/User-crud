package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Db connection
func ConnectDb() *mongo.Collection{
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	//cancel the db connection if db not connected within 10 seconds
	if cancel != nil {
		ctx.Done()
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")) 
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("dB successfully connected...")
	collection := client.Database("mydb").Collection("mycollection")
	return collection
}