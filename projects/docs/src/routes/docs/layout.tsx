import { component$, Slot, useStyles$ } from "@builder.io/qwik";
import { ContentNav } from "~/modules/docs/component/content-nav";
import OnThisPage from "~/modules/docs/component/on-this-page/on-this-page";
import { createBreadcrumbs, SideBar } from "~/modules/docs/component/sidebar";
import { Container } from "~/modules/shared/components/container/container";
import { Header } from "~/modules/shared/components/header/header";
import { Spacer } from "~/modules/shared/components/spacer";
import Styles from "./docs.scss?inline";
import { useContent, useLocation } from "@builder.io/qwik-city";

export default component$(() => {
  useStyles$(Styles);
  const { url } = useLocation();
  const { menu } = useContent();

  const breadcrumbs = createBreadcrumbs(menu, url.pathname);

  return (
    <div class="cmk-docs">
      <Header disableScrollHandler />
      <Spacer size={40} />
      <main>
        <Container>
          {breadcrumbs.length > 0 ? (
            <ol class="cmk-breadcrumbs">
              {breadcrumbs.map((b, i) => (
                <li key={i}>{b.text}</li>
              ))}
            </ol>
          ) : null}
          <SideBar />
          <div class="cmk-content">
            <article class="cmk-container">
              <Slot />
              <ContentNav />
            </article>
            <OnThisPage />
          </div>
        </Container>
      </main>
    </div>
  );
});
