import React, { useState } from "react";
import toast, { Toaster } from "react-hot-toast";
import useSWR from "swr";
import { MdDelete } from "react-icons/md";

import { getTodos, addTodo, deleteTodo } from "./api";

export default function Home() {
  const [text, setText] = useState("");
  const { data, mutate } = useSWR("/api/todos", getTodos);

  return (
    <div>
      <Toaster toastOptions={{ position: "bottom-center" }} />
      <h1>Next.js TODO APP</h1>
      <form onSubmit={(ev) => ev.preventDefault()}>
        <input
          value={text}
          onChange={(e) => setText(e.target.value)}
          autoFocus
        />
        <button
          type="submit"
          style={{ marginLeft: 10 }}
          onClick={async () => {
            setText("");
            const newTodo = {
              id: Date.now(),
              text,
            };
            try {
              await mutate(addTodo(newTodo), {
                optimisticData: [...data, newTodo],
                rollbackOnError: true,
                populateCache: true,
                revalidate: false,
              });
              toast.success("Successfully added the new item.");
            } catch (e) {
              toast.error("Failed to add the new item.");
            }
          }}
        >
          Add
        </button>
      </form>
      <ul>
        {data
          ? data.map((todo,index) => {
              return (
                <li key={index}>
                  {todo.text}
                  <button
                    style={{ marginLeft: 10, marginTop: 10 }}
                    type="submit"
                    onClick={async () => {
                      try {
                        await mutate(deleteTodo(todo), {
                          optimisticData: [...data, todo],
                          rollbackOnError: true,
                          populateCache: true,
                          revalidate: false,
                        });
                        toast.success("Successfully remove the item.");
                      } catch (e) {
                        toast.error("Failed to remove the item.");
                      }
                    }}
                  >
                    <MdDelete size={10} color="red" />
                  </button>
                </li>
              );
            })
          : null}
      </ul>
    </div>
  );
}
