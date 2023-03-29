//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"galaxywave.com/go-todo/todo"
)

func InitTodoAPI(db *gorm.DB) todo.TodoAPI {
	wire.Build(todo.ProvideTodoRepository, todo.ProvideTodoService, todo.ProvideTodoAPI)

	return todo.TodoAPI{}
}
