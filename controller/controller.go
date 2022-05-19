package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/octoi/go-todo-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// MONGO HELPERS

func AddTodo(ginContext *gin.Context, todo string) {
	inserted, err := collection.InsertOne(context.Background(), model.Todo{
		Todo:       todo,
		IsResolved: false,
	})

	if err != nil {
		ginContext.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginContext.JSON(200, model.Todo{
		ID:         inserted.InsertedID.(primitive.ObjectID),
		Todo:       todo,
		IsResolved: false,
	})
}

func UpdateTodo(ginContext *gin.Context, todo model.Todo) {
	id, _ := primitive.ObjectIDFromHex(todo.ID.Hex())
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"todo":       todo.Todo,
		"isResolved": todo.IsResolved,
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		ginContext.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ginContext.JSON(200, todo)
}

func DeleteTodo(ginContext *gin.Context, todoId string) {
	id, _ := primitive.ObjectIDFromHex(todoId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		ginContext.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Println("delete count", deleteCount)
}

func GetAllTodo(ginContext *gin.Context) {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		ginContext.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	var todos []primitive.M

	for cursor.Next(context.Background()) {
		var todo bson.M
		err := cursor.Decode(&todo)

		if err != nil {
			ginContext.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		todos = append(todos, todo)
	}

	defer cursor.Close(context.Background())

	ginContext.JSON(200, todos)
}
