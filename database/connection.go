package database

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

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error while loading .env")
	}

	mongodbURI := os.Getenv("DB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURI))

	if err != nil {
		log.Fatal("Database connection error")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.Connect(ctx)

	fmt.Print("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(client.Database("book-manager").Collection(collectionName))
	return collection
}
