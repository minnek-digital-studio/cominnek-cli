import type { ITitleProps } from "./title.d";
import { Slot, component$ } from "@builder.io/qwik";
import Style from "./title.module.scss";

export const Title = component$<ITitleProps>((props) => {
  const Type = props.type || "h2";

  return (
    <Type
      class={[
        props.class,
        Style.title,
        {
          [Style.bold]: props.bold,
        },
      ]}
      cmk-font={props.font}
    >
      <Slot />
    </Type>
  );
});
