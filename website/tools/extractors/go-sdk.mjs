import fs from "node:fs/promises";
import path from "node:path";
import { docsToolsRoot, publicGoPackages, repoBrowserUrl, repoRoot } from "../config/site.mjs";
import { normalizeGeneratedMarkdown, renderMarkdownPage, stripLeadingHeading } from "../lib/frontmatter.mjs";
import { ensureDir } from "../lib/fs.mjs";
import { makeEntity, localeTitle } from "../lib/site-model.mjs";
import { run } from "../lib/process.mjs";

export async function extractGoSDK() {
  const root = path.join(docsToolsRoot, "go-sdk");
  await ensureDir(root);
  const entities = [];
  const pages = [];

  for (const pkg of publicGoPackages) {
    const outPath = path.join(root, `${pkg.id}.md`);
    await run(
      "go",
      [
        "run",
        "github.com/princjef/gomarkdoc/cmd/gomarkdoc@v1.1.0",
        "--repository.url",
        "https://github.com/777genius/plugin-kit-ai",
        "--repository.default-branch",
        "main",
        "--output",
        outPath,
        `./${pkg.relativePath}`
      ],
      { cwd: repoRoot }
    );
    const body = stripLeadingHeading(normalizeGeneratedMarkdown(await fs.readFile(outPath, "utf8")));
    const canonicalId = `go-package:${pkg.importPath}`;
    const slug = pkg.id === "root" ? "sdk" : pkg.id;
    const packageLabel = pkg.id === "root" ? "sdk" : pkg.id;
    entities.push(
      makeEntity({
        canonicalId,
        kind: "package",
        surface: "go-sdk",
        localeStrategy: "mirrored",
        title: packageLabel,
        summary: `Public Go package ${pkg.importPath}`,
        stability: pkg.id === "platformmeta" ? "public-beta" : "public-stable",
        maturity: pkg.id === "platformmeta" ? "beta" : "stable",
        sourceKind: "gomarkdoc",
        sourceRef: pkg.relativePath,
        pathEn: `/en/api/go-sdk/${slug}`,
        pathRu: `/ru/api/go-sdk/${slug}`,
        relatedIds: pkg.id === "claude" ? ["event-platform:claude"] : pkg.id === "codex" ? ["event-platform:codex"] : [],
        searchTerms: [pkg.importPath, slug]
      })
    );
    for (const locale of ["en", "ru"]) {
      const intro =
        locale === "ru"
          ? "Сгенерировано из публичного Go package через gomarkdoc."
          : "Generated from the public Go package via gomarkdoc.";
      pages.push({
        locale,
        relativePath: path.join(locale, "api", "go-sdk", `${slug}.md`),
        content: renderMarkdownPage(
          {
            title: localeTitle(locale, packageLabel, packageLabel),
            description: `Generated Go SDK package reference for ${pkg.importPath}`,
            canonicalId,
            surface: "go-sdk",
            section: "api",
            locale,
            generated: true,
            editLink: false,
            stability: pkg.id === "platformmeta" ? "public-beta" : "public-stable",
            maturity: pkg.id === "platformmeta" ? "beta" : "stable",
            sourceRef: pkg.relativePath,
            translationRequired: false
          },
          `<DocMetaCard surface="go-sdk" stability="${pkg.id === "platformmeta" ? "public-beta" : "public-stable"}" maturity="${pkg.id === "platformmeta" ? "beta" : "stable"}" source-ref="${pkg.relativePath}" source-href="${repoBrowserUrl(pkg.relativePath)}" />\n\n# ${packageLabel}\n\n${intro}\n\n**Import path:** \`${pkg.importPath}\`\n\n${body}`
        )
      });
    }
  }

  for (const locale of ["en", "ru"]) {
    const packageRows = publicGoPackages
      .map((pkg) => {
        const packageLabel = pkg.id === "root" ? "sdk" : pkg.id;
        const slug = pkg.id === "root" ? "sdk" : pkg.id;
        const summary =
          pkg.id === "root"
            ? locale === "ru"
              ? "Корневой composition/runtime entry package."
              : "Root composition and runtime entry package."
            : pkg.id === "claude"
              ? locale === "ru"
                ? "Публичные Claude-oriented handlers и event wiring."
                : "Public Claude-oriented handlers and event wiring."
              : pkg.id === "codex"
                ? locale === "ru"
                  ? "Публичные Codex-oriented handlers и runtime integration."
                  : "Public Codex-oriented handlers and runtime integration."
                : locale === "ru"
                  ? "Platform metadata и support-oriented helpers."
                  : "Platform metadata and support-oriented helpers.";
        return `| [\`${packageLabel}\`](/${locale}/api/go-sdk/${slug}) | ${summary} |`;
      })
      .join("\n");
    const intro =
      locale === "ru"
        ? "Go SDK — рекомендуемый путь по умолчанию, когда нужен самый сильный и предсказуемый контракт для продакшена."
        : "The Go SDK is the recommended default path when you want the strongest production contract.";
    const guidance =
      locale === "ru"
        ? "- Открывайте эту зону, когда строите production-oriented plugin на Go.\n- Это лучший старт, если вы хотите минимальную зависимость от внешних runtime на машинах пользователей.\n- Если вы ещё выбираете между Go, Python и Node, начните с `/guide/what-you-can-build` и `/concepts/choosing-runtime`."
        : "- Open this area when you are building a production-oriented Go plugin.\n- This is the best starting point when you want the least downstream runtime friction.\n- If you are still choosing between Go, Python, and Node, start with `/guide/what-you-can-build` and `/concepts/choosing-runtime`.";
    pages.push({
      locale,
      relativePath: path.join(locale, "api", "go-sdk", "index.md"),
      content: renderMarkdownPage(
        {
          title: "Go SDK",
          description: "Generated Go SDK package reference",
          canonicalId: "page:api:go-sdk:index",
          surface: "go-sdk",
          section: "api",
          locale,
          generated: true,
          editLink: false,
          stability: "public-stable",
          maturity: "stable",
          sourceRef: "sdk",
          translationRequired: false
        },
        `# Go SDK\n\n${intro}\n\n${guidance}\n\n| Package | Summary |\n| --- | --- |\n${packageRows}`
      )
    });
  }

  return { entities, pages };
}
