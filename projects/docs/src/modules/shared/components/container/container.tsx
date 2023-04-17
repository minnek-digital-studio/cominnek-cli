import { Slot, component$ } from "@builder.io/qwik";
import Styles from "./container.module.scss";
import type { IContainerProps } from "./container.d";

export const Container = component$<IContainerProps>((props) => {
  return (
    <div class={[Styles.container, props.class]}>
      <Slot />
    </div>
  );
});
