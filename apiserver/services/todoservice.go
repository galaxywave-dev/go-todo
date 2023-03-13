package services

import (
	"errors"

	"galaxywave.com/go-todo/apiserver/models"
)

var TODOChan chan *models.Todo

func Init() {
	TODOChan = make(chan *models.Todo)
}
func CreateTodo(todo *models.Todo) error {
	// Create todo
	count := models.DB.Where(todo).Find(&models.Todo{}).RowsAffected
	if count > 0 {
		return errors.New("record already exists")
	}
	models.DB.Create(todo)
	TODOChan <- todo
	return nil
}
