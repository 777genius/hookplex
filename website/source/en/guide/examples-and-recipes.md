---
title: "Examples And Recipes"
description: "A guided map of the public example repos, starter repos, local runtime references, and skill examples in plugin-kit-ai."
canonicalId: "page:guide:examples-and-recipes"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# Examples And Recipes

Use this page when you want to see what `plugin-kit-ai` looks like in real repositories instead of only reading abstract guidance.

## Read This Page The Right Way

Examples are not the product by themselves. They show how one managed plugin project can end up in different output shapes.

- Start with a production example when you want to see a finished public repo.
- Start with a starter repo when you want the fastest copy-first entrance.
- Start with a local runtime reference when you want to understand the stable Node or Python local-runtime path.
- Start with skill examples only when you are extending the surrounding authoring workflow.

If you still think `plugin-kit-ai` is only a Claude or Codex starter collection, read [One Project, Multiple Targets](/en/guide/one-project-multiple-targets) first and come back here.

## Best First Stops

- Want the clearest finished runtime repo: open [`codex-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-basic-prod) or [`claude-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/claude-basic-prod).
- Want to see packaging and workspace-config outputs: open [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod), [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package), [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic), or [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic).
- Want the fastest clean starting point: open the [canonical starter catalog](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters) and pair it with [Choose A Starter Repo](/en/guide/choose-a-starter).
- Want the stable interpreted runtime story: open the [repo-local runtime catalog](https://github.com/777genius/plugin-kit-ai/tree/main/examples/local).
- Want small supporting patterns around authoring: open the [skill examples catalog](https://github.com/777genius/plugin-kit-ai/tree/main/examples/skills).

## 1. Production Plugin Examples

These are the clearest examples of finished public shapes:

- [`codex-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-basic-prod): Codex runtime production repo
- [`claude-basic-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/claude-basic-prod): Claude production repo
- [`codex-package-prod`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/codex-package-prod): Codex package target
- [`gemini-extension-package`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/gemini-extension-package): Gemini extension packaging target
- [`cursor-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/cursor-basic): Cursor workspace-config target
- [`opencode-basic`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/plugins/opencode-basic): OpenCode workspace-config target

Read these when you want:

- a concrete repo layout
- real rendered outputs
- a truthful public example of what “healthy” looks like

These examples are the fastest way to understand that `plugin-kit-ai` is broader than one runtime. The same managed model can end up as a runtime plugin, a package-style output, or a workspace-config integration.

## 2. Starter Repos

Use starter repos when you want to begin from a known-good baseline instead of from an empty directory.

They are best for:

- first-time setup
- team onboarding
- choosing between Go, Python, Node, Claude, and Codex starting points

If you are still choosing, pair this with [Choose A Starter Repo](/en/guide/choose-a-starter).

Recommended first opens:

- [Starter catalog README](https://github.com/777genius/plugin-kit-ai/blob/main/examples/starters/README.md): fastest overview of all canonical starters
- [`codex-go-starter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters/codex-go-starter): strongest self-contained default for Codex teams
- [`claude-go-starter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters/claude-go-starter): strongest self-contained default for Claude teams
- [`codex-python-starter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters/codex-python-starter): stable Python copy-first path
- [`claude-node-typescript-starter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/starters/claude-node-typescript-starter): stable Node/TypeScript copy-first path

Starter repos are entrance points, not the long-term boundary of the product.

## 3. Local Runtime References

The `examples/local` area shows repo-local Python and Node runtime references.

These are useful when:

- you want to understand the interpreted runtime story more deeply
- you want to compare JavaScript, TypeScript, and Python local-runtime setups
- you need a concrete reference beyond the starter repos

Recommended first opens:

- [Local runtime catalog README](https://github.com/777genius/plugin-kit-ai/blob/main/examples/local/README.md)
- [`codex-python-local`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/local/codex-python-local)
- [`codex-node-local`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/local/codex-node-local)
- [`codex-node-typescript-local`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/local/codex-node-typescript-local)

## 4. Skill Examples

The `examples/skills` area shows supporting skill examples and helper integrations.

These are not the main entrypoint for most plugin authors, but they are valuable when:

- you want to wire docs, review, or formatting helpers into the broader workflow
- you want to understand how adjacent skills can fit around plugin repos

Recommended first opens:

- [Skill examples README](https://github.com/777genius/plugin-kit-ai/blob/main/examples/skills/README.md)
- [`go-command-lint`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/skills/go-command-lint)
- [`cli-wrapper-formatter`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/skills/cli-wrapper-formatter)
- [`docs-only-review`](https://github.com/777genius/plugin-kit-ai/tree/main/examples/skills/docs-only-review)

## Suggested Reading By Goal

- Want the strongest runtime example: start with the Codex or Claude production example, then read [Build A Team-Ready Plugin](/en/guide/team-ready-plugin).
- Want packaging or workspace-config examples: start with Codex package, Gemini, Cursor, or OpenCode examples, then read [Package And Workspace Targets](/en/guide/package-and-workspace-targets).
- Want a clean starting point, not a finished example: go to [Starter Templates](/en/guide/starter-templates).
- Want to choose the product target before looking at repos: read [Choose A Target](/en/guide/choose-a-target).
- Want to understand the central model behind all of those examples: read [Managed Project Model](/en/concepts/managed-project-model).
- Want the full product map first: read [What You Can Build](/en/guide/what-you-can-build).

## Final Rule

Examples should clarify the public contract, not replace it.

Use example repos to see shape, layout, and healthy outputs. Use the rest of the docs to understand what is stable, what is optional, and what the project actually promises.
