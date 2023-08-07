import { useContent, useLocation } from "@builder.io/qwik-city";
import { component$, useStyles$ } from "@builder.io/qwik";
import styles from "./on-this-page.scss?inline";
import AppConfig from "~/config/app.config";

const makeEditPageUrl = (url: string) => {
  const path = url.replace(/\/$/, "");
  return `${AppConfig.repoUrl}/tree/master/projects/docs/src/routes${path}/index.mdx`;
};

export default component$(() => {
  useStyles$(styles);

  const { headings } = useContent();
  const contentHeadings =
    headings?.filter((h) => h.level === 2 || h.level === 3) || [];

  const { url } = useLocation();
  const editUrl = makeEditPageUrl(url.pathname);

  return (
    <aside class="on-this-page">
      {contentHeadings.length > 0 ? (
        <>
          <h6>On This Page</h6>
          <ul>
            {contentHeadings.map((h) => (
              <li key={h.id}>
                <a
                  href={`#${h.id}`}
                  class={{
                    block: true,
                    indent: h.level > 2,
                  }}
                >
                  {h.text}
                </a>
              </li>
            ))}
          </ul>
        </>
      ) : null}

      <h6>More</h6>
      <ul>
        <li>
          <a href={editUrl} target="_blank">
            Edit this page
          </a>
        </li>
        <li>
          <a href={AppConfig.repoUrl} target="_blank">
            GitHub
          </a>
        </li>
      </ul>
    </aside>
  );
});
