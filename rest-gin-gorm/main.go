package main

import (
	"rest-gin-gorm/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})

	return db
}

func main() {
	db := initDB()

	productAPI := InitProductAPI(db)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	err := r.Run(":8088")
	if err != nil {
		panic(err)
	}
}
