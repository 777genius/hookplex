# Migration Registry

This file records pre-`v1` beta-breaking changes that affect user code, generated scaffold output, or release usage.

Before the `v1.0` tag is cut, beta-breaking changes must still be recorded here and in the relevant changelog.

## Entry Format

For each beta-breaking change, record:

- date
- affected surface
- old behavior
- new behavior
- required user action
- changelog reference

## Current Entries

### 2026-03-26: SDK registration moved from root-Claude methods to platform registrars

- Affected surface:
  - `sdk/hookplex`
  - generated scaffold entrypoints
- Old behavior:
  - Claude registration happened through root methods such as `app.OnStop(...)`.
- New behavior:
  - Registration now happens through platform registrars:
    - `app.Claude().OnStop(...)`
    - `app.Claude().OnPreToolUse(...)`
    - `app.Claude().OnUserPromptSubmit(...)`
    - `app.Codex().OnNotify(...)`
- Required user action:
  - Recreate app construction with `hookplex.New(hookplex.Config{...})`.
  - Move Claude hook registration to `app.Claude()`.
  - Register Codex notify handlers through `app.Codex()`.
- Changelog reference:
  - `sdk/hookplex/CHANGELOG.md` `[Unreleased]`
