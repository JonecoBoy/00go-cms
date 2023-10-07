package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func connectToMongoDB() (*mongo.Client, error) {
	// Set the MongoDB URI
	mongoURI := "mongodb://localhost:27017" // Replace with your MongoDB URI

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

func mongoDB() {
	// Connect to MongoDB
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Select a database and collection
	database := client.Database("your-database-name")
	collection := database.Collection("your-collection-name")

	// Insert a document
	insertResult, err := collection.InsertOne(context.TODO(), bson.M{"name": "John"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)
}
