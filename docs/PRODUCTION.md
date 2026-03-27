# Production Plugin Workflow

This document is the canonical production authoring path for plugin authors using `plugin-kit-ai`.

## Current Target Boundary

- Claude: production-ready within the stable `Stop`, `PreToolUse`, and `UserPromptSubmit` event set
- Codex: production-ready within the stable `Notify` path
- Gemini: packaging-only target through `render|import`; not a production-ready runtime target

Repo-local executable runtime boundary:

| Runtime | Current tier | Production guidance |
|---------|--------------|---------------------|
| `go` | stable | default production path |
| `python` | public-beta | repo-local only, prefer `.venv`, fallback to system Python `3.10+` |
| `node` | public-beta | repo-local only, system Node.js `20+`, TypeScript only after build-to-JS |
| `shell` | public-beta | repo-local only, POSIX shell on Unix, `bash` required on Windows |

Interpreted runtimes are production-hardened for scaffold, validate, launcher execution, and repo-local bootstrap only.
This workflow does not imply support for dependency installation, package management, or packaged distribution through `plugin-kit-ai install`.

## Canonical Production Lane

Run this exact sequence before shipping a plugin repo:

```bash
plugin-kit-ai normalize .
plugin-kit-ai render .
plugin-kit-ai render --check .
plugin-kit-ai validate . --platform <claude|codex> --strict
```

Then run the target-specific smoke:

- Claude: execute the built binary with a documented stable hook payload such as `Stop`
- Codex: execute the built binary with a documented `notify` payload

For interpreted runtimes, add the bootstrap step before `validate --strict`:

- `python`: create `.venv` with `python3 -m venv .venv` when using a project-local runtime
- `node`: run `npm install`; commit `package-lock.json` or an equivalent deterministic lockfile for production repos
- `shell`: ensure the launcher target remains executable on Unix and `bash` is available on Windows

## Claude Release-Ready Path

- Start from `plugin-kit-ai init --platform claude` or `plugin-kit-ai import --from claude`
- Keep `plugin.yaml` as the canonical authoring manifest
- Commit generated `.claude-plugin/plugin.json` and `hooks/hooks.json`
- Treat the stable promise as applying only to `Stop`, `PreToolUse`, and `UserPromptSubmit`
- Treat additional runtime-supported Claude hooks as `public-beta` unless separately promoted

Reference implementation:

- [examples/plugins/claude-basic-prod](../examples/plugins/claude-basic-prod)

## Codex Release-Ready Path

- Start from `plugin-kit-ai init --platform codex` or `plugin-kit-ai import --from codex`
- Keep `plugin.yaml` as the canonical authoring manifest
- Commit generated `.codex-plugin/plugin.json` and `.codex/config.toml`
- Keep `AGENTS.md` repo-local and review it as part of release
- Treat the stable promise as applying only to the `Notify` path

Reference implementation:

- [examples/plugins/codex-basic-prod](../examples/plugins/codex-basic-prod)

## What This Workflow Guarantees

- normalized `plugin.yaml` with no deprecated or unknown fields
- generated native artifacts are in sync
- strict validation passes with no manifest drift
- the committed example-shaped repo can build and execute a deterministic local smoke path

## What It Does Not Guarantee

- external Claude CLI health before hook execution
- external Codex CLI health before `notify` execution
- runtime parity for Gemini
- promotion of runtime-supported beta hooks into the stable promise
- dependency bootstrap or packaged distribution for interpreted runtimes
