# Target Support Matrix

| Target | Target Class | Production Class | Runtime Contract | Import | Render | Validate | Portable Components | Target-native Components | Managed Artifacts | Summary |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| claude | hook_runtime | production-ready | public-stable stable-subset runtime | yes | yes | yes | skills, mcp_servers, agents | package_metadata, hooks, commands, contexts | .claude-plugin/plugin.json, hooks/hooks.json, .mcp.json | Claude plugin packages compile portable skills and MCP plus target-native hook bindings. |
| codex | mixed_package_runtime | production-ready | public-stable notify runtime | yes | yes | yes | skills, mcp_servers | package_metadata, commands, contexts | .codex-plugin/plugin.json, .codex/config.toml, .mcp.json | Codex packages compile portable skills and MCP plus target metadata such as model hints. |
| gemini | mcp_extension | packaging-only target | not a production-ready runtime target | yes | yes | yes | skills, mcp_servers, agents | package_metadata, hooks, commands, policies, themes, settings, contexts | gemini-extension.json, commands/**, hooks/**, policies/**, contexts/** | Gemini currently compiles as an extension package with MCP, skills, and target-native extension assets. |
