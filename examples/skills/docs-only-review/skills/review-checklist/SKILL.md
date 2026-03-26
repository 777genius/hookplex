---
name: review-checklist
description: Apply a short, consistent review checklist before merging changes.
execution_mode: docs_only
supported_agents:
  - claude
  - codex
allowed_tools: []
compatibility:
  repo_required: true
  notes:
    - Works best when the agent can inspect diffs and the touched files locally.
---

# Review Checklist

## What it does

Provides a repeatable review checklist for code, docs, and generated artifact changes.

## When to use

Use this when you want a quick review pass before merging or handing work off to another maintainer.

## How to run

Read the checklist, inspect the changed files, and record concrete findings or the absence of findings.

## Constraints

- This skill is instructional only and has no command to execute.
- Keep findings concrete and tied to changed files or behaviors.
