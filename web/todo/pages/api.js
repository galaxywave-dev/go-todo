// const fetcher = (url) => fetch(url).then((res) => res.json());
const baseUrl = 'http://localhost:8088';

export async function getTodos() {
  let todos = await fetch(baseUrl+'/todos/').then((res) => res.json());
  return todos.data;
}
export async function addTodo2(url ,todo) {
  const newTodo= await fetch(`${baseUrl}${url}`, {
    method: "POST",
    body: JSON.stringify(todo),
  }).then((res) => res.json());
  console.log("Added data %o", newTodo.data)
  return newTodo.data;
}

export async function addTodo(url ,todo) {
  const newTodo= await fetch(`${baseUrl}${url}`, {
    method: "POST",
    body: JSON.stringify(todo.arg),
  }).then((res) => res.json());
  console.log("Added data %o", newTodo.data)
  return newTodo.data;
}

export async function deleteTodo(url, todo){
  await fetch(`${baseUrl}${url}${todo.arg}`, { method: 'DELETE' });
}