package services

import (
	"errors"

	"galaxywave.com/go-todo/apiserver/models"
	"github.com/tjgq/broadcast"
)

var TODOChan chan *models.Todo
var Cast *broadcast.Broadcaster

func Init() {
	TODOChan = make(chan *models.Todo)
	Cast = broadcast.New(10)
}
func CreateTodo(todo *models.Todo) error {
	// Create todo
	count := models.DB.Where(todo).Find(&models.Todo{}).RowsAffected
	if count > 0 {
		return errors.New("record already exists")
	}
	models.DB.Create(todo)
	//TODOChan <- todo
	Cast.Send(todo)
	return nil
}
