# Architecture Notes

This document describes the current shipped monorepo architecture.

Current public contract docs live in:

- [SUPPORT.md](./SUPPORT.md)
- [STATUS.md](./STATUS.md)
- [generated/support_matrix.md](./generated/support_matrix.md)

Historical maintainer references live in:

- [FOUNDATION_REWRITE_VNEXT.md](./FOUNDATION_REWRITE_VNEXT.md)
- [adr/README.md](./adr/README.md)

## Composition Roots

| Layer | Location | Role |
|-------|----------|------|
| SDK runtime | `sdk/hookplex/hookplex.go` | Platform-neutral composition root that wires the generic engine, generated descriptor lookup, middleware, and platform registrars |
| SDK generator | `cmd/hookplex-gen/main.go` | Generates descriptor-derived runtime, scaffold, validate, and docs artifacts |
| Plugin install library | `install/plugininstall/install.go` | Public install facade that wires use case and concrete adapters |
| CLI | `cli/hookplex/cmd/hookplex/main.go` | Process entrypoint; commands parse flags and call `internal/app`, `internal/scaffold`, and `internal/validate` |

Rule: the CLI must not construct `plugininstall` adapters directly. It uses the `plugininstall` facade.

## SDK Runtime

- `sdk/hookplex` exposes only shared runtime composition.
- Public platform APIs are peer packages:
  - `sdk/hookplex/claude`
  - `sdk/hookplex/codex`
- Core runtime lives under `sdk/hookplex/internal/runtime`.
- Descriptor definitions live under `sdk/hookplex/internal/descriptors/defs`.
- Generated runtime registries and resolvers live under `sdk/hookplex/internal/descriptors/gen`.
- Platform wire codecs live under:
  - `sdk/hookplex/internal/platforms/claude`
  - `sdk/hookplex/internal/platforms/codex`

Current runtime carriers:

- Claude events use `stdin_json`
- Codex `Notify` uses `argv_json`

## CLI Application Layer

`cli/hookplex/internal/app` keeps Cobra out of install/init application logic:

- `InstallRunner` delegates to `plugininstall.Install`
- `InitRunner` resolves generated scaffold definitions and delegates rendering to `scaffold`

`cli/hookplex/internal/validate` enforces generated platform rules for scaffolded projects.

## Generated Sources

`go run ./cmd/hookplex-gen` is the canonical generation entrypoint.

Generated artifacts include:

- descriptor registry and invocation resolvers
- public platform registrars
- scaffold platform definitions
- validation rules
- support contract documentation

Generator drift is enforced by tests in `sdk/hookplex/generator`.

## Exit Codes

- `hookplex install`: domain errors map through `plugininstall.ExitCodeFromErr` and CLI `exitx`
- `hookplex init`: failures surface as CLI errors and exit code `1`
- `hookplex validate`: invalid scaffold or buildability failures exit non-zero

## Tests

- `sdk/hookplex/...`: runtime, descriptors, generator drift, examples
- `cli/hookplex/...`: app and scaffold coverage
- `repotests/...`: generated project integration and installer integration

Note: installer integration tests create a local `httptest` server and require loopback bind permissions.
