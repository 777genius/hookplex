package gen

import (
	"fmt"
	"github.com/hookplex/hookplex/sdk/internal/runtime"
	"strings"
)

func ResolveInvocation(args []string, _ runtime.Env) (runtime.Invocation, error) {
	if len(args) < 2 {
		return runtime.Invocation{}, fmt.Errorf("usage: <binary> <hookName>")
	}
	raw := args[1]
	if strings.EqualFold(raw, "Stop") {
		return runtime.Invocation{Platform: "claude", Event: "Stop", RawName: raw}, nil
	}
	if strings.EqualFold(raw, "PreToolUse") {
		return runtime.Invocation{Platform: "claude", Event: "PreToolUse", RawName: raw}, nil
	}
	if strings.EqualFold(raw, "UserPromptSubmit") {
		return runtime.Invocation{Platform: "claude", Event: "UserPromptSubmit", RawName: raw}, nil
	}
	if raw == "notify" {
		return runtime.Invocation{Platform: "codex", Event: "Notify", RawName: raw}, nil
	}
	return runtime.Invocation{}, fmt.Errorf("unknown invocation %q", raw)
}
