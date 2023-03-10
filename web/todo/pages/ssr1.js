import React, { useState } from "react";
import toast, { Toaster } from "react-hot-toast";
import useSWR, {SWRConfig} from "swr";
import { MdDelete } from "react-icons/md";
import Link from "next/link";

import { getTodos } from "./api";

export async function getStaticProps () {
  // `getStaticProps` is executed on the server side.
  const todos = await getTodos()
  return {
    props: {
      fallback: {
        '/api/todos': todos
      }
    }
  }
}
export default function Home({ fallback }) {
  // SWR hooks inside the `SWRConfig` boundary will use those values.
  return (
    <SWRConfig value={{ fallback }}>
      <Todos />
    </SWRConfig>
  )
}
function Todos() {
  {console.log("Rendering...")}
  const url = "http://localhost:8088/todos/";

  const [text, setText] = useState("");
  const { data, error, isLoading, mutate } = useSWR("api/todos", getTodos,{revalidateOnFocus:false});
  if (error) return "An error has occurred.";
  if (isLoading) {
    if (data) {
      console.log("isloading " , data.length)
    }
    return "Loading..." + (data ? data.length : null)
  };
  
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
              await fetch(url, {
                method: "POST",
                body: JSON.stringify(newTodo),
              }).then((res) => res.json());
              mutate();
              toast.success("Successfully added the new item.");
            } catch (e) {
              toast.error("Failed to add the new item." + e);
            }
          }}
        >
          Add
        </button>
        <button
          type="submit"
          style={{ marginLeft: 10 }}
          onClick={() => {
            setText("");
          }}>Clear</button>
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
                        await fetch(`${url}${todo.id}`, { method: "DELETE" });
                        mutate();
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
