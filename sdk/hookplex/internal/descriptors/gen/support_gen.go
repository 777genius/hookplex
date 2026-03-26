package gen

import "github.com/hookplex/hookplex/sdk/internal/runtime"

func AllSupportEntries() []runtime.SupportEntry {
	return []runtime.SupportEntry{
		{
			Platform:       "claude",
			Event:          "Stop",
			Status:         "runtime_supported",
			Maturity:       "beta",
			V1Target:       true,
			InvocationKind: "argv_command_casefold",
			Carrier:        runtime.CarrierStdinJSON,
			TransportModes: []runtime.TransportMode{
				"process",
			},
			ScaffoldSupport: true,
			ValidateSupport: true,
			Capabilities: []runtime.CapabilityID{
				"stop_gate",
			},
			Summary:         "Claude Stop command hook",
			LiveTestProfile: "claude_cli",
		},
		{
			Platform:       "claude",
			Event:          "PreToolUse",
			Status:         "runtime_supported",
			Maturity:       "beta",
			V1Target:       true,
			InvocationKind: "argv_command_casefold",
			Carrier:        runtime.CarrierStdinJSON,
			TransportModes: []runtime.TransportMode{
				"process",
			},
			ScaffoldSupport: true,
			ValidateSupport: true,
			Capabilities: []runtime.CapabilityID{
				"tool_gate",
			},
			Summary:         "Claude PreToolUse command hook",
			LiveTestProfile: "claude_cli",
		},
		{
			Platform:       "claude",
			Event:          "UserPromptSubmit",
			Status:         "runtime_supported",
			Maturity:       "beta",
			V1Target:       true,
			InvocationKind: "argv_command_casefold",
			Carrier:        runtime.CarrierStdinJSON,
			TransportModes: []runtime.TransportMode{
				"process",
			},
			ScaffoldSupport: true,
			ValidateSupport: true,
			Capabilities: []runtime.CapabilityID{
				"prompt_submit_gate",
			},
			Summary:         "Claude UserPromptSubmit command hook",
			LiveTestProfile: "claude_cli",
		},
		{
			Platform:       "codex",
			Event:          "Notify",
			Status:         "runtime_supported",
			Maturity:       "beta",
			V1Target:       true,
			InvocationKind: "argv_command",
			Carrier:        runtime.CarrierArgvJSON,
			TransportModes: []runtime.TransportMode{
				"process",
			},
			ScaffoldSupport: true,
			ValidateSupport: true,
			Capabilities: []runtime.CapabilityID{
				"notify",
			},
			Summary:         "Codex notify hook",
			LiveTestProfile: "codex_notify",
		},
	}
}
