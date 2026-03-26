# lint-repo

> Codex render for `skills/lint-repo/SKILL.md`. Keep the authored skill canonical and regenerate this file when it changes.

## Summary

Run a fast repository lint pass and report actionable failures.
## Command

- Runtime: `go`
- Invocation: `go run ./cmd/lint-repo`
## Compatibility
- Requires: go >=1.25
- Supported OS: darwin, linux
- Requires a repository checkout
- Run from the repository root so the command can discover files consistently.
## Allowed tools
- `bash`
- `go`

## Canonical instructions

# Lint Repository

## What it does

Runs a lightweight repository lint pass and prints actionable failures.

## When to use

Use this when you want a quick quality gate before a commit, release rehearsal, or broad refactor.

## How to run

Run `go run ./cmd/lint-repo` from the repository root.

## Constraints

- This is a non-interactive command.
- It assumes the current directory is a repository checkout.
- Diagnostics go to stdout and stderr; fix the reported issues before rerunning.
