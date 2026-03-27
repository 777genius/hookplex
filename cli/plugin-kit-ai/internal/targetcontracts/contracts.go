package targetcontracts

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/tabwriter"
)

type Entry struct {
	Target                string   `json:"target"`
	TargetClass           string   `json:"target_class"`
	ProductionClass       string   `json:"production_class"`
	RuntimeContract       string   `json:"runtime_contract"`
	ImportSupport         bool     `json:"import_support"`
	RenderSupport         bool     `json:"render_support"`
	ValidateSupport       bool     `json:"validate_support"`
	PortableComponentKinds []string `json:"portable_component_kinds"`
	TargetComponentKinds  []string `json:"target_component_kinds"`
	ManagedArtifacts      []string `json:"managed_artifacts"`
	Summary               string   `json:"summary"`
}

func All() []Entry {
	return []Entry{
		{
			Target:                 "claude",
			TargetClass:            "hook_runtime",
			ProductionClass:        "production-ready",
			RuntimeContract:        "public-stable stable-subset runtime",
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
			ProductionClass:        "production-ready",
			RuntimeContract:        "public-stable notify runtime",
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
			ProductionClass:        "packaging-only target",
			RuntimeContract:        "not a production-ready runtime target",
			ImportSupport:          true,
			RenderSupport:          true,
			ValidateSupport:        true,
			PortableComponentKinds: []string{"skills", "mcp_servers", "agents"},
			TargetComponentKinds:   []string{"package_metadata", "hooks", "commands", "policies", "themes", "settings", "contexts"},
			ManagedArtifacts:       []string{"gemini-extension.json", "commands/**", "hooks/**", "policies/**", "contexts/**"},
			Summary:                "Gemini currently compiles as an extension package with MCP, skills, and target-native extension assets.",
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
	_, _ = w.Write([]byte("TARGET\tCLASS\tPRODUCTION\tRUNTIME\tIMPORT\tRENDER\tVALIDATE\tPORTABLE\tTARGET-NATIVE\tMANAGED\tSUMMARY\n"))
	for _, entry := range entries {
		_, _ = w.Write([]byte(
			entry.Target + "\t" +
				entry.TargetClass + "\t" +
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
	b.WriteString("| Target | Target Class | Production Class | Runtime Contract | Import | Render | Validate | Portable Components | Target-native Components | Managed Artifacts | Summary |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |\n")
	for _, entry := range entries {
		b.WriteString("| " + entry.Target + " | " + entry.TargetClass + " | " + entry.ProductionClass + " | " + entry.RuntimeContract + " | " + yesNo(entry.ImportSupport) + " | " + yesNo(entry.RenderSupport) + " | " + yesNo(entry.ValidateSupport) + " | " + join(entry.PortableComponentKinds) + " | " + join(entry.TargetComponentKinds) + " | " + join(entry.ManagedArtifacts) + " | " + entry.Summary + " |\n")
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
