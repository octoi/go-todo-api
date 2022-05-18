package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PORT = 5000
const MONGO_DB_URL = "mongodb://localhost:27017"

func main() {
	mongoClient := connectWithMongo()
	todoCollection := mongoClient.Database("todo").Collection("todo")

	router := gin.Default()

	router.GET("/todo", func(c *gin.Context) {
		getAllTodo(c, todoCollection)
	})

	portNumber := 5000
	port := fmt.Sprintf(":%v", portNumber)

	router.Run(port)
}

func connectWithMongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_DB_URL))

	if err != nil {
		panic(err)
	}

	return client
}

func getAllTodo(ginContext *gin.Context, todoCollection *mongo.Collection) {
	cursor, err := todoCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		ginContext.JSON(500, gin.H{
			"message": "Error getting data from database",
		})
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		ginContext.JSON(500, gin.H{
			"message": "Failed to decode data from database",
		})
	}

	ginContext.JSON(200, gin.H{
		"todos": results,
	})
}
