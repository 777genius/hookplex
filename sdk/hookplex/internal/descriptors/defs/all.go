package defs

import "github.com/hookplex/hookplex/sdk/internal/runtime"

func Profiles() []PlatformProfile {
	return []PlatformProfile{
		{
			Platform:        "claude",
			Status:          runtime.StatusRuntimeSupported,
			PublicPackage:   "claude",
			InternalPackage: "claude",
			InternalImport:  "github.com/hookplex/hookplex/sdk/internal/platforms/claude",
			TransportModes:  []runtime.TransportMode{runtime.ProcessMode},
			LiveTestProfile: "claude_cli",
			Scaffold: ScaffoldMeta{
				RequiredFiles: []string{
					"go.mod",
					"README.md",
					".claude-plugin/plugin.json",
					"hooks/hooks.json",
				},
				OptionalFiles: []string{
					"Makefile",
					".goreleaser.yml",
					"skills/{{.ProjectName}}/SKILL.md",
					"commands/{{.ProjectName}}.md",
				},
				ForbiddenFiles: []string{
					"AGENTS.md",
					".codex/config.toml",
				},
				TemplateFiles: []TemplateFile{
					{Path: "go.mod", Template: "go.mod.tmpl"},
					{Path: "cmd/{{.ProjectName}}/main.go", Template: "main.go.tmpl"},
					{Path: ".claude-plugin/plugin.json", Template: "plugin.json.tmpl"},
					{Path: "hooks/hooks.json", Template: "hooks.json.tmpl"},
					{Path: "README.md", Template: "README.md.tmpl"},
					{Path: "Makefile", Template: "Makefile.tmpl", Extra: true},
					{Path: ".goreleaser.yml", Template: "goreleaser.yml.tmpl", Extra: true},
					{Path: "skills/{{.ProjectName}}/SKILL.md", Template: "SKILL.md.tmpl", Extra: true},
					{Path: "commands/{{.ProjectName}}.md", Template: "command.md.tmpl", Extra: true},
				},
			},
			Validate: ValidateMeta{
				RequiredFiles: []string{
					"go.mod",
					"README.md",
					".claude-plugin/plugin.json",
					"hooks/hooks.json",
				},
				ForbiddenFiles: []string{
					"AGENTS.md",
					".codex/config.toml",
				},
				BuildTargets: []string{"./..."},
			},
		},
		{
			Platform:        "codex",
			Status:          runtime.StatusRuntimeSupported,
			PublicPackage:   "codex",
			InternalPackage: "codex",
			InternalImport:  "github.com/hookplex/hookplex/sdk/internal/platforms/codex",
			TransportModes:  []runtime.TransportMode{runtime.ProcessMode},
			LiveTestProfile: "codex_notify",
			Scaffold: ScaffoldMeta{
				RequiredFiles: []string{
					"go.mod",
					"README.md",
					"AGENTS.md",
					".codex/config.toml",
				},
				OptionalFiles: []string{
					"Makefile",
					".goreleaser.yml",
					"skills/{{.ProjectName}}/SKILL.md",
					"commands/{{.ProjectName}}.md",
				},
				ForbiddenFiles: []string{
					".claude-plugin/plugin.json",
					"hooks/hooks.json",
				},
				TemplateFiles: []TemplateFile{
					{Path: "go.mod", Template: "codex.go.mod.tmpl"},
					{Path: "cmd/{{.ProjectName}}/main.go", Template: "codex.main.go.tmpl"},
					{Path: "AGENTS.md", Template: "codex.AGENTS.md.tmpl"},
					{Path: ".codex/config.toml", Template: "codex.config.toml.tmpl"},
					{Path: "README.md", Template: "codex.README.md.tmpl"},
					{Path: "Makefile", Template: "Makefile.tmpl", Extra: true},
					{Path: ".goreleaser.yml", Template: "goreleaser.yml.tmpl", Extra: true},
					{Path: "skills/{{.ProjectName}}/SKILL.md", Template: "SKILL.md.tmpl", Extra: true},
					{Path: "commands/{{.ProjectName}}.md", Template: "command.md.tmpl", Extra: true},
				},
			},
			Validate: ValidateMeta{
				RequiredFiles: []string{
					"go.mod",
					"README.md",
					"AGENTS.md",
					".codex/config.toml",
				},
				ForbiddenFiles: []string{
					".claude-plugin/plugin.json",
					"hooks/hooks.json",
				},
				BuildTargets: []string{"./..."},
			},
		},
	}
}

