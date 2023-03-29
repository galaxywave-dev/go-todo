// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"galaxywave.com/go-todo/todo"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitTodoAPI(db *gorm.DB) todo.TodoAPI {
	todoRepository := todo.ProvideTodoRepository(db)
	todoService := todo.ProvideTodoService(todoRepository)
	todoAPI := todo.ProvideTodoAPI(todoService)
	return todoAPI
}
