import { Head } from "$fresh/runtime.ts";

export default function Home() {
  return (
    <>
      <Head>
        <title>Fresh Hello</title>
      </Head>
      <div class="p-4 mx-auto max-w-screen-md">
        <img
          src="/logo.svg"
          class="w-32 h-32"
          alt="the fresh logo: a sliced lemon dripping with juice"
        />
        <h1 class="text-2xl font-bold my-4">Hello World</h1>
        <p class="my-4">
          This is a simple Fresh app.
        </p>
        <input
          type="text"
          class="mt-4 p-2 border rounded w-full"
          placeholder="Enter text here... (this doesn't do anything yet)"
          aria-label="Input box"
        />
      </div>
    </>
  );
}
