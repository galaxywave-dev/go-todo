import React, { useState } from "react";
import toast, { Toaster } from "react-hot-toast";
import useSWR from "swr";
import { MdDelete } from "react-icons/md";
import Link from "next/link";

import { getTodos, addTodo2 } from "./api";

export default function Home() {
  const url = "http://localhost:8088/todos/";

  const [text, setText] = useState("");
  const { data, error, mutate } = useSWR("api/todos", getTodos, {revalidateOnFocus:false});
  if (error) return "An error has occurred.";

  return (
    <div>
      <Toaster toastOptions={{ position: "bottom-center" }} />
      <h1>TODO APP</h1>
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
              title: text,
            };
            try {
              await mutate(addTodo2("/todos/", newTodo), {
                optimisticData : data => [...data, newTodo],
                populateCache: (addedData, data) => ([...data, addedData]),
                revalidate: false
              });

              //mutate(data => [...data, newTodo], {revalidate: false});
              // var result = await fetch(url, {
              //   method: "POST",
              //   body: JSON.stringify(newTodo),
              // }).then((res) => res.json());
              //mutate(data => [...(data.filter(x => x.id != newTodo.id)), result.data], {revalidate: false});
              toast.success("Successfully added the new item.");
            } catch (e) {
              //mutate(data => [...(data.filter(x => x.id != newTodo.id))], {revalidate: false})
              toast.error("Failed to add the new item." + e);

            }
          }}
        >
          Add
        </button>
        <Link href="/" type="submit" style={{ marginLeft: 10 }}>
          Home
        </Link>
      </form>
      <ul>
        {data
          ? data.map((todo, index) => {
              return (
                <li key={index}>
                  {todo.title}
                  <button
                    style={{ marginLeft: 10, marginTop: 10 }}
                    type="submit"
                    onClick={async () => {
                      try {
                        var result = await fetch(`${url}${todo.id}`, { method: 'DELETE' });
                        console.log("result", result)
                        if(result.status == 200) {
                          mutate(data => data.filter(x => x.id != todo.id), {revalidate: false});
                          toast.success("Successfully remove the item.");
                        }
                        else{
                          toast.error("Failed to remove the item.");
                        }
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
