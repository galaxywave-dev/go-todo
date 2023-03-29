package todo

type TodoService struct {
	TodoRepository TodoRepository
}

func ProvideTodoService(t TodoRepository) TodoService {
	return TodoService{TodoRepository: t}
}

func (t *TodoService) FindAll() []Todo {
	return t.TodoRepository.FindAll()
}

func (t *TodoService) Save(todo Todo) Todo {
	t.TodoRepository.Save(todo)

	return todo
}
