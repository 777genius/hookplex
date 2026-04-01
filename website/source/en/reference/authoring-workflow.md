---
title: "Authoring Workflow"
description: "The main workflow from init to render, validate, test, and handoff."
canonicalId: "page:reference:authoring-workflow"
section: "reference"
locale: "en"
generated: false
translationRequired: true
---

# Authoring Workflow

The recommended workflow is intentionally simple:

```text
init -> render -> validate --strict -> test -> handoff
```

## What Each Step Means

| Step | Purpose |
| --- | --- |
| `init` | Create a package-standard project layout |
| `render` | Generate target artifacts from the project source |
| `validate --strict` | Run the main readiness check |
| `test` | Run stable smoke tests where applicable |
| `export` / bundle flow | Produce handoff artifacts for supported Python and Node cases |

## Rules That Keep The Repo Healthy

- the project source lives in the package-standard project layout
- generated target files are outputs, not the long-term source of truth
- strict validation is a required check, not an optional extra

## When The Workflow Changes

The workflow can widen for special cases:

- `doctor` and `bootstrap` matter for Python and Node runtime paths
- `import` and `normalize` matter when migrating native config into the managed project model
- bundle commands matter for portable Python and Node handoff flows

Start with [Quickstart](/en/guide/quickstart) when you need the shortest path, or [Migrate Existing Native Config](/en/guide/migrate-existing-config) when you are entering from legacy target files.
