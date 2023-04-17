import { Slot, component$ } from "@builder.io/qwik";
import Styles from "./card.module.scss";

export const Heading = component$(() => {
  return (
    <div class={Styles.heading}>
      <Slot />
    </div>
  );
});
