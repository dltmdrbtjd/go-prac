package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	err = client.Ping(ctx, readpref.Primary()) // Primary DB에 대한 연결 체크

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Sparata").Collection("webclass")
	println(collection)
	println("MongoDB PRAC")
}
