import fs from "node:fs/promises";
import path from "node:path";
import { docsToolsRoot, repoBrowserUrl, sourceRefs, websiteRoot } from "../config/site.mjs";
import { normalizeGeneratedMarkdown, renderMarkdownPage, stripLeadingHeading, stripLeadingTypedocPrelude } from "../lib/frontmatter.mjs";
import { ensureDir, listMarkdownFiles } from "../lib/fs.mjs";
import { makeEntity } from "../lib/site-model.mjs";
import { run } from "../lib/process.mjs";

export async function extractNodeRuntime() {
  const root = path.join(docsToolsRoot, "node-runtime");
  await ensureDir(root);
  await run(
    "pnpm",
    [
      "exec",
      "typedoc",
      "--plugin",
      "typedoc-plugin-markdown",
      "--entryPoints",
      "../npm/plugin-kit-ai-runtime/index.d.ts",
      "--readme",
      "none",
      "--out",
      root
    ],
    { cwd: websiteRoot }
  );

  const markdownFiles = await listMarkdownFiles(root);
  const entities = [];
  const pages = [];

  for (const filePath of markdownFiles) {
    const stem = path.basename(filePath, ".md");
    const body = stripLeadingHeading(stripLeadingTypedocPrelude(normalizeGeneratedMarkdown(await fs.readFile(filePath, "utf8"))));
    const slug = stem === "README" ? "runtime" : stem.toLowerCase();
    const displayTitle = humanizeNodeTitle(stem);
    const canonicalId = `node-runtime:${stem}`;
    entities.push(
      makeEntity({
        canonicalId,
        kind: "package",
        surface: "runtime-node",
        localeStrategy: "mirrored",
        title: displayTitle,
        summary: `Node runtime reference: ${stem}`,
        stability: "public-stable",
        maturity: "stable",
        sourceKind: "typedoc-markdown",
        sourceRef: sourceRefs.nodeRuntime,
        pathEn: `/en/api/runtime-node/${slug}`,
        pathRu: `/ru/api/runtime-node/${slug}`,
        searchTerms: [stem, "plugin-kit-ai-runtime", "node runtime"]
      })
    );
    for (const locale of ["en", "ru"]) {
      const intro =
        locale === "ru"
          ? "Сгенерировано через TypeDoc и typedoc-plugin-markdown."
          : "Generated via TypeDoc and typedoc-plugin-markdown.";
      pages.push({
        locale,
        relativePath: path.join(locale, "api", "runtime-node", `${slug}.md`),
        content: renderMarkdownPage(
          {
            title: displayTitle,
            description: `Generated Node runtime reference for ${stem}`,
            canonicalId,
            surface: "runtime-node",
            section: "api",
            locale,
            generated: true,
            editLink: false,
            stability: "public-stable",
            maturity: "stable",
            sourceRef: sourceRefs.nodeRuntime,
            translationRequired: false
          },
          `<DocMetaCard surface="runtime-node" stability="public-stable" maturity="stable" source-ref="${sourceRefs.nodeRuntime}" source-href="${repoBrowserUrl(sourceRefs.nodeRuntime)}" />\n\n# ${displayTitle}\n\n${intro}\n\n${body}`
        )
      });
    }
  }

  for (const locale of ["en", "ru"]) {
    const list = entities
      .map((entry) => `- [\`${entry.title}\`](/${locale}/api/runtime-node/${entry.pathEn.split("/").pop()})`)
      .join("\n");
    pages.push({
      locale,
      relativePath: path.join(locale, "api", "runtime-node", "index.md"),
      content: renderMarkdownPage(
        {
          title: "Node Runtime",
          description: "Generated Node runtime reference",
          canonicalId: "page:api:runtime-node:index",
          surface: "runtime-node",
          section: "api",
          locale,
          generated: true,
          editLink: false,
          stability: "public-stable",
          maturity: "stable",
          sourceRef: sourceRefs.nodeRuntime,
          translationRequired: false
        },
        `# Node Runtime\n\n${
          locale === "ru"
            ? "Открывайте эту зону, когда вам нужен helper-level API для поддерживаемого repo-local Node/TypeScript runtime lane."
            : "Open this area when you need the helper-level API for the supported repo-local Node/TypeScript runtime lane."
        }\n\n${
          locale === "ru"
            ? "- Это runtime helpers, а не install wrappers.\n- Этот путь подходит для mainstream non-Go stable lane.\n- Для общего выбора между lanes начните с `/guide/what-you-can-build` и `/concepts/choosing-runtime`."
            : "- These are runtime helpers, not install wrappers.\n- This is the mainstream non-Go stable lane.\n- For the broader lane choice, start with `/guide/what-you-can-build` and `/concepts/choosing-runtime`."
        }\n\n${locale === "ru" ? "Сгенерированные Node runtime страницы:" : "Generated Node runtime pages:"}\n\n${list}`
      )
    });
  }

  return { entities, pages };
}

function humanizeNodeTitle(stem) {
  if (stem === "README") {
    return "Overview";
  }

  return stem.replace(/([a-z])([A-Z])/g, "$1 $2");
}
