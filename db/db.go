package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBinstance function
func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoURI := os.Getenv("MONGODB_URL")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

// Client database instance
var Client *mongo.Client = DBinstance()

// OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)

	return collection
}