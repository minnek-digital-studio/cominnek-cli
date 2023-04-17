import { component$ } from "@builder.io/qwik";
import Styles from "./hero.module.scss";
import { Title } from "~/modules/shared/components/title/title";
import { Text } from "~/modules/shared/components/text/text";
import { useTextAnimation } from "../../hooks/useTextAnimation";
import type { ITypeAnimationProps } from "./hero.d";

export const Hero = component$(() => {
  return (
    <section class={Styles.main} id="hero">
      <Title type="h3" font="heading" class={Styles.title}>
        Save time on each{" "}
        <TypeAnimation
          words={[
            "commit",
            "pull request",
            "branch",
            "merge",
            "push",
            "release",
          ]}
          class={Styles.type}
        />
      </Title>
      <Text class={Styles.subtitle}>
        A simple and easy way to work with your team
      </Text>
    </section>
  );
});

const TypeAnimation = component$<ITypeAnimationProps>((props) => {
  const text = useTextAnimation(props.words);

  return (
    <span class={[props.class]}>
      {text}
      <i class={Styles.type_line} aria-hidden="true"></i>
    </span>
  );
});
