import fs from "node:fs/promises";
import path from "node:path";
import { docsToolsRoot, repoBrowserUrl, repoRoot, sourceRefs } from "../config/site.mjs";
import { renderMarkdownPage } from "../lib/frontmatter.mjs";
import { ensureDir } from "../lib/fs.mjs";
import { makeEntity } from "../lib/site-model.mjs";
import { run } from "../lib/process.mjs";

export async function extractPlatformData() {
  const root = path.join(docsToolsRoot, "platform");
  const eventsPath = path.join(root, "events.json");
  const targetsPath = path.join(root, "targets.json");
  const capabilitiesPath = path.join(root, "capabilities.json");
  await ensureDir(root);
  await run(
    "go",
    [
      "run",
      "./cli/plugin-kit-ai/cmd/plugin-kit-ai",
      "__docs",
      "export-support",
      "--events-path",
      eventsPath,
      "--targets-path",
      targetsPath,
      "--capabilities-path",
      capabilitiesPath
    ],
    { cwd: repoRoot }
  );
  const events = JSON.parse(await fs.readFile(eventsPath, "utf8"));
  const targets = JSON.parse(await fs.readFile(targetsPath, "utf8"));
  const capabilities = JSON.parse(await fs.readFile(capabilitiesPath, "utf8"));

  const entities = [];
  const pages = [];

  const platforms = new Map();
  for (const event of events) {
    if (!platforms.has(event.platform)) {
      platforms.set(event.platform, []);
    }
    platforms.get(event.platform).push(event);
  }

  for (const [platform, platformEvents] of platforms) {
    const canonicalId = `event-platform:${platform}`;
    entities.push(
      makeEntity({
        canonicalId,
        kind: "event",
        surface: "platform-events",
        localeStrategy: "mirrored",
        title: platform,
        summary: `${platform} event surface`,
        stability: platformEvents.every((entry) => entry.maturity === "stable") ? "public-stable" : "public-beta",
        maturity: platformEvents.every((entry) => entry.maturity === "stable") ? "stable" : "beta",
        sourceKind: "support-export",
        sourceRef: sourceRefs.supportMatrix,
        pathEn: `/en/api/platform-events/${platform}`,
        pathRu: `/ru/api/platform-events/${platform}`,
        searchTerms: [platform, ...platformEvents.map((entry) => entry.event)]
      })
    );
    const table = platformEvents
      .map(
        (entry) =>
          `| ${entry.event} | ${entry.maturity} | ${entry.contract_class} | ${entry.summary} |\n`
      )
      .join("");
    for (const locale of ["en", "ru"]) {
      pages.push({
        locale,
        relativePath: path.join(locale, "api", "platform-events", `${platform}.md`),
        content: renderMarkdownPage(
          {
            title: platform,
            description: `Event reference for ${platform}`,
            canonicalId,
            surface: "platform-events",
            section: "api",
            locale,
            generated: true,
            editLink: false,
            stability: platformEvents.every((entry) => entry.maturity === "stable") ? "public-stable" : "public-beta",
            maturity: platformEvents.every((entry) => entry.maturity === "stable") ? "stable" : "beta",
            sourceRef: sourceRefs.supportMatrix,
            translationRequired: false
          },
          `<DocMetaCard surface="platform-events" stability="${platformEvents.every((entry) => entry.maturity === "stable") ? "public-stable" : "public-beta"}" maturity="${platformEvents.every((entry) => entry.maturity === "stable") ? "stable" : "beta"}" source-ref="${sourceRefs.supportMatrix}" source-href="${repoBrowserUrl(sourceRefs.supportMatrix)}" />\n\n# ${platform}\n\n| Event | Maturity | Contract | Summary |\n| --- | --- | --- | --- |\n${table}`
        )
      });
    }
  }

  for (const capability of capabilities) {
    const relatedEvents = events.filter((entry) => entry.capabilities.includes(capability));
    const canonicalId = `capability:${capability}`;
    entities.push(
      makeEntity({
        canonicalId,
        kind: "capability",
        surface: "capabilities",
        localeStrategy: "mirrored",
        title: capability,
        summary: `Capability ${capability}`,
        stability: "public-beta",
        maturity: "beta",
        sourceKind: "support-export",
        sourceRef: sourceRefs.supportMatrix,
        pathEn: `/en/api/capabilities/${capability}`,
        pathRu: `/ru/api/capabilities/${capability}`,
        searchTerms: [capability]
      })
    );
    const list = relatedEvents.map((entry) => `- \`${entry.platform}/${entry.event}\``).join("\n");
    for (const locale of ["en", "ru"]) {
      pages.push({
        locale,
        relativePath: path.join(locale, "api", "capabilities", `${capability}.md`),
        content: renderMarkdownPage(
          {
            title: capability,
            description: `Capability reference for ${capability}`,
            canonicalId,
            surface: "capabilities",
            section: "api",
            locale,
            generated: true,
            editLink: false,
            stability: "public-beta",
            maturity: "beta",
            sourceRef: sourceRefs.supportMatrix,
            translationRequired: false
          },
          `<DocMetaCard surface="capabilities" stability="public-beta" maturity="beta" source-ref="${sourceRefs.supportMatrix}" source-href="${repoBrowserUrl(sourceRefs.supportMatrix)}" />\n\n# ${capability}\n\n${locale === "ru" ? "Связанные runtime events:" : "Related runtime events:"}\n\n${list}`
        )
      });
    }
  }

  for (const locale of ["en", "ru"]) {
    const platformList = [...platforms.keys()]
      .map((platform) => `- [\`${platform}\`](/${locale}/api/platform-events/${platform})`)
      .join("\n");
    const capabilityList = capabilities
      .map((capability) => `- [\`${capability}\`](/${locale}/api/capabilities/${capability})`)
      .join("\n");
    const targetRows = targets
      .map(
        (entry) =>
          `| ${entry.target} | ${compactProductionClass(entry.production_class, locale)} | ${compactRuntimeContract(entry.runtime_contract, entry.target, locale)} | ${compactInstallModel(entry.install_model, locale)} |`
      )
      .join("\n");
    pages.push({
      locale,
      relativePath: path.join(locale, "api", "platform-events", "index.md"),
      content: renderMarkdownPage(
        {
          title: "Platform Events",
          description: "Generated platform event reference",
          canonicalId: "page:api:platform-events:index",
          surface: "platform-events",
          section: "api",
          locale,
          generated: true,
          editLink: false,
          stability: "public-stable",
          maturity: "stable",
          sourceRef: sourceRefs.supportMatrix,
          translationRequired: false
        },
        `# ${locale === "ru" ? "События платформ" : "Platform Events"}\n\n${
          locale === "ru"
            ? "Эта зона показывает event surfaces по платформам и помогает не смешивать stable lane с beta runtime coverage."
            : "This area shows event surfaces by platform and helps you separate the stable lane from wider beta runtime coverage."
        }\n\n${
          locale === "ru"
            ? "- Открывайте её, когда уже знаете target и хотите увидеть event-level contract.\n- Используйте `Capabilities`, когда нужен cross-platform взгляд вместо platform-first view."
            : "- Open this when you already know the target and need the event-level contract.\n- Use `Capabilities` when you want a cross-platform view instead of a platform-first view."
        }\n\n${
          locale === "ru"
            ? "## Когда не нужно начинать отсюда\n\n- Если вы ещё не выбрали target, сначала прочитайте `/guide/choose-a-target` и `/concepts/target-model`."
            : "## When Not To Start Here\n\n- If you have not picked a target yet, start with `/guide/choose-a-target` and `/concepts/target-model`."
        }\n\n${platformList}`
      )
    });
    pages.push({
      locale,
      relativePath: path.join(locale, "api", "capabilities", "index.md"),
      content: renderMarkdownPage(
        {
          title: "Capabilities",
          description: "Generated capability reference",
          canonicalId: "page:api:capabilities:index",
          surface: "capabilities",
          section: "api",
          locale,
          generated: true,
          editLink: false,
          stability: "public-beta",
          maturity: "beta",
          sourceRef: sourceRefs.supportMatrix,
          translationRequired: false
        },
        `# Capabilities\n\n${
          locale === "ru"
            ? "Capabilities дают cross-platform view на runtime behavior, а не только package tree."
            : "Capabilities give you a cross-platform view of runtime behavior, not just a package tree."
        }\n\n${
          locale === "ru"
            ? "- Открывайте эту зону, когда хотите понять не platform name, а само действие или реакцию.\n- Это лучший вход, если вы сравниваете похожее поведение между Claude и Codex."
            : "- Open this area when you care about the behavior itself, not only the platform name.\n- This is the better entry point when you compare similar behavior across Claude and Codex."
        }\n\n${
          locale === "ru"
            ? "## Когда не нужно начинать отсюда\n\n- Если вы ещё не понимаете сами target families, сначала прочитайте `/guide/what-you-can-build` и `/guide/choose-a-target`."
            : "## When Not To Start Here\n\n- If you do not understand the target families yet, start with `/guide/what-you-can-build` and `/guide/choose-a-target`."
        }\n\n${capabilityList}`
      )
    });
    pages.push({
      locale,
      relativePath: path.join(locale, "reference", "target-support.md"),
      content: renderMarkdownPage(
        {
          title: "Target Support",
          description: "Generated target support summary",
          canonicalId: "page:reference:target-support",
          surface: "reference",
          section: "reference",
          locale,
          generated: true,
          editLink: false,
          stability: "public-stable",
          maturity: "stable",
          sourceRef: sourceRefs.targetSupportMatrix,
          translationRequired: false
        },
        `# ${locale === "ru" ? "Поддержка target’ов" : "Target Support"}\n\n${
          locale === "ru"
            ? "Используйте эту страницу, когда нужно быстро понять, какой target production-ready, а какой остаётся packaging-only или workspace-config lane."
            : "Use this page when you need to quickly see which target is production-ready and which remains packaging-only or a workspace-config lane."
        }\n\n${
          locale === "ru"
            ? "## Когда открывать эту матрицу\n\n- Когда уже известны target names, а теперь нужно быстро сравнить их production class.\n- Когда нужно проверить, не путаете ли вы runtime lane с package-only или workspace-config lane."
            : "## When To Open This Matrix\n\n- When you already know the target names and now need to compare their production class quickly.\n- When you need to verify that you are not confusing a runtime lane with a package-only or workspace-config lane."
        }\n\n| Target | Production Class | Runtime Contract | Install Model |\n| --- | --- | --- | --- |\n${targetRows}\n\n${
          locale === "ru"
            ? "Для полной framing-картины свяжите эту матрицу с [Границей поддержки](/ru/reference/support-boundary) и [Моделью target’ов](/ru/concepts/target-model)."
            : "For full framing, pair this matrix with [Support Boundary](/en/reference/support-boundary) and [Target Model](/en/concepts/target-model)."
        }\n`
      )
    });
  }

  return { entities, pages };
}

