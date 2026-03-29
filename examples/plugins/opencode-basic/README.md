# opencode-basic

Reference OpenCode workspace-config example for `plugin-kit-ai`.

This example demonstrates the current finished OpenCode workspace/config lane:

- `targets/opencode/package.yaml` for `opencode.json.plugin`
- `targets/opencode/commands/` for `.opencode/commands/`
- `targets/opencode/agents/` for `.opencode/agents/`
- `targets/opencode/themes/` for `.opencode/themes/`
- `targets/opencode/plugins/` for `.opencode/plugins/`
- `targets/opencode/package.json` for `.opencode/package.json`
- portable `mcp/servers.json` for `opencode.json.mcp`
- portable `skills/` validated against the shared `SKILL.md` contract and mirrored into `.opencode/skills/`
- `targets/opencode/config.extra.json` for non-managed config passthrough
- native import compatibility for `opencode.json`, `opencode.jsonc`, project workspace directories, local plugin code/package metadata, and explicit `--include-user-scope`

Plugin specifics in this example:

- `targets/opencode/plugins/example.js` uses the canonical official-style named async plugin export and doubles as the loader smoke fixture
- `targets/opencode/plugins/custom-tool.js` shows beta custom-tool support through plugin code using `@opencode-ai/plugin`
- `targets/opencode/package.json` is the canonical authored source for plugin-local dependencies

Validate it with:

```bash
plugin-kit-ai render --check .
plugin-kit-ai validate . --platform opencode --strict
```
