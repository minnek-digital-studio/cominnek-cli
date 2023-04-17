import { component$, useSignal, useVisibleTask$ } from "@builder.io/qwik";
import Styles from "./navbar.module.scss";
import { Links } from "./links";

export const Navbar = component$(() => {
  const open = useSignal(false);
  const hamburgerClass = [
    Styles.hamburger,
    "cmk-u-only--mobile",
    { [Styles.open]: open.value },
  ];

  if (open.value) {
    document.addEventListener("keydown", (e) => {
      if (e.key === "Escape") {
        open.value = false;
      }
    });
  }

  useVisibleTask$(({ track }) => {
    const isOpen = track(open);

    if (isOpen.value) {
      document.body.classList.add(Styles.open);
    } else {
      document.body.classList.remove(Styles.open);
    }
  });

  return (
    <div>
      <button
        class={hamburgerClass}
        onClick$={() => (open.value = !open.value)}
        title={open.value ? "Close menu" : "Open menu"}
        aria-label={open.value ? "Close menu" : "Open menu"}
      >
        <span></span>
        <span></span>
        <span></span>
        <span></span>
      </button>
      <div
        class={[
          Styles.menu,
          "cmk-u-only--mobile",
          { [Styles.open]: open.value },
        ]}
      >
        <Links />
      </div>
      <div class={[Styles.links, "cmk-u-only--desktop"]}>
        <Links />
      </div>
    </div>
  );
});
