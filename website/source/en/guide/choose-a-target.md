---
title: "Choose A Target"
description: "A practical public guide for choosing between Codex runtime, Claude, Codex package, Gemini, OpenCode, and Cursor."
canonicalId: "page:guide:choose-a-target"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# Choose A Target

Use this page when you already know you want `plugin-kit-ai`, but you are still deciding which target matches the thing you are building.

## Choose In 60 Seconds

- choose `codex-runtime` when you want an executable plugin with the strongest default story
- choose `claude` when Claude hooks are the real product requirement
- choose `codex-package` when the output is a Codex package, not an executable plugin repo
- choose `gemini` when the output is a Gemini extension package
- choose `opencode` or `cursor` when the repo owns workspace configuration instead of an executable plugin

## The Main Distinction

- Choose a **runtime target** when the repo should run plugin logic directly.
- Choose a **package or extension target** when the output is something to install or publish.
- Choose a **workspace-config target** when the repo should own integration files and configuration, not executable plugin behavior.

## Best Defaults

- Best default for a real executable plugin: `codex-runtime`
- Best default when Claude hooks are the first real requirement: `claude`
- Best package-style target: `codex-package`
- Best extension-style target: `gemini`
- Best workspace-config targets: `cursor` or `opencode`

## Target Directory

| Target | Choose it when | What it is not |
| --- | --- | --- |
| `codex-runtime` | You want the default executable plugin path | Not a packaging-only target |
| `claude` | You need Claude hooks specifically | Not the default Codex path |
| `codex-package` | You need Codex packaging output | Not an executable runtime plugin |
| `gemini` | You are shipping a Gemini extension package | Not the main runtime path |
| `opencode` | You want repo-owned OpenCode workspace config | Not an executable runtime plugin |
| `cursor` | You want repo-owned Cursor workspace config | Not an executable runtime plugin |

## Quick Decision Tree

1. Do you need the repo to execute plugin logic? Choose `codex-runtime` or `claude`.
2. Do you need a package or extension artifact instead of a running plugin? Choose `codex-package` or `gemini`.
3. Do you need repo-owned editor or tool integration files? Choose `cursor` or `opencode`.
4. Are you unsure? Start with `codex-runtime`.

## Start Here By Goal

- Need the strongest default for a real plugin repo: start with `codex-runtime`, then read [Choose A Starter Repo](/en/guide/choose-a-starter).
- Need Claude hooks: start with `claude`, then read [Starter Templates](/en/guide/starter-templates).
- Need package or extension artifacts: start with `codex-package` or `gemini`, then read [Package And Workspace Targets](/en/guide/package-and-workspace-targets).
- Need workspace config owned by the repo: start with `opencode` or `cursor`, then read [Examples And Recipes](/en/guide/examples-and-recipes).

## Best First Example By Target

- `codex-runtime`: [`codex-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-basic-prod)
- `claude`: [`claude-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/claude-basic-prod)
- `codex-package`: [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod)
- `gemini`: [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package)
- `cursor`: [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic)
- `opencode`: [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic)

## Safe Default

If you are unsure, start with `codex-runtime` and the default Go path.

That gives you the cleanest production starting point before you choose a narrower or more specialized target.

## Next Reading

- Read [Managed Project Model](/en/concepts/managed-project-model) if you want the central product model before you commit to a target.
- Read [Choosing Runtime](/en/concepts/choosing-runtime) if you picked a runtime target and still need to choose Go, Python, Node, or shell.
- Read [Package And Workspace Targets](/en/guide/package-and-workspace-targets) if you are deciding between packaging and workspace-config targets.
- Read [Examples And Recipes](/en/guide/examples-and-recipes) if you want to see real repos for each product shape.
- Read [Target Support](/en/reference/target-support) if you want the compact support matrix.
