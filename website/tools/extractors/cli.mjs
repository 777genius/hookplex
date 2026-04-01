import fs from "node:fs/promises";
import path from "node:path";
import { docsToolsRoot, repoBrowserUrl, repoRoot } from "../config/site.mjs";
import { normalizeGeneratedMarkdown, renderMarkdownPage } from "../lib/frontmatter.mjs";
import { ensureDir, listMarkdownFiles } from "../lib/fs.mjs";
import { makeEntity, localeTitle } from "../lib/site-model.mjs";
import { run } from "../lib/process.mjs";

export async function extractCLI() {
  const root = path.join(docsToolsRoot, "cli");
  const markdownDir = path.join(root, "markdown");
  const manifestPath = path.join(root, "manifest.json");
  await ensureDir(markdownDir);
  await run(
    "go",
    [
      "run",
      "./cli/plugin-kit-ai/cmd/plugin-kit-ai",
      "__docs",
      "export-cli",
      "--out-dir",
      markdownDir,
      "--manifest-path",
      manifestPath
    ],
    { cwd: repoRoot }
  );

  const manifest = JSON.parse(await fs.readFile(manifestPath, "utf8"));
  const markdownFiles = await listMarkdownFiles(markdownDir);
  const markdownMap = new Map();
  for (const filePath of markdownFiles) {
    markdownMap.set(path.basename(filePath), await fs.readFile(filePath, "utf8"));
  }

  const entities = [];
  const pages = [];
  const filenameToLink = new Map();
  for (const entry of manifest) {
    filenameToLink.set(entry.file_name, entry.slug);
  }

  for (const entry of manifest) {
    const body = normalizeGeneratedMarkdown(
      rewriteLinks(markdownMap.get(entry.file_name) || "", filenameToLink)
    );
    const canonicalId = `command:${entry.command_path.toLowerCase().replaceAll(" ", ":")}`;
    entities.push(
      makeEntity({
        canonicalId,
        kind: "command",
        surface: "cli",
        localeStrategy: "mirrored",
        title: entry.command_path,
        summary: entry.short || entry.long || "",
        stability: entry.deprecated ? "public-beta" : "public-stable",
        maturity: entry.deprecated ? "deprecated" : "stable",
        sourceKind: "cobra-doc",
        sourceRef: `cli:${entry.command_path}`,
        pathEn: `/en/api/cli/${entry.slug}`,
        pathRu: `/ru/api/cli/${entry.slug}`,
        searchTerms: [entry.command_path, ...(entry.aliases || [])]
      })
    );
    for (const locale of ["en", "ru"]) {
      const intro =
        locale === "ru"
          ? "Сгенерировано из реального Cobra command tree."
          : "Generated from the live Cobra command tree.";
      pages.push({
        locale,
        relativePath: path.join(locale, "api", "cli", `${entry.slug}.md`),
        content: renderMarkdownPage(
          {
            title: localeTitle(locale, entry.command_path, entry.command_path),
            description: entry.short || entry.long || entry.command_path,
            canonicalId,
            surface: "cli",
            section: "api",
            locale,
            generated: true,
            editLink: false,
            stability: entry.deprecated ? "public-beta" : "public-stable",
            maturity: entry.deprecated ? "deprecated" : "stable",
            sourceRef: `cli:${entry.command_path}`,
            translationRequired: false
          },
          `<DocMetaCard surface="cli" stability="${entry.deprecated ? "public-beta" : "public-stable"}" maturity="${entry.deprecated ? "deprecated" : "stable"}" source-ref="cli:${entry.command_path}" source-href="${repoBrowserUrl(`cli:${entry.command_path}`)}" />\n\n# ${entry.command_path}\n\n${intro}\n\n${body}`
        )
      });
    }
  }

  for (const locale of ["en", "ru"]) {
    const heading = locale === "ru" ? "CLI Reference" : "CLI Reference";
    const coreCommands = manifest.filter((entry) => !["bundle", "completion", "skills"].includes(entry.command_path.split(" ")[1]));
    const grouped = {
      core: coreCommands,
      bundle: manifest.filter((entry) => entry.command_path.split(" ")[1] === "bundle"),
      completion: manifest.filter((entry) => entry.command_path.split(" ")[1] === "completion"),
      skills: manifest.filter((entry) => entry.command_path.split(" ")[1] === "skills")
    };
    const renderList = (entries) =>
      entries
        .map((entry) => `- [\`${entry.command_path}\`](/${locale}/api/cli/${entry.slug})`)
        .join("\n");
    const summary =
      locale === "ru"
        ? "CLI — это основная пользовательская часть продукта: создание repo, проверка, тесты, инспекция и установка."
        : "The CLI is the main user-facing surface: scaffold, validate, test, inspect, and install flows.";
    const guidance =
      locale === "ru"
        ? [
            "Используйте `init`, чтобы начать новый plugin repo.",
            "Используйте `render` и `validate --strict` как основной рабочий путь.",
            "Используйте bundle-команды только для переносимых Python и Node bundle-артефактов."
          ]
        : [
            "Use `init` to start a new plugin repo.",
            "Use `render` and `validate --strict` as the primary authored workflow.",
            "Use bundle commands only for portable handoff of Python or Node runtime bundles."
          ];
    const sections = [
      locale === "ru" ? "## Основные команды" : "## Core Commands",
      renderList(grouped.core),
      locale === "ru" ? "## Bundle" : "## Bundle",
      renderList(grouped.bundle),
      locale === "ru" ? "## Completion" : "## Completion",
      renderList(grouped.completion),
      locale === "ru" ? "## Skills" : "## Skills",
      renderList(grouped.skills)
    ]
      .filter(Boolean)
      .join("\n");
    pages.push({
      locale,
      relativePath: path.join(locale, "api", "cli", "index.md"),
      content: renderMarkdownPage(
        {
          title: heading,
          description: "Generated CLI reference",
          canonicalId: "page:api:cli:index",
          surface: "cli",
          section: "api",
          locale,
          generated: true,
          editLink: false,
          stability: "public-stable",
          maturity: "stable",
          sourceRef: "cli/plugin-kit-ai",
          translationRequired: false
        },
        `# ${heading}\n\n${summary}\n\n${guidance.map((line) => `- ${line}`).join("\n")}\n\n${sections}`
      )
    });
  }

  return { entities, pages };
}

function rewriteLinks(body, filenameToLink) {
  return body.replace(/\(([^)]+)\.md\)/g, (_full, target) => {
    const link = filenameToLink.get(`${target}.md`);
    if (!link) {
      return `(${target}.md)`;
    }
    return `(${link})`;
  });
}
