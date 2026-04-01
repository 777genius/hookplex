# Contributing to plugin-kit-ai

Thanks for contributing.

## Start With The Right Path

- Bugs, docs gaps, and feature ideas: open a GitHub issue first at <https://github.com/777genius/plugin-kit-ai/issues/new/choose>.
- Sensitive security issues: do **not** open a public issue. Follow [SECURITY.md](./SECURITY.md).
- Small focused fixes: a pull request is welcome directly if the change is obvious and self-contained.

## Good Contributions

Good contributions usually do one of these:

- clarify public docs, guides, or support boundaries
- improve examples or starter flows
- fix user-facing CLI, SDK, runtime, or generated API behavior
- tighten tests, validation, or release safety around the public contract

## Keep The Change Focused

- Prefer one concern per pull request.
- Update docs when user-facing behavior changes.
- Update generated docs or registries when the public docs pipeline says they drifted.
- Avoid mixing unrelated refactors with a public behavior change.

## Checks

Run the checks that match the area you changed.

For docs-only changes:

```bash
cd website
pnpm run docs:check
pnpm run docs:smoke-ui
```

For SDK or CLI changes, run the relevant tests for the touched area and the repository-level checks your change could affect.

## Public Contract First

If your change touches:

- target behavior
- support boundaries
- install or runtime guidance
- generated API reference

make sure the public docs stay aligned with the code and release story.

## Useful Public Docs

- Public docs site: <https://777genius.github.io/plugin-kit-ai/>
- Managed project model: <https://777genius.github.io/plugin-kit-ai/en/concepts/managed-project-model>
- Team adoption path: <https://777genius.github.io/plugin-kit-ai/en/guide/team-adoption>
- Help and contribution guide: <https://777genius.github.io/plugin-kit-ai/en/reference/get-help-and-contribute>
