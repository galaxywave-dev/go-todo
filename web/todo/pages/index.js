import React, { useState } from "react";
import useSWR from "swr";
import Link from "next/link";

import { getTodos, addTodo, deleteTodo } from "./api";

export default function Home() {
  const [text, setText] = useState("");
  //const fetcher = (url) => fetch(url).then((res) => res.json());
  const { data, error, mutate } = useSWR("api/todos", getTodos);
  if (error) return "An error has occurred.";

  return (
    <div>
      <h1>TODO APP</h1>
      <ul>
        <li>
          <Link href="/mutate1">Mutate with No Parameters</Link>
        </li>
        <li>
          <Link href="/mutate2">Mutate Custom cache</Link>
        </li>
        <li>
          <Link href="/useSwrMutate1">
            useSwrMutate with populate cache (ADD)
          </Link>
        </li>
        <li>
          <Link href="/useSwrMutate2">
            useSwrMutate with optimisticData (ADD)
          </Link>
        </li>
        <li>
          <Link href="/useSwrMutateDel1">
            useSwrMutate with populate cache (DELETE)
          </Link>
        </li>
        <li>
          <Link href="/useSwrMutateDel2">
            useSwrMutate with optimisticData (DELETE)
          </Link>
        </li>
        <li>
          <Link href="/ssr1">
            useSwrConfig with SSR
          </Link>
        </li>
      </ul>
    </div>
  );
}
