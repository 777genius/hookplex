# hookplex SDK

Module: `github.com/hookplex/hookplex/sdk`

The SDK now exposes a platform-neutral runtime core with platform-specific public registrars.

Current contract status in this source tree: approved for `public-stable` in the pending `v1.0` release. Event-level support claims come from [../../docs/generated/support_matrix.md](../../docs/generated/support_matrix.md). Compatibility policy lives in [STABILITY.md](./STABILITY.md).

## Public API

Root package:

- `hookplex.New(hookplex.Config)`
- `(*hookplex.App).Use(...)`
- `(*hookplex.App).Claude()`
- `(*hookplex.App).Codex()`
- `(*hookplex.App).Run()`
- `(*hookplex.App).RunContext(ctx)`
- `hookplex.Supported()`

Platform packages:

- `github.com/hookplex/hookplex/sdk/claude`
- `github.com/hookplex/hookplex/sdk/codex`

## Supported Runtime Events

- `claude/Stop`
- `claude/PreToolUse`
- `claude/UserPromptSubmit`
- `codex/Notify`

Generated support matrix: [../../docs/generated/support_matrix.md](../../docs/generated/support_matrix.md)

## Generation

Runtime/scaffold/validate registries are generated from descriptor definitions.

```bash
go run ./cmd/hookplex-gen
```

## Claude Example

```go
package main

import (
	"os"

	hookplex "github.com/hookplex/hookplex/sdk"
	"github.com/hookplex/hookplex/sdk/claude"
)

func main() {
	app := hookplex.New(hookplex.Config{Name: "claude-demo"})
	app.Claude().OnStop(func(*claude.StopEvent) *claude.Response {
		return claude.Allow()
	})
	os.Exit(app.Run())
}
```

## Codex Example

```go
package main

import (
	"os"

	hookplex "github.com/hookplex/hookplex/sdk"
	"github.com/hookplex/hookplex/sdk/codex"
)

func main() {
	app := hookplex.New(hookplex.Config{Name: "codex-demo"})
	app.Codex().OnNotify(func(*codex.NotifyEvent) *codex.Response {
		return codex.Continue()
	})
	os.Exit(app.Run())
}
```
