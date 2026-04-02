---
title: "What You Can Build"
description: "A broad public overview of how one plugin repo can grow into more supported outputs."
canonicalId: "page:guide:what-you-can-build"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# What You Can Build

This page explains the main promise of `plugin-kit-ai`: start in one repo, then expand that same repo to more supported outputs as the product grows.

<MermaidDiagram
  :chart="`
flowchart TD
  Product[One repo] --> Runtime[Runtime plugins]
  Product --> Multi[More supported outputs later]
  Product --> Bundle[Portable bundle handoff]
  Product --> Shared[Shared runtime package]
  Product --> Package[Package and extension targets]
  Product --> Workspace[Workspace config targets]
`"
/>

## 1. One Repo, Many Supported Outputs

This is the core product promise.

- Start with one plugin repo.
- Add the outputs you actually need as the product grows.
- Keep one workflow through `render`, `validate`, and CI.
- Do not assume every target has the same runtime guarantees.

The safe mental model is:

- same repo
- same core workflow
- many supported outputs
- different support depth by target

## 2. Start With The Strongest First Repo

Most teams should start with a Codex runtime repo on Go.

That first repo gives you:

- the strongest production-oriented starting point
- one repo workflow instead of hand-edited target files
- a clear path through `render` and `validate --strict`

If your stack already dictates the runtime, the same first repo model also supports:

- Go for the strongest default production contract
- Node/TypeScript for the mainstream non-Go stable lane
- Python for repo-local Python-first teams

## 3. Add Claude When Hooks Are The Real Requirement

Use the Claude lane when Claude hooks are the actual product requirement.

This is the right choice when:

- you need Claude-specific runtime hooks
- the stable Claude subset is enough for your plugin
- you want a stronger authoring contract than native file editing

## 4. Expand To More Supported Outputs Later

Once the first repo is working, the same repo can grow into:

- Claude hooks outputs
- Codex package outputs
- Gemini extension packaging
- OpenCode and Cursor workspace/config outputs
- portable bundle delivery for supported Python and Node repos

That is the real cross-target story: one repo, one workflow, more supported outputs over time.

## 5. Team-Ready Plugin Repositories

`plugin-kit-ai` is not only about scaffolding. It is also about getting to a repo another teammate can understand, validate, and ship.

That means the system supports:

- strict readiness gates
- CI-friendly flows
- explicit lane and target choices
- predictable handoff between authors and downstream consumers

## 6. Portable Python And Node Handoff Bundles

For supported Python and Node lanes, you can move beyond local authoring and produce portable bundle handoff artifacts.

This matters when:

- the delivery model needs fetched artifacts instead of a live repo
- you want a cleaner downstream install story for interpreted runtime lanes
- you are using the bundle publish/fetch flow as part of release handoff

See [Bundle Handoff](/en/guide/bundle-handoff) for the actual public flow.

## 7. Shared Runtime Package Flows

Python and Node helper behavior can live either:

- in vendored helper files inside the repo
- in the shared `plugin-kit-ai-runtime` package

This gives teams a supported path for:

- reusable runtime helpers across multiple repos
- cleaner dependency upgrades
- a standardized helper API without copying scaffolded files by hand

## 8. Package, Extension, And Workspace-Config Targets

Not every public shape is a repo-local runtime plugin.

`plugin-kit-ai` also covers:

- packaging-oriented lanes
- extension-style targets
- workspace-config integration targets

These targets matter when the end product is packaging or configuration, not an executable plugin.

See [Package And Workspace Targets](/en/guide/package-and-workspace-targets) before you treat these targets like runtime plugins.

## 9. Read In This Order

If you are still deciding what to do:

1. read this page
2. use [Quickstart](/en/guide/quickstart) or [Choose A Starter Repo](/en/guide/choose-a-starter)
3. read [One Project, Multiple Targets](/en/guide/one-project-multiple-targets) when you want the honest explanation of expansion
4. read [Target Model](/en/concepts/target-model) only once you are ready to compare output types precisely

Pair this page with [Examples And Recipes](/en/guide/examples-and-recipes), [Choose A Starter Repo](/en/guide/choose-a-starter), [Choose Delivery Model](/en/guide/choose-delivery-model), [Bundle Handoff](/en/guide/bundle-handoff), [Package And Workspace Targets](/en/guide/package-and-workspace-targets), and [API Surfaces](/en/api/).
