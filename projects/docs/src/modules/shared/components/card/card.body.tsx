import { Slot, component$ } from "@builder.io/qwik";
import Styles from "./card.module.scss";

export const Body = component$(() => {
  return (
    <div class={Styles.body}>
      <Slot />
    </div>
  );
});
