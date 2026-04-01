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

## Short Rule

- choose `codex-runtime` when you want an executable plugin with the strongest default story
- choose `claude` when Claude hooks are the real product requirement
- choose `codex-package` when the output is a Codex package, not an executable plugin repo
- choose `gemini` when the output is a Gemini extension package
- choose `opencode` or `cursor` when the repo owns workspace configuration instead of an executable plugin

## Target Directory

| Target | Choose it when | What it is not |
| --- | --- | --- |
| `codex-runtime` | You want the default executable plugin path | Not a packaging-only target |
| `claude` | You need Claude hooks specifically | Not the default Codex path |
| `codex-package` | You need Codex packaging output | Not an executable runtime plugin |
| `gemini` | You are shipping a Gemini extension package | Not the main runtime path |
| `opencode` | You want repo-owned OpenCode workspace config | Not an executable runtime plugin |
| `cursor` | You want repo-owned Cursor workspace config | Not an executable runtime plugin |

## Start Here By Goal

- Need the strongest default for a real plugin repo: start with `codex-runtime`
- Need Claude hooks: start with `claude`
- Need package or extension artifacts: start with `codex-package` or `gemini`
- Need workspace config owned by the repo: start with `opencode` or `cursor`

## Safe Default

If you are unsure, start with `codex-runtime` and the default Go path.

That gives you the cleanest production starting point before you choose a narrower or more specialized target.

## Next Reading

- Read [Choosing Runtime](/en/concepts/choosing-runtime) if you picked a runtime target and still need to choose Go, Python, Node, or shell.
- Read [Package And Workspace Targets](/en/guide/package-and-workspace-targets) if you are deciding between packaging and workspace-config targets.
- Read [Examples And Recipes](/en/guide/examples-and-recipes) if you want to see real repos for each product shape.
- Read [Target Support](/en/reference/target-support) if you want the compact support matrix.
