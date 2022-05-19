package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "todo"
const collectionName = "todo"

var collection *mongo.Collection

// connect with mongoDB
func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[+] MONGO DB CONNECTED TO ", connectionString)

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("[i] COLLECTION INSTANCE IS READY")
}
