import { component$} from "@builder.io/qwik";
import {
  QwikCityProvider,
  RouterOutlet,
  ServiceWorkerRegister,
} from "@builder.io/qwik-city";
import { RouterHead } from "./modules/shared/components/router-head/router-head";

import "./modules/shared/styles/global.scss";
import { QwikPartytown } from "./modules/shared/components/partytown/partytown";

export default component$(() => {
  return (
    <QwikCityProvider>
      <head>
        <meta charSet="utf-8" />
        <RouterHead />
        <QwikPartytown forward={["dataLayer.push"]} />
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link
          rel="preconnect"
          href="https://fonts.gstatic.com"
          crossOrigin={"use-credentials"}
        />
        <link
          href="https://fonts.googleapis.com/css2?family=Bebas+Neue&family=Roboto:wght@100;300;400;500;700;900&display=swap"
          rel="stylesheet"
        />
        <script
          defer
          src="https://kit.fontawesome.com/7c47b054f1.js"
          crossOrigin="anonymous"
          type="text/partytown"
        ></script>
      </head>
      <body lang="en">
        <RouterOutlet />
        <ServiceWorkerRegister />
      </body>
    </QwikCityProvider>
  );
});
