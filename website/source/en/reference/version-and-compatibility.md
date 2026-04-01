---
title: "Version And Compatibility Policy"
description: "Understand how plugin-kit-ai frames stable versus beta surfaces, install channels, runtime helpers, and release baselines."
canonicalId: "page:reference:version-and-compatibility"
section: "reference"
locale: "en"
generated: false
translationRequired: true
---

# Version And Compatibility Policy

Use this page when you need one compact public answer to a practical question: what should a team treat as the current baseline, what actually carries a compatibility promise, and what should be read as guidance instead of guarantee?

## Choose In 60 Seconds

- Need the current user-facing baseline: start with the newest relevant [release note](/en/releases/), currently [v1.0.6](/en/releases/v1-0-6).
- Need the safest compatibility rule: stable paths carry the normal promise, beta paths are supported but not frozen.
- Need the install rule: Homebrew, npm, and PyPI wrappers install the CLI; they are not runtime APIs or SDK contracts.
- Need the support rule: pair this page with [Support Boundary](/en/reference/support-boundary) and [Target Support](/en/reference/target-support).

## What This Page Helps You Decide

- which version signal should drive your team's decisions
- whether a change is a normal upgrade, a migration, or a path change
- whether you are reading a release note as guidance, a compatibility promise, or both

## The Public Baseline Rule

- The current public baseline is set by the latest relevant user-facing release note.
- Today, the main baseline in the public docs set is [v1.0.6](/en/releases/v1-0-6).
- For older repos or narrower surfaces such as the Go SDK, read the matching release note before you assume the newest note applies unchanged.

## Stable, Beta, Experimental

- `public-stable`: normal production expectation
- `public-beta`: supported, but still moving
- `public-experimental`: useful for early adopters, not a long-term compatibility promise

The practical rule is simple: do not roll beta or experimental behavior into a long-lived team standard without accepting churn.

## Install Channels vs Runtime Contracts

- Homebrew, npm, and PyPI are install channels for the CLI.
- They do not turn wrapper packages into runtime APIs.
- Public runtime and SDK contracts live in the guides, reference rules, and generated [API](/en/api/) surfaces.

## Runtime Helper Compatibility

- Shared runtime helpers such as `plugin-kit-ai-runtime` should be read through the release notes and delivery guidance, not as an isolated permanent promise.
- If the current guidance changes around `--runtime-package`, use the newest release note plus [Choose Delivery Model](/en/guide/choose-delivery-model).
- Do not assume that “new package version exists” means “every repo should upgrade immediately.”

## What Release Notes Mean

- A release note tells you what changed for users, what did not change, and whether a recommendation became stronger.
- A release note is not a universal parity promise across every target, runtime, or language surface.
- The safe reading pattern is:
  1. read the newest matching note
  2. check migration callouts
  3. go back to the matching guide or reference page for the exact mechanics

## What Teams Should Pin In Practice

- Pin the workflow around `doctor`, `render`, and `validate --strict`.
- Pin the chosen runtime and target path in repo docs and CI.
- Pin shared runtime helper versions where the chosen delivery model expects them.
- Keep the team aligned to one published baseline instead of mixing several half-remembered release states.

## What Not To Assume

- do not assume every target has the same compatibility promise
- do not assume wrapper packages are SDKs
- do not assume a beta convenience flow has the same long-term contract as the stable path
- do not assume one release note replaces the need to check support and target boundaries

## Best First Stops

- Need the maturity vocabulary: [Stability Model](/en/concepts/stability-model)
- Need the exact support line: [Support Boundary](/en/reference/support-boundary)
- Need the exact target matrix: [Target Support](/en/reference/target-support)
- Need the current change journal: [Releases](/en/releases/)
- Need the rollout path across live repos: [Upgrade And Migration Playbook](/en/guide/upgrade-and-migration-playbook)
