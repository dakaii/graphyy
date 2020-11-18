package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
// GetDatabase returns a database instance.
func GetDatabase(userCollection string) (context.Context, *mongo.Database) {
	url, exists := os.LookupEnv("MONGODB_URL")
	if !exists {
		url = "mongodb://localhost:27017"
	}
	dbName, exists := os.LookupEnv("MONGODB_DB_NAME")
	if !exists {
		dbName = "testingDB"
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(dbName)
	collection := database.Collection(userCollection)
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	models := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "username", Value: bsonx.String("text")}},
		},
	}
	_, err = collection.Indexes().CreateMany(ctx, models, opts)
	return ctx, database
}
