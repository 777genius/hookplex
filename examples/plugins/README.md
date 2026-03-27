# Production Plugin Examples

These examples are reference implementations for the current production plugin workflow.

- [claude-basic-prod](./claude-basic-prod): Claude plugin repo with `plugin.yaml`, generated native artifacts, and deterministic local smoke path
- [codex-basic-prod](./codex-basic-prod): Codex plugin repo with `plugin.yaml`, generated native artifacts, and deterministic local smoke path

Use them together with [../../docs/PRODUCTION.md](../../docs/PRODUCTION.md).

These reference repos document the current stable Go-first production path.
Executable `python`, `node`, and `shell` plugins remain `public-beta`, repo-local only, and are covered through scaffold/runtime docs plus polyglot smoke tests rather than checked-in production example repos.
