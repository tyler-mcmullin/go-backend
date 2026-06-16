package mongo

import (

	"context"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "time"
	"os"
	"log"

)

var MongoClient *mongo.Client

func CreateMongoConnection() error {

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Error: Failed to set MongoDB URI")
	}

	_, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	MongoClient = client

	return nil
	
}