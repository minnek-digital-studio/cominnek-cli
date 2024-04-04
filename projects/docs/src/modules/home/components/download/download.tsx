import { component$ } from "@builder.io/qwik";
import Styles from "./download.module.scss";
import { Container } from "~/modules/shared/components/container/container";
import { Title } from "../../../shared/components/title/title";
import { Button } from "~/modules/shared/components/button";

export const Download = component$(() => {
  return (
    <section id="download" class={Styles.root}>
      <Container class={Styles.content}>
        <Title class={Styles.title} bold>
          Download
        </Title>
        <Button.Group variant="inline" class={Styles.buttons}>
          <Button
            href="https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/cominnek-4.0.0.dmg"
            variant="dark"
            class={Styles.button}
            title="MacOS"
          >
            <i class={["fa-brands fa-apple", Styles.icon]}></i>
          </Button>
          <Button
            href="https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/cominnek-4.0.0.exe"
            variant="dark"
            class={Styles.button}
            title="Windows"
          >
            <i class={["fa-brands fa-windows", Styles.icon]}></i>
          </Button>
          <Button
            href="https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/cominnek-4.0.0.deb"
            _target="_blank"
            variant="dark"
            class={Styles.button}
            title="Linux"
          >
            <i class={["fa-brands fa-linux", Styles.icon]}></i>
          </Button>
        </Button.Group>
      </Container>
    </section>
  );
});
