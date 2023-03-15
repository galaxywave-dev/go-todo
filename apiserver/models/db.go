package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDBConnection() (*gorm.DB, error) {
	db, err := newDBConnection()
	DB = db
	return db, err
}
func newDBConnection() (*gorm.DB, error) {
	var database *gorm.DB
	var err error
	database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{}, &Todo{}, &User{})
	if err != nil {
		panic("Failed to migrate to database!")
	}

	return database, nil
}
