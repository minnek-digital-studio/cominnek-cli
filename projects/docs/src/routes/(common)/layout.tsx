import { component$, Slot } from '@builder.io/qwik';
import { Header } from "~/modules/shared/components/header/header";

export default component$(() => {
  return (
    <>
      <Header />
      <main>
        <Slot />
      </main>
    </>
  );
});
