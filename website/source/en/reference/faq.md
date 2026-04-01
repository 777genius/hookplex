---
title: "FAQ"
description: "Common questions about choosing paths, wrappers, targets, and authoring flows."
canonicalId: "page:reference:faq"
section: "reference"
locale: "en"
generated: false
translationRequired: true
---

# FAQ

## Choose In 60 Seconds

- Need the safest default: start with Go.
- Need the main supported non-Go path: choose Node/TypeScript.
- Need a Python-first repo-local path on purpose: choose Python.
- Need to know whether something is an API or just an install path: wrappers install the CLI, they are not runtime APIs.
- Need the support promise, not the quick answer: go to [Support Boundary](/en/reference/support-boundary) and [Target Support](/en/reference/target-support).

## Should I Start With Go, Python, Or Node?

Start with Go unless you have a real reason not to. Choose Node/TypeScript for the main supported non-Go path. Choose Python when the plugin stays local to the repo and your team is already Python-first.

## Are npm And PyPI `plugin-kit-ai` Packages Runtime APIs?

No. They are ways to install the CLI. They are not public runtime APIs and they are not SDKs.

## Should I Choose A Starter Or `init`?

Choose a starter when you want the fastest known-good repo layout. Choose `plugin-kit-ai init` when you already have a repo or want to build from first principles.

## Do Starter Names Lock Me To One Agent Family Forever?

No. The starter name tells you the first correct path, not the permanent boundary of the project.

## When Should I Use Bundle Commands?

Use bundle commands when you need portable Python or Node artifacts that another machine can fetch or install. Do not confuse them with the main CLI install path.

## Can I Keep Native Target Files As My Source Of Truth?

That is not the intended long-term model. The project source of truth should live in the package-standard layout, and target files should be generated outputs.

## Is `render` Optional?

Not if you want the managed project model. `render` is part of the workflow, not an optional extra.

## Is `validate --strict` Optional?

Treat it as the main readiness check, especially for local Python and Node runtime projects.

## Are All Targets Equally Stable?

No. Runtime, packaging, extension, and workspace-configuration targets do not all carry the same support promise.

## When Should I Read API Instead Of Guides?

Read API when you already know which surface you need and now want exact commands, packages, runtime helpers, events, or capabilities. Start with guides if you are still choosing a path.

## What This Page Is Good For

- quick path-selection answers
- quick wrapper-versus-runtime answers
- quick stability questions before you go deeper into the full reference

See [Support Boundary](/en/reference/support-boundary), [Target Support](/en/reference/target-support), and [Authoring Workflow](/en/reference/authoring-workflow).
