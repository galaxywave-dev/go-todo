package todo

func (d TodoDTO) ToTodo() Todo {
	return Todo{Title: d.Title}
}

func ToTodoDTO(todo Todo) TodoDTO {
	return TodoDTO{ID: todo.ID, Title: todo.Title}
}

func ToTodoDTOs(todos []Todo) []TodoDTO {
	tododtos := make([]TodoDTO, len(todos))

	for i, itm := range todos {
		tododtos[i] = ToTodoDTO(itm)
	}

	return tododtos
}
