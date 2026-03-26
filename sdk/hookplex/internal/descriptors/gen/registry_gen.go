package gen

import (
	internal_claude "github.com/hookplex/hookplex/sdk/internal/platforms/claude"
	internal_codex "github.com/hookplex/hookplex/sdk/internal/platforms/codex"
	"github.com/hookplex/hookplex/sdk/internal/runtime"
)

type key struct {
	platform runtime.PlatformID
	event    runtime.EventID
}

var registry = map[key]runtime.Descriptor{
	{platform: "claude", event: "Stop"}: {
		Platform: "claude",
		Event:    "Stop",
		Carrier:  runtime.CarrierStdinJSON,
		Decode:   internal_claude.DecodeStop,
		Encode:   internal_claude.EncodeStop,
	},
	{platform: "claude", event: "PreToolUse"}: {
		Platform: "claude",
		Event:    "PreToolUse",
		Carrier:  runtime.CarrierStdinJSON,
		Decode:   internal_claude.DecodePreToolUse,
		Encode:   internal_claude.EncodePreToolUse,
	},
	{platform: "claude", event: "UserPromptSubmit"}: {
		Platform: "claude",
		Event:    "UserPromptSubmit",
		Carrier:  runtime.CarrierStdinJSON,
		Decode:   internal_claude.DecodeUserPromptSubmit,
		Encode:   internal_claude.EncodeUserPromptSubmit,
	},
	{platform: "codex", event: "Notify"}: {
		Platform: "codex",
		Event:    "Notify",
		Carrier:  runtime.CarrierArgvJSON,
		Decode:   internal_codex.DecodeNotify,
		Encode:   internal_codex.EncodeNotify,
	},
}

func Lookup(platform runtime.PlatformID, event runtime.EventID) (runtime.Descriptor, bool) {
	d, ok := registry[key{platform: platform, event: event}]
	return d, ok
}
