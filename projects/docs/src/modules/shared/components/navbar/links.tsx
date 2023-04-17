import { component$ } from "@builder.io/qwik";
import AppConfig from "~/config/app.config";

export const Links = component$(() => {
  return (
    <>
      {Object.entries(AppConfig.links).map(
        ([key, link]) =>
          !link.hidden && (
            <a
              href={link.path}
              key={key}
              target={link._blank ? "_blank" : undefined}
            >
              {link.label}
            </a>
          )
      )}
    </>
  );
});
