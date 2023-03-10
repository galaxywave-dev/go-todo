package main

import (
	"galaxywave.com/go-todo/controllers"
	"galaxywave.com/go-todo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	books := r.Group("/books")
	{
		books.GET("/", controllers.FindBooks)
		books.POST("/", controllers.CreateBook)
		books.GET("/:id", controllers.FindBook)
		books.PATCH("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook) // new
	}

	todos := r.Group("/todos")
	{
		todos.GET("/", controllers.FindTodos)
		todos.POST("/", controllers.CreateTodo)
		todos.GET("/:id", controllers.FindTodo)
		todos.PATCH("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo) // new
		// Define OPTIONS route to handle preflight request
		todos.OPTIONS("/:id", func(c *gin.Context) {
			// Set CORS headers
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type")

			// Respond with status 204 (no content)
			c.Status(204)
		})

	}

	models.InitDBConnection() // new

	r.Run(":8088")
}
