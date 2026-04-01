import path from "node:path";
import { docsToolsRoot, repoBrowserUrl, sourceRefs, websiteRoot } from "../config/site.mjs";
import { normalizeGeneratedMarkdown, renderMarkdownPage } from "../lib/frontmatter.mjs";
import { ensureDir } from "../lib/fs.mjs";
import { makeEntity } from "../lib/site-model.mjs";
import { run } from "../lib/process.mjs";
import fs from "node:fs/promises";

async function ensurePythonVenv(venvDir) {
  const pythonPath = path.join(venvDir, "bin", "python");
  const cliPath = path.join(venvDir, "bin", "pydoc-markdown");
  try {
    await fs.access(cliPath);
    return { pythonPath, cliPath };
  } catch {
    await run("python3", ["-m", "venv", venvDir], { cwd: websiteRoot });
    await run(pythonPath, ["-m", "pip", "install", "--disable-pip-version-check", "pydoc-markdown==4.8.2"], {
      cwd: websiteRoot
    });
    return { pythonPath, cliPath };
  }
}

export async function extractPythonRuntime() {
  const root = path.join(docsToolsRoot, "python-runtime");
  const venvDir = path.join(docsToolsRoot, "python-venv");
  await ensureDir(root);
  const { cliPath } = await ensurePythonVenv(venvDir);
  const body = normalizeGeneratedMarkdown(await run(
    cliPath,
    [
      "--module",
      "plugin_kit_ai_runtime",
      "--search-path",
      "../python/plugin-kit-ai-runtime/src",
      "--render-toc"
    ],
    { cwd: websiteRoot }
  ));

  const canonicalId = "python-runtime:plugin_kit_ai_runtime";
  const entities = [
    makeEntity({
      canonicalId,
      kind: "package",
      surface: "runtime-python",
      localeStrategy: "mirrored",
      title: "plugin_kit_ai_runtime",
      summary: "Public Python runtime helpers",
      stability: "public-stable",
      maturity: "stable",
      sourceKind: "pydoc-markdown",
      sourceRef: sourceRefs.pythonRuntime,
      pathEn: "/en/api/runtime-python/plugin-kit-ai-runtime",
      pathRu: "/ru/api/runtime-python/plugin-kit-ai-runtime",
      searchTerms: ["plugin_kit_ai_runtime", "python runtime"]
    })
  ];

  const pages = ["en", "ru"].map((locale) => ({
    locale,
    relativePath: path.join(locale, "api", "runtime-python", "plugin-kit-ai-runtime.md"),
    content: renderMarkdownPage(
      {
        title: "plugin_kit_ai_runtime",
        description: "Generated Python runtime reference",
        canonicalId,
        surface: "runtime-python",
        section: "api",
        locale,
        generated: true,
        editLink: false,
        stability: "public-stable",
        maturity: "stable",
        sourceRef: sourceRefs.pythonRuntime,
        translationRequired: false
      },
      `<DocMetaCard surface="runtime-python" stability="public-stable" maturity="stable" source-ref="${sourceRefs.pythonRuntime}" source-href="${repoBrowserUrl(sourceRefs.pythonRuntime)}" />\n\n# plugin_kit_ai_runtime\n\n${locale === "ru" ? "Сгенерировано через pydoc-markdown." : "Generated via pydoc-markdown."}\n\n${body}`
    )
  }));

  for (const locale of ["en", "ru"]) {
    pages.push({
      locale,
      relativePath: path.join(locale, "api", "runtime-python", "index.md"),
      content: renderMarkdownPage(
        {
          title: "Python Runtime",
          description: "Generated Python runtime reference",
          canonicalId: "page:api:runtime-python:index",
          surface: "runtime-python",
          section: "api",
          locale,
          generated: true,
          editLink: false,
          stability: "public-stable",
          maturity: "stable",
          sourceRef: sourceRefs.pythonRuntime,
          translationRequired: false
        },
        `# Python Runtime\n\n${
          locale === "ru"
            ? "Открывайте эту зону, когда нужен helper-level API для поддерживаемого repo-local Python runtime lane."
            : "Open this area when you need the helper-level API for the supported repo-local Python runtime lane."
        }\n\n${
          locale === "ru"
            ? "- Это runtime helpers, а не install wrappers.\n- Этот путь лучше всего подходит Python-first командам, которые осознанно принимают repo-local runtime tradeoff.\n- Для широкого выбора формы проекта начните с `/guide/what-you-can-build` и `/concepts/choosing-runtime`."
            : "- These are runtime helpers, not install wrappers.\n- This lane fits Python-first teams that intentionally accept the repo-local runtime tradeoff.\n- For the broader product-shape decision, start with `/guide/what-you-can-build` and `/concepts/choosing-runtime`."
        }\n\n${
          locale === "ru"
            ? "## Когда не нужно начинать отсюда\n\n- Если вы ещё не решили, нужен ли вам Python lane вообще, сначала прочитайте `/guide/what-you-can-build`, `/guide/choose-a-target` и `/concepts/choosing-runtime`."
            : "## When Not To Start Here\n\n- If you are still deciding whether you need the Python lane at all, start with `/guide/what-you-can-build`, `/guide/choose-a-target`, and `/concepts/choosing-runtime`."
        }\n\n- [\`plugin_kit_ai_runtime\`](/${locale}/api/runtime-python/plugin-kit-ai-runtime)`
      )
    });
  }

  return { entities, pages };
}
