package todo

import (
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func ProvideTodoRepository(DB *gorm.DB) TodoRepository {
	return TodoRepository{DB: DB}
}

func (t *TodoRepository) FindAll() []Todo {
	var todos []Todo
	t.DB.Find(&todos)

	return todos
}

func (t *TodoRepository) Save(todo Todo) Todo {
	t.DB.Save(&todo)

	return todo
}
