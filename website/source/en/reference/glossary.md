---
title: "Glossary"
description: "Canonical terms used across the public plugin-kit-ai documentation."
canonicalId: "page:reference:glossary"
section: "reference"
locale: "en"
generated: false
translationRequired: true
---

# Glossary

## Authored State

The project source of truth that lives in the package-standard layout. `generate` turns this state into target-specific output files.

## Generated Target Files

Files produced for a target after generating. They are not the preferred long-term source of truth.

## Lane

A practical path with its own operating rules. Examples include the default Go path, the Node/TypeScript local runtime path, and workspace-configuration paths.

## Target

The integration or delivery surface you are aiming at, such as `codex-runtime`, `claude`, `codex-package`, `gemini`, `opencode`, or `cursor`.

## Runtime Lane

A path where the project owns executable plugin behavior directly and runtime choice, handler behavior, and strict validation matter the most.

## Package Or Extension Lane

A path focused on producing the right package or extension artifacts rather than running a local executable plugin.

## Workspace-Config Lane

A path where the main product is repo-owned configuration, not an executable plugin runtime.

## Wrapper Install Channel

A way to install the CLI, such as Homebrew, npm, or PyPI. It is not a public runtime API.

## Shared Runtime Package

The `plugin-kit-ai-runtime` dependency used by approved Python and Node flows instead of copying helper files into every repo.

## Support Boundary

The public line between what the project treats as stable, what remains beta, and what is intentionally outside the long-term promise.

## Readiness Gate

The command or flow you should treat as the public signal that a repo is healthy. For most projects this is `validate --strict`, often paired with `doctor` and `generate`.

## Handoff

The point where a repo, artifact, or package is ready for another teammate, another machine, or another user without hidden knowledge.

Pair this glossary with [Target Model](/en/concepts/target-model), [Support Boundary](/en/reference/support-boundary), and [Production Readiness](/en/guide/production-readiness).
