package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/octoi/go-todo-api/controller"
	"github.com/octoi/go-todo-api/model"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartServer(port int32) {
	router.GET("/todos", getAllTodoRoute)
	router.GET("/todos/:id", getOneTodoRoute)
	router.POST("/todos", createTodoRoute)
	router.PATCH("/todos/:id", updateTodoRoute)
	router.DELETE("/todos/:id", deleteTodoRoute)

	// listen on port
	addr := fmt.Sprintf(":%v", port)
	router.Run(addr)
}

// API routes

func getAllTodoRoute(c *gin.Context) {
	controller.GetAllTodo(c)
}

func getOneTodoRoute(c *gin.Context) {
	id := c.Param("id")
	controller.GetOneTodo(c, id)
}

func createTodoRoute(c *gin.Context) {
	todo := getTodoFromBody(c)
	controller.AddTodo(c, todo.Todo)
}

func updateTodoRoute(c *gin.Context) {
	todo := getTodoFromBody(c)
	controller.UpdateTodo(c, todo)
}

func deleteTodoRoute(c *gin.Context) {
	id := c.Param("id")
	controller.DeleteTodo(c, id)
}

// helpers
func getTodoFromBody(c *gin.Context) model.Todo {
	body := c.Request.Body
	valueRaw, err := ioutil.ReadAll(body)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	var todo model.Todo
	json.Unmarshal(valueRaw, &todo)
	return todo
}
