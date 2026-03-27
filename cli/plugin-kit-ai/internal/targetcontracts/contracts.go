package targetcontracts

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/tabwriter"
)

type Entry struct {
	Target                 string   `json:"target"`
	TargetClass            string   `json:"target_class"`
	TargetNoun             string   `json:"target_noun,omitempty"`
	ProductionClass        string   `json:"production_class"`
	RuntimeContract        string   `json:"runtime_contract"`
	InstallModel           string   `json:"install_model,omitempty"`
	DevModel               string   `json:"dev_model,omitempty"`
	ActivationModel        string   `json:"activation_model,omitempty"`
	NativeRoot             string   `json:"native_root,omitempty"`
	ImportSupport          bool     `json:"import_support"`
	RenderSupport          bool     `json:"render_support"`
	ValidateSupport        bool     `json:"validate_support"`
	PortableComponentKinds []string `json:"portable_component_kinds"`
	TargetComponentKinds   []string `json:"target_component_kinds"`
	ManagedArtifacts       []string `json:"managed_artifacts"`
	Summary                string   `json:"summary"`
}

func All() []Entry {
	return []Entry{
		{
			Target:                 "claude",
			TargetClass:            "hook_runtime",
			TargetNoun:             "plugin",
			ProductionClass:        "production-ready",
			RuntimeContract:        "public-stable stable-subset runtime",
			InstallModel:           "marketplace or local plugin install",
			DevModel:               "reload plugins",
			ActivationModel:        "reload or restart",
			NativeRoot:             "~/.claude/plugins/...",
			ImportSupport:          true,
			RenderSupport:          true,
			ValidateSupport:        true,
			PortableComponentKinds: []string{"skills", "mcp_servers", "agents"},
			TargetComponentKinds:   []string{"package_metadata", "hooks", "commands", "contexts"},
			ManagedArtifacts:       []string{".claude-plugin/plugin.json", "hooks/hooks.json", ".mcp.json"},
			Summary:                "Claude plugin packages compile portable skills and MCP plus target-native hook bindings.",
		},
		{
			Target:                 "codex",
			TargetClass:            "mixed_package_runtime",
			TargetNoun:             "plugin",
			ProductionClass:        "production-ready",
			RuntimeContract:        "public-stable notify runtime",
			InstallModel:           "plugin directory or marketplace cache",
			DevModel:               "local plugin workspace",
			ActivationModel:        "config reload or restart",
			NativeRoot:             "~/.codex/plugins/...",
			ImportSupport:          true,
			RenderSupport:          true,
			ValidateSupport:        true,
			PortableComponentKinds: []string{"skills", "mcp_servers"},
			TargetComponentKinds:   []string{"package_metadata", "commands", "contexts"},
			ManagedArtifacts:       []string{".codex-plugin/plugin.json", ".codex/config.toml", ".mcp.json"},
			Summary:                "Codex packages compile portable skills and MCP plus target metadata such as model hints.",
		},
		{
			Target:                 "gemini",
			TargetClass:            "mcp_extension",
			TargetNoun:             "extension",
			ProductionClass:        "packaging-only target",
			RuntimeContract:        "not a production-ready runtime target",
			InstallModel:           "copy install",
			DevModel:               "link",
			ActivationModel:        "restart required",
			NativeRoot:             "~/.gemini/extensions/<name>",
			ImportSupport:          true,
			RenderSupport:          true,
			ValidateSupport:        true,
			PortableComponentKinds: []string{"skills", "mcp_servers", "agents"},
			TargetComponentKinds:   []string{"package_metadata", "hooks", "commands", "policies", "themes", "settings", "contexts", "manifest_extra"},
			ManagedArtifacts:       []string{"gemini-extension.json", "commands/**", "hooks/**", "policies/**", "GEMINI.md or selected root context", "contexts/**"},
			Summary:                "Gemini compiles as an official-style extension package with MCP, a primary root context, and target-native extension assets.",
		},
	}
}

func ByTarget(name string) []Entry {
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		return All()
	}
	out := make([]Entry, 0, 1)
	for _, entry := range All() {
		if entry.Target == name {
			out = append(out, entry)
		}
	}
	return out
}

func Lookup(name string) (Entry, bool) {
	for _, entry := range All() {
		if entry.Target == strings.ToLower(strings.TrimSpace(name)) {
			return entry, true
		}
	}
	return Entry{}, false
}

func JSON(entries []Entry) ([]byte, error) {
	return json.MarshalIndent(entries, "", "  ")
}

func Table(entries []Entry) []byte {
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)
	_, _ = w.Write([]byte("TARGET\tCLASS\tNOUN\tINSTALL\tDEV\tACTIVATION\tNATIVE ROOT\tPRODUCTION\tRUNTIME\tIMPORT\tRENDER\tVALIDATE\tPORTABLE\tTARGET-NATIVE\tMANAGED\tSUMMARY\n"))
	for _, entry := range entries {
		_, _ = w.Write([]byte(
			entry.Target + "\t" +
				entry.TargetClass + "\t" +
				entry.TargetNoun + "\t" +
				entry.InstallModel + "\t" +
				entry.DevModel + "\t" +
				entry.ActivationModel + "\t" +
				entry.NativeRoot + "\t" +
				entry.ProductionClass + "\t" +
				entry.RuntimeContract + "\t" +
				yesNo(entry.ImportSupport) + "\t" +
				yesNo(entry.RenderSupport) + "\t" +
				yesNo(entry.ValidateSupport) + "\t" +
				join(entry.PortableComponentKinds) + "\t" +
				join(entry.TargetComponentKinds) + "\t" +
				join(entry.ManagedArtifacts) + "\t" +
				entry.Summary + "\n"))
	}
	_ = w.Flush()
	return buf.Bytes()
}

func Markdown(entries []Entry) []byte {
	var b bytes.Buffer
	b.WriteString("# Target Support Matrix\n\n")
	b.WriteString("| Target | Target Class | Target Noun | Install Model | Dev Model | Activation Model | Native Root | Production Class | Runtime Contract | Import | Render | Validate | Portable Components | Target-native Components | Managed Artifacts | Summary |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |\n")
	for _, entry := range entries {
		b.WriteString("| " + entry.Target + " | " + entry.TargetClass + " | " + entry.TargetNoun + " | " + entry.InstallModel + " | " + entry.DevModel + " | " + entry.ActivationModel + " | " + entry.NativeRoot + " | " + entry.ProductionClass + " | " + entry.RuntimeContract + " | " + yesNo(entry.ImportSupport) + " | " + yesNo(entry.RenderSupport) + " | " + yesNo(entry.ValidateSupport) + " | " + join(entry.PortableComponentKinds) + " | " + join(entry.TargetComponentKinds) + " | " + join(entry.ManagedArtifacts) + " | " + entry.Summary + " |\n")
	}
	return b.Bytes()
}

func join(items []string) string {
	if len(items) == 0 {
		return "-"
	}
	return strings.Join(items, ", ")
}

func yesNo(v bool) string {
	if v {
		return "yes"
	}
	return "no"
}
