---
title: "Migrate Existing Native Config"
description: "Move from hand-managed native config to the package-standard project model."
canonicalId: "page:guide:migrate-existing-config"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# Migrate Existing Native Config

Use this path when you already have native Claude, Codex, Gemini, OpenCode, or Cursor configuration and want to move into the package-standard project model.

## Migration Shape

```bash
plugin-kit-ai import ./native-plugin --from codex-runtime
plugin-kit-ai normalize ./native-plugin
plugin-kit-ai render ./native-plugin
plugin-kit-ai validate ./native-plugin --platform codex-runtime --strict
```

## Goal Of The Migration

The goal is not to keep native files as your long-term source of truth. The goal is to move into a repo-owned project model and let `render` produce the target artifacts deterministically.

## Good Migration Discipline

- import once to establish the project model
- normalize when you need the package-standard shape cleaned up
- render to regenerate target artifacts from the project source
- validate strictly before trusting the migrated project

## When This Is Worth It

- your team already has native config drift
- you want a single repo-owned source of truth
- you want to stop hand-editing target artifacts as if they were the main project source

See [Support Boundary](/en/reference/support-boundary) and [CLI Reference](/en/api/cli/) for the formal contract around this flow.
