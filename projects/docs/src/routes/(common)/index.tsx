import type { DocumentHead } from "@builder.io/qwik-city";
import { HomeView } from "~/modules/home/views/home.view";

export default HomeView;

export const head: DocumentHead = {
  title: "Cominnek",
  meta: [
    {
      name: "description",
      content:
        "Create commits & pull requests easily. A git & github CLI manager with super powers",
    },
    {
      name: "keywords",
      content:
        "git, github, cli, manager, pull request, commit, super powers, cominnek, Minnek Digital Studio",
    },
    {
      name: "author",
      content: "Minnek Digital Agency",
    },
    {
      name: "robots",
      content: "index, follow",
    },
    {
      name: "og:title",
      content: "Cominnek",
    },
    {
      name: "og:description",
      content:
        "Create commits & pull requests easily. A git & github CLI manager with super powers",
    },
    {
      name: "og:image",
      content: "https://www.cominnek.com/img/og-image.png",
    },
  ],
};
