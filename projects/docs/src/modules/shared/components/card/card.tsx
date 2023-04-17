import { Slot, component$ } from "@builder.io/qwik";
import Styles from "./card.module.scss";
import type { ICardsMainProps } from "./card.d";

export const Main = component$<ICardsMainProps>((props) => {
  const maxWidth = props.maxWidth || "500px";

  return (
    <div
      class={[
        Styles.root,
        props.class,
        {
          [Styles.shadow]: props.shadow,
        },
      ]}
      style={{ maxWidth: maxWidth !== "0" ? maxWidth : undefined }}
    >
      <Slot />
    </div>
  );
});
