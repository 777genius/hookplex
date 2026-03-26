## Skill: lint-repo

- Description: Run a fast repository lint pass and report actionable failures.
- Canonical source: `skills/lint-repo/SKILL.md`
- Command: `go run ./cmd/lint-repo`
- Runtime: `go`
- Allowed tools:
  - `bash`
  - `go`
- Compatibility:
  - Requires: go >=1.25
  - Supported OS: darwin, linux
  - Requires a repository checkout
  - Run from the repository root so the command can discover files consistently.

Use the rendered skill instructions alongside this snippet when wiring repository guidance.
