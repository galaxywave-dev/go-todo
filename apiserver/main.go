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

	models.InitDBConnection() // new

	r.Run(":8088")
}
