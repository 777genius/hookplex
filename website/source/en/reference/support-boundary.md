---
title: "Support Boundary"
description: "A compact guide to what plugin-kit-ai treats as stable, beta, and intentionally out of scope."
canonicalId: "page:reference:support-boundary"
section: "reference"
locale: "en"
generated: false
translationRequired: true
---

# Support Boundary

This page answers one practical question: what can you safely promise today, and what should you describe more carefully?

## Safe Defaults

- Go is the recommended production path.
- `validate --strict` is the main readiness check for local Python and Node runtime projects.
- The CLI install wrappers are ways to install the CLI, not runtime APIs.
- One repo can cover many supported outputs, but support depth still depends on the target.

## Stable By Default

- the main public CLI contract
- the recommended Go SDK path
- the stable local Python and Node subset on supported runtime targets
- the targets explicitly marked stable in the generated support matrix

Today that means:

- Claude is production-ready only for the stable default hook subset, not for every Claude package surface
- Codex is production-ready for the `Notify` runtime lane and for the official `codex-package` lane
- Gemini extension packaging is production-ready, while the optional Gemini Go runtime remains `public-beta`
- OpenCode and Cursor are supported as workspace-config lanes, not as production-ready runtime targets

## Use Carefully

- beta paths that are still evolving
- workspace-configuration and packaging targets when what you really need is executable plugin behavior
- install wrappers when what you really want is a runtime or SDK API

## Out Of Scope

- treating every target as if it had the same runtime guarantees
- treating wrapper packages as SDKs or runtime contracts
- assuming experimental surfaces carry long-term compatibility promises

Pair this page with [Version And Compatibility Policy](/en/reference/version-and-compatibility), [Target Support](/en/reference/target-support), and [Stability Model](/en/concepts/stability-model).
