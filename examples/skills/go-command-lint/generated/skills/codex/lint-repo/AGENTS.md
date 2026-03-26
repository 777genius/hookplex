## Skill: lint-repo

- Description: Check that this skill package is internally consistent and report actionable failures.
- Canonical source: `skills/lint-repo/SKILL.md`
- Command: `go run ./cmd/lint-repo`
- Runtime: `go`
- Allowed tools:
  - `bash`
  - `go`
- Compatibility:
  - Requires: go >=1.22
  - Supported OS: darwin, linux
  - Requires a repository checkout
  - Run from the example root so the command can inspect the authored and generated skill files together.

Use the rendered skill instructions alongside this snippet when wiring repository guidance.
