//go:build wireinject
// +build wireinject

package models

import (
	"github.com/google/wire"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDBConnection() (*gorm.DB, error) {
	wire.Build(
		newDBConnection,
	)
	return nil, nil
}
func newDBConnection() (*gorm.DB, error) {
	var database *gorm.DB
	var err error
	database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to migrate to database!")
	}

	return database, nil
}
