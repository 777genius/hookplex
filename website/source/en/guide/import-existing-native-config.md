---
title: "Import Existing Native Config"
description: "Bring hand-managed native config into one repo-owned workflow."
canonicalId: "page:guide:import-existing-native-config"
section: "guide"
locale: "en"
generated: false
translationRequired: true
---

# Import Existing Native Config

Use this path when you already have native Claude, Codex, Gemini, OpenCode, or Cursor configuration and want to bring it into one repo-owned workflow.

## Import Shape

```bash
plugin-kit-ai import ./native-plugin --from codex-runtime
plugin-kit-ai normalize ./native-plugin
plugin-kit-ai render ./native-plugin
plugin-kit-ai validate ./native-plugin --platform codex-runtime --strict
```

## Goal Of The Import

The goal is not to keep native files as your long-term editing surface. The goal is to move into a repo-owned workflow and let `render` produce the target artifacts deterministically.

## Good Import Discipline

- import once to establish the project model
- normalize when you need the package-standard shape cleaned up
- render to regenerate target artifacts from the project source
- validate strictly before trusting the imported project

## When This Is Worth It

- your team already has native config drift
- you want one repo-owned workflow
- you want to stop hand-editing target artifacts as if they were the main project source

See [Support Boundary](/en/reference/support-boundary) and [CLI Reference](/en/api/cli/) for the formal contract around this flow.
