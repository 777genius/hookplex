# plugin-kit-ai

AI CLI plugin runtime with a first-class Go SDK.

## Contract Status

This source tree contains the approved `v1.0` contract plus explicitly marked **public-experimental** surfaces.

Stable now:

- SDK root API and approved Claude/Codex event surfaces
- CLI commands `init`, `validate`, `capabilities`, `install`, `version`
- generated Go-first Claude/Codex scaffold contract

Beta now:

- optional scaffold extras from `plugin-kit-ai init --extras`
- executable runtime scaffolds for `python`, `node`, and `shell`
- experimental `plugin-kit-ai skills` authoring/rendering subsystem
- any future surfaces not explicitly promoted through the audit ledger

Canonical sources of truth:

- event support contract: [docs/generated/support_matrix.md](docs/generated/support_matrix.md)
- compatibility and public-surface policy: [docs/SUPPORT.md](docs/SUPPORT.md)
- delivery status ledger: [docs/STATUS.md](docs/STATUS.md)
- release lanes and shipping gates: [docs/RELEASE.md](docs/RELEASE.md)
- release notes template: [docs/RELEASE_NOTES_TEMPLATE.md](docs/RELEASE_NOTES_TEMPLATE.md)
- release rehearsal worksheet: [docs/REHEARSAL_TEMPLATE.md](docs/REHEARSAL_TEMPLATE.md)
- `v0.9` stable-candidate audit: [docs/V0_9_AUDIT.md](docs/V0_9_AUDIT.md)
- beta-breaking migration registry: [docs/MIGRATIONS.md](docs/MIGRATIONS.md)
- post-`v1` hardening mode: [docs/V1_0_X_HARDENING.md](docs/V1_0_X_HARDENING.md)
- diagnostics contract: [docs/DIAGNOSTICS.md](docs/DIAGNOSTICS.md)
- install compatibility contract: [docs/INSTALL_COMPATIBILITY.md](docs/INSTALL_COMPATIBILITY.md)
- executable plugin ABI: [docs/EXECUTABLE_ABI.md](docs/EXECUTABLE_ABI.md)
- threat model: [docs/THREAT_MODEL.md](docs/THREAT_MODEL.md)

Maintainer-only historical context:

- [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)
- [docs/FOUNDATION_REWRITE_VNEXT.md](docs/FOUNDATION_REWRITE_VNEXT.md)
- [docs/adr/README.md](docs/adr/README.md)

## Shipped Scope

What ships now:

- `sdk/plugin-kit-ai`: generated multi-platform runtime with peer public packages for Claude and Codex
- `cli/plugin-kit-ai`: `plugin-kit-ai init`, `plugin-kit-ai validate`, `plugin-kit-ai capabilities`, `plugin-kit-ai install`, `plugin-kit-ai version`
- `cli/plugin-kit-ai` plugin authoring flow: repo-root `plugin.yaml` plus `plugin-kit-ai render|import|normalize`
- `cli/plugin-kit-ai` experimental skills layer: `plugin-kit-ai skills init|validate|render`
- `install/plugininstall`: GitHub Releases installer with checksum verification

For the experimental skills layer, handwritten `skills/<name>/SKILL.md` is supported directly. `plugin-kit-ai skills init` is convenience scaffold, not a required authoring path.
For `plugin-kit-ai install`, the stable contract covers verified third-party plugin installation only. It does not promise self-update or an auto-update subsystem for the `plugin-kit-ai` CLI itself.
`plugin-kit-ai init` now keeps Go as the default runtime and can also scaffold executable plugins for `python`, `node`, and `shell`. These executable runtimes are repo-local and `public-beta`; install/update dependency management for interpreted runtimes remains out of scope.
New plugin projects use repo-root `plugin.yaml` as the canonical authoring manifest; native Claude/Codex/Gemini files are rendered artifacts, while `.plugin-kit-ai/project.toml` is retained only for legacy compatibility and migration. The supported `plugin.yaml` v1 surface is intentionally small, and `plugin-kit-ai validate` warns on unknown or deprecated manifest keys instead of silently treating them as supported.

Current runtime support:

- Claude stable: `Stop`, `PreToolUse`, `UserPromptSubmit`
- Claude beta: `SessionStart`, `SessionEnd`, `Notification`, `PostToolUse`, `PostToolUseFailure`, `PermissionRequest`, `SubagentStart`, `SubagentStop`, `PreCompact`, `Setup`, `TeammateIdle`, `TaskCompleted`, `ConfigChange`, `WorktreeCreate`, `WorktreeRemove`
- Codex: `Notify`

Release boundary notes:

- Codex stable support does not guarantee the health of the external `codex exec` runtime before hook execution.
- Claude stable support covers the declared event set only.
- Additional official Claude hooks may be runtime-supported in `public-beta` before they are promoted through the audit ledger.
- Experimental typed custom Claude hooks can be registered locally through `sdk/claude` generic helper functions when upstream support lags behind.
- Experimental typed custom Codex hooks can be registered locally through `sdk/codex` generic helper functions for future argv-JSON hook additions.

