# hookplex Skills

`hookplex skills` is an experimental, compatibility-first authoring layer for repository-local skills.

The goal is not to replace the broader `SKILL.md` ecosystem. The goal is to make the common workflow easier to maintain when you want:

- one canonical authored `SKILL.md`
- validation before you publish or share a skill
- rendered artifacts for both Claude and Codex from the same source
- a strong executable path for Go without forcing Go on everyone

## What It Is

`hookplex skills` treats a skill as:

- authored instructions in `skills/<name>/SKILL.md`
- YAML frontmatter for light validation and rendering metadata
- optional executable behavior through a non-interactive command contract
- optional `scripts/`, `references/`, and `assets/`

The canonical authored file stays:

- `skills/<name>/SKILL.md`

Generated files are derived outputs. Edit the canonical skill, then re-render.

## How It Differs From Hooks

Skills and hooks solve different problems.

- Hooks are deterministic runtime integrations triggered by Claude or Codex lifecycle events.
- Skills are authored instruction packages that may optionally point to an executable command.

Use hooks when you need lifecycle automation.
Use skills when you need reusable guidance, workflows, or command-backed procedures.

## Why Use This Instead Of Writing SKILL.md By Hand

If all you need is a tiny `SKILL.md` and a trivial one-off command, writing it by hand is fine.

`hookplex skills` becomes useful when you want:

- a canonical authoring workflow: `init -> edit -> validate -> render`
- clearer validation than "hope the frontmatter is right"
- one authored source for both Claude and Codex render targets
- an executable path that is not biased toward shell scripts

## Canonical Workflow

```bash
hookplex skills init lint-repo --template go-command
# edit skills/lint-repo/SKILL.md
hookplex skills validate .
hookplex skills render . --target all
```

Templates:

- `go-command`: best default for typed, testable executable skills
- `cli-wrapper`: for an existing Python, Node, shell, or external CLI workflow
- `docs-only`: for instructional skills with no executable step

## Execution Model

Supported execution modes:

- `docs_only`
- `command`

`command` is language-neutral. It can describe:

- a Go binary or Go command
- a shell script
- Python, Node, or Deno commands
- `npx`, `uvx`, `go run`, Docker wrappers, or another external CLI

`hookplex` does not execute commands during validation. It only validates the authored contract statically.

## When Not To Use hookplex Skills

You probably do not need this subsystem when:

- your skill is a tiny standalone `SKILL.md`
- you do not need validation
- you do not care about rendering for multiple agents
- you are already happy with handwritten artifacts for one agent only

## Examples

See:

- [examples/skills/README.md](../examples/skills/README.md)
- [examples/skills/go-command-lint](../examples/skills/go-command-lint/README.md)
- [examples/skills/cli-wrapper-formatter](../examples/skills/cli-wrapper-formatter/README.md)
- [examples/skills/docs-only-review](../examples/skills/docs-only-review/README.md)

## Stability

This subsystem is `public-experimental`.

That means:

- `SKILL.md` remains the canonical authored contract
- the current workflow is intended for real use
- validation and renderer output may still evolve
- this subsystem is not part of the stable `v1.0` runtime compatibility promise
