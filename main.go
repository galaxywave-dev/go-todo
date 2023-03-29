package main

import (
	"galaxywave.com/go-todo/todo"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&todo.Todo{})

	return db
}

func main() {
	db := initDB()

	todoAPI := InitTodoAPI(db)

	r := gin.Default()

	r.GET("/todos", todoAPI.FindAll)
	r.POST("/todos", todoAPI.Create)

	err := r.Run(":8088")
	if err != nil {
		panic(err)
	}
}
