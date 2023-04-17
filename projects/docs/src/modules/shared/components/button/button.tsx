import type { Component } from "@builder.io/qwik";
import { Slot, component$ } from "@builder.io/qwik";
import Styles from "./button.module.scss";
import type { IButtonGroupProps, IButtonProps } from "./button.d";
import { Link } from "@builder.io/qwik-city";

interface IButtonComponent extends Component<IButtonProps> {
  Group: typeof ButtonGroup;
}

const Button = component$<IButtonProps>((props) => {
  const variant = props.variant || "primary";
  const _class = [Styles.root, Styles[variant], props.class];
  const type = (props.href === undefined && props.type) || "button";

  if (props.href) {
    return (
      <Link
        href={props.href}
        class={_class}
        target={props._target}
        title={props.title}
      >
        <Slot />
      </Link>
    );
  }

  return (
    <button class={_class} type={type} title={props.title}>
      <Slot />
    </button>
  );
}) as IButtonComponent;

const ButtonGroup = component$<IButtonGroupProps>((props) => {
  const variant = props.variant || "spaced";
  const style = {
    gap: variant === "spaced" ? `${(props as any).gap}px` : undefined,
    flexWrap: props.wrap ? "wrap" : undefined,
  };

  return (
    <div class={[Styles.group, Styles[variant]]} style={style}>
      <Slot />
    </div>
  );
});

Button.Group = ButtonGroup;

export { Button };