Current CLI scaffold targets:

- `--platform codex` (default)
- `--platform claude`
- `--runtime go` (default)
- `--runtime python`
- `--runtime node`
- `--runtime shell`

Generator-backed artifacts:

- runtime descriptor registry and invocation resolvers
- public platform registrars
- scaffold platform definitions
- validate rules
- capabilities registry
- generated support contract matrix

## Repository Layout

- `sdk/plugin-kit-ai`: SDK runtime, public platform packages, descriptor generator
- `cli/plugin-kit-ai`: CLI scaffold, validation, install wiring
- `install/plugininstall`: installer subsystem
- `repotests`: integration and guard tests
- `docs`: support policy, status ledger, release policy, generated contract docs

## Build And Test

Requirements:

- Go `1.22+`

Common commands from repo root:

```bash
go run ./cmd/plugin-kit-ai-gen
go build -o bin/plugin-kit-ai ./cli/plugin-kit-ai/cmd/plugin-kit-ai
./bin/plugin-kit-ai version
make test-polyglot-smoke

go test ./sdk/plugin-kit-ai/...
go test ./cli/plugin-kit-ai/...
go test ./install/plugininstall/...
go test ./repotests -run TestPluginKitAIInitGeneratesBuildableModule -count=1
go test ./...
```

## SDK

Root package `plugin-kit-ai` is now composition/runtime only. Platform APIs live in peer public packages:

- `github.com/plugin-kit-ai/plugin-kit-ai/sdk`
- `github.com/plugin-kit-ai/plugin-kit-ai/sdk/claude`
- `github.com/plugin-kit-ai/plugin-kit-ai/sdk/codex`

Claude example:

```go
package main

import (
	"os"

	pluginkitai "github.com/plugin-kit-ai/plugin-kit-ai/sdk"
	"github.com/plugin-kit-ai/plugin-kit-ai/sdk/claude"
)

func main() {
	app := pluginkitai.New(pluginkitai.Config{Name: "claude-demo"})
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response {
		return claude.Allow()
	})
	os.Exit(app.Run())
}
```

Codex example:

```go
package main

import (
	"os"

	pluginkitai "github.com/plugin-kit-ai/plugin-kit-ai/sdk"
	"github.com/plugin-kit-ai/plugin-kit-ai/sdk/codex"
)

func main() {
	app := pluginkitai.New(pluginkitai.Config{Name: "codex-demo"})
	app.Codex().OnNotify(func(*codex.NotifyEvent) *codex.Response {
		return codex.Continue()
	})
	os.Exit(app.Run())
}
```

See:

- [sdk/plugin-kit-ai/README.md](sdk/plugin-kit-ai/README.md)
- [docs/generated/support_matrix.md](docs/generated/support_matrix.md)
- [docs/SUPPORT.md](docs/SUPPORT.md)

## CLI

Build the CLI:

```bash
go build -o bin/plugin-kit-ai ./cli/plugin-kit-ai/cmd/plugin-kit-ai
```

Examples:

```bash
./bin/plugin-kit-ai init my-plugin
./bin/plugin-kit-ai init my-plugin --runtime python
./bin/plugin-kit-ai init my-plugin --platform claude --extras
./bin/plugin-kit-ai init my-plugin --platform claude --runtime shell
./bin/plugin-kit-ai render ./my-plugin
./bin/plugin-kit-ai render ./my-plugin --check
./bin/plugin-kit-ai import ./legacy-plugin --from codex
./bin/plugin-kit-ai normalize ./my-plugin
./bin/plugin-kit-ai validate ./my-plugin --platform codex
./bin/plugin-kit-ai validate ./my-plugin --platform codex --strict
./bin/plugin-kit-ai skills init lint-repo --template go-command
./bin/plugin-kit-ai skills validate .
./bin/plugin-kit-ai skills render . --target all
./bin/plugin-kit-ai capabilities --format json --platform claude
./bin/plugin-kit-ai install owner/repo --tag v1.0.0 --goos linux --goarch amd64
```

`plugin-kit-ai install` success output is intentionally compact but deterministic:

- installed file path
- resolved release ref and source (`--tag` or `--latest`)
- selected asset
- target GOOS/GOARCH
- overwrite notice only when `--force` replaced an existing file

The command verifies `checksums.txt` from the target release and installs third-party plugin binaries only. Self-update remains out of scope.
Supported and refused release layouts are documented in [docs/INSTALL_COMPATIBILITY.md](docs/INSTALL_COMPATIBILITY.md).

See:

- [cli/plugin-kit-ai/README.md](cli/plugin-kit-ai/README.md)
- [docs/EXECUTABLE_ABI.md](docs/EXECUTABLE_ABI.md)
- [docs/SKILLS.md](docs/SKILLS.md)
- [docs/RELEASE.md](docs/RELEASE.md)
