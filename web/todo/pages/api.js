let todos = [];
const delay = () => new Promise((res) => setTimeout(() => res(), 800));

export async function getTodos() {
  await delay();
  return todos;
}

export async function addTodo(todo) {
  await delay();
  //   if (Math.random() < 0.5) throw new Error("Failed to add new item!");
  todos = [...todos, todo];
  console.log("TODOS %o", todos);
  return todos;
}

export async function deleteTodo(todo){
    await delay;
    todos = todos.filter( x => x.id !== todo.id);
    return todos;
}
