package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

var (
	todos  = []Todo{}
	nextID = 1
	mu     sync.Mutex
)

func main() {
	r := gin.Default()

	// Enable CORS (for frontend requests)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	// Routes
	r.GET("/todos", getTodos)
	r.POST("/todos", addTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.Run(":9999") // Run server on port 8080
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	mu.Lock()
	newTodo.ID = nextID
	nextID++
	todos = append(todos, newTodo)
	mu.Unlock()

	c.JSON(http.StatusOK, newTodo)
}

func deleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	mu.Lock()
	defer mu.Unlock()

	for i, todo := range todos {
		if id == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
