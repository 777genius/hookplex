# opencode-basic

Reference OpenCode workspace-config example for `plugin-kit-ai`.

This example demonstrates the current OpenCode v1 lane:

- `targets/opencode/package.yaml` for `opencode.json.plugin`
- portable `mcp/servers.json` for `opencode.json.mcp`
- portable `skills/` mirrored into `.opencode/skills/`
- `targets/opencode/config.extra.json` for non-managed config passthrough

Validate it with:

```bash
plugin-kit-ai render --check .
plugin-kit-ai validate . --platform opencode --strict
```