func Events() []EventDescriptor {
	return []EventDescriptor{
		{
			Platform: "claude",
			Event:    "Stop",
			Invocation: InvocationBinding{
				Kind: runtime.InvocationArgvCommandCaseFold,
				Name: "Stop",
			},
			Carrier: runtime.CarrierStdinJSON,
			Contract: ContractMeta{
				Maturity: runtime.MaturityBeta,
				V1Target: true,
			},
			DecodeFunc: "DecodeStop",
			EncodeFunc: "EncodeStop",
			Registrar: RegistrarMeta{
				MethodName:   "OnStop",
				EventType:    "*StopEvent",
				ResponseType: "*Response",
				WrapFunc:     "wrapStop",
			},
			Docs: DocsMeta{
				SnippetKey: "claude-stop",
				TableGroup: "claude",
				Summary:    "Claude Stop command hook",
			},
			Capabilities: []runtime.CapabilityID{"stop_gate"},
			CapabilityMappings: []CapabilityMapping{
				{Unified: "stop_gate", Platform: "stop_gate"},
			},
		},
		{
			Platform: "claude",
			Event:    "PreToolUse",
			Invocation: InvocationBinding{
				Kind: runtime.InvocationArgvCommandCaseFold,
				Name: "PreToolUse",
			},
			Carrier: runtime.CarrierStdinJSON,
			Contract: ContractMeta{
				Maturity: runtime.MaturityBeta,
				V1Target: true,
			},
			DecodeFunc: "DecodePreToolUse",
			EncodeFunc: "EncodePreToolUse",
			Registrar: RegistrarMeta{
				MethodName:   "OnPreToolUse",
				EventType:    "*PreToolUseEvent",
				ResponseType: "*PreToolResponse",
				WrapFunc:     "wrapPreToolUse",
			},
			Docs: DocsMeta{
				SnippetKey: "claude-pretooluse",
				TableGroup: "claude",
				Summary:    "Claude PreToolUse command hook",
			},
			Capabilities: []runtime.CapabilityID{"tool_gate"},
			CapabilityMappings: []CapabilityMapping{
				{Unified: "tool_gate", Platform: "tool_gate"},
			},
		},
		{
			Platform: "claude",
			Event:    "UserPromptSubmit",
			Invocation: InvocationBinding{
				Kind: runtime.InvocationArgvCommandCaseFold,
				Name: "UserPromptSubmit",
			},
			Carrier: runtime.CarrierStdinJSON,
			Contract: ContractMeta{
				Maturity: runtime.MaturityBeta,
				V1Target: true,
			},
			DecodeFunc: "DecodeUserPromptSubmit",
			EncodeFunc: "EncodeUserPromptSubmit",
			Registrar: RegistrarMeta{
				MethodName:   "OnUserPromptSubmit",
				EventType:    "*UserPromptEvent",
				ResponseType: "*UserPromptResponse",
				WrapFunc:     "wrapUserPromptSubmit",
			},
			Docs: DocsMeta{
				SnippetKey: "claude-userpromptsubmit",
				TableGroup: "claude",
				Summary:    "Claude UserPromptSubmit command hook",
			},
			Capabilities: []runtime.CapabilityID{"prompt_submit_gate"},
			CapabilityMappings: []CapabilityMapping{
				{Unified: "prompt_submit_gate", Platform: "prompt_submit_gate"},
			},
		},
		{
			Platform: "codex",
			Event:    "Notify",
			Invocation: InvocationBinding{
				Kind: runtime.InvocationArgvCommand,
				Name: "notify",
			},
			Carrier: runtime.CarrierArgvJSON,
			Contract: ContractMeta{
				Maturity: runtime.MaturityBeta,
				V1Target: true,
			},
			DecodeFunc: "DecodeNotify",
			EncodeFunc: "EncodeNotify",
			Registrar: RegistrarMeta{
				MethodName:   "OnNotify",
				EventType:    "*NotifyEvent",
				ResponseType: "*Response",
				WrapFunc:     "wrapNotify",
			},
			Docs: DocsMeta{
				SnippetKey: "codex-notify",
				TableGroup: "codex",
				Summary:    "Codex notify hook",
			},
			Capabilities: []runtime.CapabilityID{"notify"},
		},
	}
}