function compactProductionClass(value, locale) {
  if (value === "production-ready") {
    return locale === "ru" ? "production-ready" : "production-ready";
  }
  if (value === "production-ready package lane") {
    return locale === "ru" ? "package lane" : "package lane";
  }
  if (value === "production-ready runtime lane") {
    return locale === "ru" ? "runtime lane" : "runtime lane";
  }
  if (value === "packaging-only target") {
    return locale === "ru" ? "packaging-only" : "packaging-only";
  }
  return value;
}

function compactRuntimeContract(value, target, locale) {
  if (target === "claude") {
    return locale === "ru" ? "stable runtime subset" : "stable runtime subset";
  }
  if (target === "codex-runtime") {
    return locale === "ru" ? "stable notify runtime" : "stable notify runtime";
  }
  if (target === "codex-package") {
    return locale === "ru" ? "official package only" : "official package only";
  }
  if (target === "gemini") {
    return locale === "ru" ? "packaging, not runtime" : "packaging, not runtime";
  }
  if (target === "cursor" || target === "opencode") {
    return locale === "ru" ? "workspace-config lane" : "workspace-config lane";
  }
  return value;
}

function compactInstallModel(value, locale) {
  if (value.includes("marketplace")) {
    return locale === "ru" ? "marketplace or local" : "marketplace or local";
  }
  if (value.includes("plugin directory")) {
    return locale === "ru" ? "plugin dir or cache" : "plugin dir or cache";
  }
  if (value.includes("repo-local")) {
    return locale === "ru" ? "repo-local" : "repo-local";
  }
  if (value.includes("workspace")) {
    return locale === "ru" ? "workspace config" : "workspace config";
  }
  if (value.includes("copy install")) {
    return locale === "ru" ? "copy install" : "copy install";
  }
  return value;
}
