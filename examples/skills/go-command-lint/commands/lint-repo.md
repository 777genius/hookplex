# /lint-repo

Generated command reference for `lint-repo`.

## Purpose

Run a fast repository lint pass and report actionable failures.

## Invocation

`go run ./cmd/lint-repo`
## Runtime

`go`
## Allowed tools
- `bash`
- `go`
## Compatibility
- Requires: go >=1.25
- Supported OS: darwin, linux
- Requires a repository checkout
- Run from the repository root so the command can discover files consistently.

## Notes

- Safe to retry: yes
- Writes files: no
- Produces JSON: no
- This file is generated from `skills/lint-repo/SKILL.md`.
- Regenerate with `hookplex skills render`.
