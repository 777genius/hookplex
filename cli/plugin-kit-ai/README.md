# plugin-kit-ai CLI

Module: `github.com/plugin-kit-ai/plugin-kit-ai/cli`. Builds the **`plugin-kit-ai`** binary: `init`, `render`, `import`, `normalize`, `validate`, `capabilities`, `install`, `version`, plus experimental `skills` authoring commands.

Current CLI contract status in this source tree: `public-stable` shipped in `v1.0.0`, with additional post-`v1.0.x` hardening on `main`. Repository-wide compatibility and release policy live in [../../docs/SUPPORT.md](../../docs/SUPPORT.md) and [../../docs/RELEASE.md](../../docs/RELEASE.md).

`plugin-kit-ai init` scaffolds either a **Codex project** (`--platform codex`, default) or a **Claude plugin** (`--platform claude`) and supports `--runtime go|python|node|shell` with Go as the default runtime.
`plugin-kit-ai validate` checks either a legacy scaffold or a `plugin.yaml`-based plugin project, including generated-artifact drift and manifest warnings for unknown or deprecated `plugin.yaml` keys.
`plugin-kit-ai render` regenerates native vendor artifacts from `plugin.yaml`, `plugin-kit-ai import` backfills `plugin.yaml` from an existing native-only project, and `plugin-kit-ai normalize` rewrites `plugin.yaml` into the supported v1 shape.
`plugin-kit-ai capabilities` prints generated runtime support and capability metadata.

```bash
# from repository root
go build -o bin/plugin-kit-ai ./cli/plugin-kit-ai/cmd/plugin-kit-ai
```

Current-state behavior:

- `init`: project scaffold for `codex` or `claude`, with Go-first or executable runtimes
- `render`: generate native Claude/Codex runtime artifacts and Gemini packaging-only artifacts from `plugin.yaml`
- `import`: create `plugin.yaml` from an existing native plugin layout; Gemini import remains packaging-only
- `normalize`: rewrite `plugin.yaml` into the supported v1 schema and drop deprecated/unknown fields
- `validate`: legacy validation plus `plugin.yaml`-based project validation, generated-artifact drift checks, and non-failing manifest warnings; `--strict` promotes warnings to errors for CI
- `capabilities`: generated runtime support/contract introspection in table or JSON for Claude and Codex runtime events
- `install`: plugin binary from GitHub Releases with checksum verification
- `version`: build/version info
- `skills init|validate|render`: experimental SKILL.md authoring and agent render tooling

For the experimental skills subsystem, handwritten `skills/<name>/SKILL.md` is supported directly. `skills init` is convenience scaffold, not a required entrypoint.
For `install`, the stable CLI promise is limited to verified installation of third-party plugin binaries from GitHub Releases. It does not include self-update for the `plugin-kit-ai` CLI itself.
Executable runtime scaffolds for `python`, `node`, and `shell` are `public-beta`, repo-local, and do not add managed install/update handling for interpreted runtimes. `plugin.yaml` is the canonical authoring manifest for new projects and intentionally stays small in `v1`; unknown or deprecated manifest keys warn via `validate`, while `.plugin-kit-ai/project.toml` remains legacy compatibility only. Gemini is currently a `packaging-only target` in this CLI surface, not a production-ready runtime target.

Executable runtime matrix:

| Runtime | Status | Scope | Bootstrap |
|---------|--------|-------|-----------|
| `go` | stable | default production path | Go `1.22+` |
| `python` | public-beta | repo-local only | prefer `.venv`, fallback to system Python `3.10+` |
| `node` | public-beta | repo-local only | system Node.js `20+`; TypeScript via build-to-JS only |
| `shell` | public-beta | repo-local only | POSIX shell on Unix, `bash` in `PATH` on Windows |

For interpreted runtimes, `validate --strict` is the canonical CI-grade readiness gate.
`plugin-kit-ai install` remains binary-only and does not bootstrap or distribute Python/Node/Shell plugin dependencies.

Production-ready target boundary in the current contract:

- Claude: production-ready within the stable `Stop`, `PreToolUse`, and `UserPromptSubmit` set
- Codex: production-ready within the stable `Notify` path
- Gemini: packaging-only beta target through `render|import`

Canonical production plugin lane:

```bash
./bin/plugin-kit-ai normalize ./my-plugin
./bin/plugin-kit-ai render ./my-plugin
./bin/plugin-kit-ai render ./my-plugin --check
./bin/plugin-kit-ai validate ./my-plugin --platform codex --strict
```

`plugin-kit-ai install` prints a deterministic success summary:

- installed file path
- release ref with source (`tag` or `latest`)
- selected asset name
- target GOOS/GOARCH
- overwrite notice only when an existing file was replaced

Supported and unsupported release layouts for `install` are documented in [../../docs/INSTALL_COMPATIBILITY.md](../../docs/INSTALL_COMPATIBILITY.md).
Production authoring guidance and reference examples live in [../../docs/PRODUCTION.md](../../docs/PRODUCTION.md) and [../../examples/plugins/README.md](../../examples/plugins/README.md).

See the root [README.md](../../README.md) for current CLI behavior, shipped scope, and canonical support links.
See [../../docs/EXECUTABLE_ABI.md](../../docs/EXECUTABLE_ABI.md) for the low-level executable plugin contract.
See [../../docs/SKILLS.md](../../docs/SKILLS.md) for the skills workflow, positioning, and examples.

`go.mod` uses:

- `replace github.com/plugin-kit-ai/plugin-kit-ai/sdk => ../../sdk/plugin-kit-ai`
- `replace github.com/plugin-kit-ai/plugin-kit-ai/plugininstall => ../../install/plugininstall`
