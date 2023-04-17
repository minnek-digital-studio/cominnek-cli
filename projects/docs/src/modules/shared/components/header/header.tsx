import { $, component$, useSignal } from "@builder.io/qwik";
import { Title } from "../title/title";
import Styles from "./header.module.scss";
import { Link } from "@builder.io/qwik-city";
import { Container } from "../container/container";
import { Navbar } from "../navbar/navbar";
import AppConfig from "~/config/app.config";
import type { IHeaderProps } from "./header.d";

export const Header = component$<IHeaderProps>((props) => {
  const header = useSignal<Element>();
  const minScroll = props.minScroll || 300;

  const handleScroll = $(() => {
    const headerElement = header.value;
    if (headerElement) {
      const scrollPosition = window.scrollY;
      if (scrollPosition > minScroll) {
        headerElement.classList.add(Styles.scrolled);
      } else {
        headerElement.classList.remove(Styles.scrolled);
      }
    }
  });

  return (
    <header
      class={[
        Styles.main,
        {
          [Styles.scrolled]: props.disableScrollHandler,
        },
      ]}
      ref={header}
      document:onScroll$={
        !props.disableScrollHandler ? handleScroll : undefined
      }
    >
      <Container class={Styles.content}>
        <div>
          <Title class={Styles.title}>
            <Link href="/">{AppConfig.appName}</Link>
          </Title>
        </div>
        <Navbar />
      </Container>
    </header>
  );
});
