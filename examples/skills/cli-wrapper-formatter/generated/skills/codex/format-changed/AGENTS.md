## Skill: format-changed

- Description: Format changed files through an existing external formatter command.
- Canonical source: `skills/format-changed/SKILL.md`
- Command: `npx prettier@3.4.2 --write .`
- Runtime: `node`
- Allowed tools:
  - `bash`
  - `node`
- Compatibility:
  - Requires: node >=20
  - Supported OS: darwin, linux
  - Requires a repository checkout
  - May require network access
  - The first run may download the pinned package through npm.

Use the rendered skill instructions alongside this snippet when wiring repository guidance.
