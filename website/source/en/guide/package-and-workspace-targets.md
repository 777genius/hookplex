---
title: "Package And Workspace Targets"
description: "How to use Codex package, Gemini, OpenCode, and Cursor targets without confusing them with executable plugin paths."
canonicalId: "page:guide:package-and-workspace-targets"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# Package And Workspace Targets

Not every `plugin-kit-ai` target is an executable plugin path.

Read this page before you choose `codex-package`, `gemini`, `opencode`, or `cursor`, because these targets solve different problems than `codex-runtime` or `claude`.

## Choose In 60 Seconds

- choose `codex-runtime` or `claude` when the product is an executable plugin
- choose `codex-package` or `gemini` when the product is a packaged target or extension artifact
- choose `opencode` or `cursor` when the product is repo-owned workspace configuration

## The Short Rule

- runtime target = executable plugin behavior
- package or extension target = publishable or installable artifact
- workspace-config target = repo-owned integration and configuration files

## Best Defaults

- Best package-style default: `codex-package`
- Best extension-style default: `gemini`
- Best workspace-config defaults: `cursor` or `opencode`

## Best First Examples

- `codex-package`: [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod)
- `gemini`: [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package)
- `cursor`: [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic)
- `opencode`: [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic)

## Quick Decision Tree

1. Need something users install or publish as an artifact? Choose `codex-package` or `gemini`.
2. Need the repo to own editor or tool integration files? Choose `cursor` or `opencode`.
3. Need the repo itself to run plugin logic? Stop and go back to `codex-runtime` or `claude`.

## Codex Package

Use `codex-package` when the end result is a Codex package, not an executable plugin repo.

This is useful when:

- packaging is the real delivery contract
- you want the project source to stay managed in one place
- you do not want to pretend this target has the same runtime contract as `codex-runtime`

Best first example: [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod)

## Gemini

Use `gemini` when the goal is a Gemini CLI extension package.

This target is intentionally packaging-oriented.

Treat it as:

- a full extension-packaging lane through `render`, `import`, and `validate`
- not the main runtime path
- something you choose when Gemini extension artifacts are the actual product

Best first example: [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package)

## OpenCode

Use `opencode` when the repo owns OpenCode workspace configuration and related project assets.

This target is valuable when:

- the project needs managed `opencode.json`
- the repo should own workspace-level MCP and config shape
- you want a documented config authoring path instead of hand-edited files

Do not confuse that with the strongest runtime contract.

Best first example: [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic)

## Cursor

Use `cursor` when the repo should manage Cursor workspace configuration.

The documented subset includes:

- `.cursor/mcp.json`
- project-root `.cursor/rules/**`
- optional shared root `AGENTS.md`

This is a workspace-configuration target, not the main runtime path.

Best first example: [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic)

## Practical Decision Rule

Choose these targets when the project output is:

- package artifacts
- extension packaging
- workspace config

Do not choose them just because they sound close to a runtime path.

If what you really need is executable plugin behavior, go back to [Choosing Runtime](/en/concepts/choosing-runtime) and start there.

## What Not To Expect

- Do not expect `codex-package` or `gemini` to behave like executable runtime plugins.
- Do not expect `cursor` or `opencode` to replace the main runtime path when what you really need is plugin logic.
- Do not choose these targets only because the ecosystem name looks familiar. Choose them only when the output shape is the actual product requirement.

## Readiness Rule

For these targets, the healthy repo rule is still the same:

- the project source stays in the package-standard layout
- rendered files are outputs
- `render --check` and `validate --strict` are the core checks

## Next Reading

- Read [Choose A Target](/en/guide/choose-a-target) if you want the full target map first.
- Read [Examples And Recipes](/en/guide/examples-and-recipes) if you want to compare runtime, package, and workspace examples side by side.
- Read [Target Support](/en/reference/target-support) when you need the compact support matrix.

## Pair It With

Read this page with [Target Model](/en/concepts/target-model), [Target Support](/en/reference/target-support), and [Support Boundary](/en/reference/support-boundary).
