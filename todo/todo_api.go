package todo

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoAPI struct {
	TodoService TodoService
}

func ProvideTodoAPI(t TodoService) TodoAPI {
	return TodoAPI{TodoService: t}
}

func (t *TodoAPI) FindAll(c *gin.Context) {
	todos := t.TodoService.FindAll()
	c.JSON(http.StatusOK, gin.H{"todos": ToTodoDTOs(todos)})
}

func (t *TodoAPI) Create(c *gin.Context) {
	var todoDTO TodoDTO
	err := c.BindJSON(&todoDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdTodo := t.TodoService.Save(todoDTO.ToTodo())

	c.JSON(http.StatusOK, gin.H{"todo": ToTodoDTO(createdTodo)})
}
