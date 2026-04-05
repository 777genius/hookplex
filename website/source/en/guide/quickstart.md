---
title: "Quickstart"
description: "The fastest supported path to a working plugin-kit-ai project."
canonicalId: "page:guide:quickstart"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# Quickstart

This is the shortest supported path when you want one plugin repo that can later expand to more supported targets.

Start with one clear path first. Expand the same repo later when you need Claude, Codex package, Gemini, or workspace/config outputs.

## If You Only Read One Thing

Start with the default Go path unless you already know you need Claude hooks, Node/TypeScript, or Python.

Do not confuse the first path with the final product boundary: the repo can grow later.

## Recommended Default

If you do not have a strong reason to choose another path, start here:

```bash
brew install 777genius/homebrew-plugin-kit-ai/plugin-kit-ai
plugin-kit-ai version
plugin-kit-ai init my-plugin
cd my-plugin
plugin-kit-ai generate .
plugin-kit-ai validate . --platform codex-runtime --strict
```

That gives you the strongest default path today: a Go-based Codex runtime repo that stays easy to validate, hand off, and expand later.

## Why This Is The Default

- one repo from day one
- the cleanest production story today
- the easiest base for later expansion to other supported outputs

## What Expands Later

- You still keep one repo and one validation workflow as you add more outputs.
- You can generate supported outputs for Claude, Codex, Gemini, and other targets from the same repo.
- Support depth depends on the target you add.
- Runtime plugins, package outputs, and workspace-managed config do not all behave the same way.

## Choose The First Path

| If you want | Best starting path |
| --- | --- |
| Strongest production path | `codex-runtime` with `--runtime go` |
| Repo-local TypeScript plugin | `codex-runtime --runtime node --typescript` |
| Repo-local Python plugin | `codex-runtime --runtime python` |

Choose `claude` first only when Claude hooks are already the real product requirement.

Choose package, extension, and workspace/config targets later as expansion paths, not as the first default decision.

## Common First Commands

```bash
plugin-kit-ai init my-plugin --platform codex-runtime --runtime node --typescript
plugin-kit-ai doctor ./my-plugin
plugin-kit-ai bootstrap ./my-plugin
plugin-kit-ai generate ./my-plugin
plugin-kit-ai validate ./my-plugin --platform codex-runtime --strict
```

## Read This Before Choosing Python Or Node

- Python and Node are supported first-class for the stable repo-local subset.
- They still require Python `3.10+` or Node.js `20+` on the machine that runs the plugin.
- Go remains the recommended default when you want the cleanest production and distribution story.

## After Quickstart

- Continue with [Build Your First Plugin](/en/guide/first-plugin) if you want the narrowest recommended tutorial.
- Continue with [Build A Python Runtime Plugin](/en/guide/python-runtime) if your team is Python-first and the plugin stays repo-local.
- Continue with [What You Can Build](/en/guide/what-you-can-build) if you want to see how the same repo can later cover more outputs.
- Continue with [Choose A Starter Repo](/en/guide/choose-a-starter) if you want to start from a template instead of a blank repo.
- Continue with [One Project, Multiple Targets](/en/guide/one-project-multiple-targets) when you are ready to expand beyond the first path.
- Continue with [Choose A Target](/en/guide/choose-a-target) only after you already understand the basic product shape.

See [Choosing Runtime](/en/concepts/choosing-runtime) for the decision model and [Installation](/en/guide/installation) for CLI install channels.
