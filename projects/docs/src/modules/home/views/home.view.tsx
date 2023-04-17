import { component$ } from "@builder.io/qwik";
import { Hero } from "../components/hero/hero";
import { Download } from "../components/download";

export const HomeView = component$(() => {
  return (
    <>
      <h1 class="sr-only">Cominnek</h1>
      <Hero />
      <Download />
    </>
  );
});
