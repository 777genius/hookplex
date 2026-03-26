---
name: lint-repo
description: Run a fast repository lint pass and report actionable failures.
execution_mode: command
supported_agents:
  - claude
  - codex
allowed_tools:
  - bash
  - go
command: go run ./cmd/lint-repo
runtime: go
compatibility:
  requires:
    - go >=1.25
  supported_os:
    - darwin
    - linux
  repo_required: true
  notes:
    - Run from the repository root so the command can discover files consistently.
safe_to_retry: true
writes_files: false
produces_json: false
---

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
