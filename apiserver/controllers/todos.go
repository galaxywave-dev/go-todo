package controllers

import (
	"fmt"
	"net/http"

	"galaxywave.com/go-todo/apiserver/models"
	"galaxywave.com/go-todo/apiserver/services"
	"github.com/gin-gonic/gin"
)

// GET /todos
// Get all todos
type CreateTodoInput struct {
	Title string `json:"title" binding:"required"`
}

type TodoInput struct {
	ID uint `uri:"id"`
}

type UpdateTodoInput struct {
	Title string `json:"title"`
}

type FindTodoInput struct {
	Title string `json:"title"`
}

func FindTodos(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	var todos []models.Todo
	models.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Validate input
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create todo
	todo := models.Todo{Title: input.Title}
	if err := services.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func FindTodo(c *gin.Context) { // Get model if exist
	c.Header("Access-Control-Allow-Origin", "*")

	var todo models.Todo

	// if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }

	var input TodoInput
	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input.ID)
	if err := models.DB.Where("id = ?", input.ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&todo).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Get model if exist
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
}
