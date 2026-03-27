# gemini-extension-package

Reference Gemini CLI extension repo for the current `plugin-kit-ai` packaging workflow.

Packaging contract:

- official-style `gemini-extension.json`
- shared MCP from `mcp/servers.json`
- one primary root context file
- native Gemini commands and policies
- manifest-driven settings and themes

This example is intentionally `packaging-only`. It does not claim Gemini runtime parity with Claude or Codex.

## Workflow

```bash
plugin-kit-ai normalize .
plugin-kit-ai render .
plugin-kit-ai render --check .
plugin-kit-ai validate . --platform gemini --strict
go test ./...
go build -o bin/gemini-extension-package ./cmd/gemini-extension-package
```
