---
title: "Troubleshooting"
description: "The most common failure modes when installing, rendering, validating, or bootstrapping plugin-kit-ai projects."
canonicalId: "page:reference:troubleshooting"
section: "reference"
locale: "en"
generated: false
translationRequired: true
---

# Troubleshooting

## Check These First

1. Run `plugin-kit-ai doctor <path>`.
2. Confirm you chose the correct target and runtime path.
3. Run `render` again before assuming generated files are correct.
4. Treat `validate --strict` as signal, not noise.

## The CLI Installs But Does Not Run

Check that the binary is actually on your shell `PATH`. If you used npm or PyPI to install the CLI, verify that it downloaded the published binary successfully instead of assuming the package itself is the runtime.

## Python Or Node Runtime Projects Fail Early

Check the real runtime first:

- Python runtime projects require Python `3.10+`
- Node runtime projects require Node.js `20+`

Use `plugin-kit-ai doctor <path>` before assuming the project itself is broken.

## `validate --strict` Fails

Treat this as signal, not noise. The point of strict validation is to catch drift or readiness problems before you treat the project as healthy.

Common causes:

- generated artifacts are stale because `render` was skipped
- the selected platform does not match the project source
- the runtime path needs bootstrap or environment fixes

First check:

- run `plugin-kit-ai render <path>`
- rerun `plugin-kit-ai validate <path> --platform <target> --strict`
- verify that the target really matches the project you are trying to build

## `render` Output Looks Different Than Expected

That usually means the project source and your mental model have drifted apart. Re-check the package-standard layout instead of hand-editing generated target files to “fix” the output.

## I Am Unsure Which Path I Should Use

Start with the default Go path if you want the strongest contract. Move to Node/TypeScript or Python only when the repo-local runtime tradeoff is real and intentional.

## I Already Have Native Config Files

Use the migration flow:

```bash
plugin-kit-ai import ./native-plugin --from codex-runtime
plugin-kit-ai normalize ./native-plugin
plugin-kit-ai render ./native-plugin
plugin-kit-ai validate ./native-plugin --platform codex-runtime --strict
```

## When To Stop Debugging And Go Back To The Guides

- If you still are not sure whether you need runtime, package, or workspace-config outputs, go back to [Choose A Target](/en/guide/choose-a-target).
- If you are not sure whether you should be on Go, Node, or Python, go back to [Choosing Runtime](/en/concepts/choosing-runtime).
- If you are not sure whether your repo should start from a starter or `init`, go back to [Choose A Starter Repo](/en/guide/choose-a-starter).

See [Authoring Workflow](/en/reference/authoring-workflow), [FAQ](/en/reference/faq), and [Migrate Existing Native Config](/en/guide/migrate-existing-config).
