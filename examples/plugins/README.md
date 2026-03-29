# Production Plugin Examples

These examples are reference implementations for the current production plugin workflow.

- [claude-basic-prod](./claude-basic-prod): Claude plugin repo with `plugin.yaml`, generated native artifacts, and deterministic local smoke path
- [codex-basic-prod](./codex-basic-prod): Codex runtime lane repo with `plugin.yaml`, generated `.codex/config.toml`, and deterministic local notify smoke path
- [codex-package-prod](./codex-package-prod): official Codex package lane with `plugin.yaml`, rendered `.codex-plugin/plugin.json`, and skills-first bundle output
- [gemini-extension-package](./gemini-extension-package): Gemini CLI extension repo with `plugin.yaml`, rendered `gemini-extension.json`, shared MCP, and packaging-only validation coverage
- [opencode-basic](./opencode-basic): OpenCode workspace-config repo with `plugin.yaml`, rendered `opencode.json`, shared MCP, and mirrored portable skills

Use them together with [../../docs/PRODUCTION.md](../../docs/PRODUCTION.md).
For copy-first Python/Node starter repos, see [../starters/README.md](../starters/README.md).
For deeper repo-local Python/Node entrance references, see [../local/README.md](../local/README.md).

These reference repos document the current stable Go-first production path.
Their authored source of truth is `plugin.yaml` plus `targets/<platform>/...`; committed native Claude/Codex/Gemini/OpenCode files are rendered managed artifacts.
Gemini and OpenCode remain packaging/workspace-config-only in this reference set. Executable `python` and `node` plugins are now the stable repo-local local-runtime subset and are covered through scaffold/runtime docs plus polyglot smoke tests rather than checked-in production example repos. Launcher-based `shell` authoring remains `public-beta`.
