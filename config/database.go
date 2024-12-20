package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func ConnectDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://rajneeshkumar2545:2O2JAZ7kfaHOjkbT@stock-glimpse-db.ypstj.mongodb.net/golang-fiber-hrms?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("MongoDB ping error: ", err)
	}

	Client = client
	Database = client.Database("hrms")
	log.Println("Successfully connected to MongoDB")
}