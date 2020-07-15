package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Method from: https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

// DB is a global variable to hold db connection
var Client *mongo.Client

func ConnectDB() error {

	// Get the environment variable
	host := os.Getenv("HOST_ENV")
	username := os.Getenv("USER_ENV")
	password := os.Getenv("PASS_ENV")

	// May be of use: https://stackoverflow.com/questions/17588876/mongodb-conf-bind-ip-127-0-0-1-does-not-work-but-0-0-0-0-works
	var (
		client *mongo.Client
		//mongoURL = "mongodb://mongo:27017"		// Needs to be used in docker container
		mongoURL = fmt.Sprintf("mongodb://%s:27017", host) // Needs to be used for local dev
	)

	fmt.Println("Initalising connection...")
	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL).SetAuth(options.Credential{Username: username, Password: password}))

	fmt.Println("Connecting to MongoDB Server...")
	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("could not connect to mongo db service: %v\n", err)
		return err
	}

	fmt.Println("Pinging MongoDB")
	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
		return err
	}

	fmt.Println("connected to nosql database:", mongoURL)

	Client = client

	return nil
}
