package validate

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/pluginmanifest"
	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/targetcontracts"
	"gopkg.in/yaml.v3"
)

type FailureKind string

const (
	FailureUnknownPlatform          FailureKind = "unknown_platform"
	FailureCannotInferPlatform      FailureKind = "cannot_infer_platform"
	FailureManifestMissing          FailureKind = "manifest_missing"
	FailureManifestInvalid          FailureKind = "manifest_invalid"
	FailureRequiredFileMissing      FailureKind = "required_file_missing"
	FailureForbiddenFilePresent     FailureKind = "forbidden_file_present"
	FailureBuildFailed              FailureKind = "build_failed"
	FailureRuntimeNotFound          FailureKind = "runtime_not_found"
	FailureEntrypointMismatch       FailureKind = "entrypoint_mismatch"
	FailureLauncherInvalid          FailureKind = "launcher_invalid"
	FailureRuntimeTargetMissing     FailureKind = "runtime_target_missing"
	FailureGeneratedContractInvalid FailureKind = "generated_contract_invalid"
	FailureSourceFileMissing        FailureKind = "source_file_missing"
	FailureUnsupportedTargetKind    FailureKind = "unsupported_target_kind"
)

type Failure struct {
	Kind    FailureKind
	Path    string
	Target  string
	Message string
}

type WarningKind string

const (
	WarningManifestUnknownField  WarningKind = "manifest_unknown_field"
	WarningGeminiDirNameMismatch WarningKind = "gemini_dir_name_mismatch"
	WarningGeminiMCPCommandStyle WarningKind = "gemini_mcp_command_style"
	WarningGeminiPolicyIgnored   WarningKind = "gemini_policy_ignored"
)

type Warning struct {
	Kind    WarningKind
	Path    string
	Message string
}

type Report struct {
	Platform string
	Checks   []string
	Warnings []Warning
	Failures []Failure
}

type ReportError struct {
	Report Report
}

func (e *ReportError) Error() string {
	if len(e.Report.Failures) == 0 {
		return "validation failed"
	}
	f := e.Report.Failures[0]
	switch f.Kind {
	case FailureRequiredFileMissing:
		return "required file missing: " + f.Path
	case FailureForbiddenFilePresent:
		return fmt.Sprintf("forbidden file present for platform %s: %s", e.Report.Platform, f.Path)
	case FailureBuildFailed:
		return fmt.Sprintf("go build %s: %s", f.Target, f.Message)
	case FailureManifestMissing, FailureManifestInvalid, FailureRuntimeNotFound, FailureEntrypointMismatch, FailureLauncherInvalid, FailureRuntimeTargetMissing:
		return f.Message
	default:
		return f.Message
	}
}

type Rule struct {
	Platform       string
	RequiredFiles  []string
	ForbiddenFiles []string
	BuildTargets   []string
}

func Run(root, platform string) error {
	report, err := Validate(root, platform)
	if err != nil {
		return err
	}
	if len(report.Failures) > 0 {
		return &ReportError{Report: report}
	}
	return nil
}

func Validate(root, platform string) (Report, error) {
	if fileExists(filepath.Join(root, pluginmanifest.FileName)) {
		return validatePluginProject(root, platform)
	}
	if fileExists(filepath.Join(root, ".plugin-kit-ai", "project.toml")) {
		return Report{}, &ReportError{Report: Report{
			Failures: []Failure{{
				Kind:    FailureManifestInvalid,
				Message: "unsupported project format: .plugin-kit-ai/project.toml is not supported; use plugin.yaml and targets/<platform>/...",
			}},
		}}
	}
	return Report{}, &ReportError{Report: Report{
		Failures: []Failure{{
			Kind:    FailureManifestMissing,
			Message: "required manifest missing: plugin.yaml",
		}},
	}}
}

func validatePluginProject(root, platform string) (Report, error) {
	manifest, warnings, err := pluginmanifest.LoadWithWarnings(root)
	if err != nil {
		return Report{}, &ReportError{Report: Report{
			Failures: []Failure{{
				Kind:    FailureManifestInvalid,
				Message: err.Error(),
			}},
		}}
	}

	report := Report{
		Platform: strings.Join(manifest.EnabledTargets(), ","),
		Checks:   []string{"plugin_manifest", "package_graph", "generated_artifacts", "runtime"},
	}
	if strings.TrimSpace(platform) != "" && !slices.Contains(manifest.EnabledTargets(), strings.TrimSpace(platform)) {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureManifestInvalid,
			Message: fmt.Sprintf("plugin.yaml does not enable target %q", platform),
		})
	}
	for _, warning := range warnings {
		report.Warnings = append(report.Warnings, Warning{
			Kind:    mapManifestWarningKind(warning.Kind),
			Path:    warning.Path,
			Message: warning.Message,
		})
	}
	graph, _, err := pluginmanifest.Discover(root)
	if err != nil {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureManifestInvalid,
			Message: err.Error(),
		})
		return report, nil
	}
	for _, rel := range graph.SourceFiles {
		if _, err := os.Stat(filepath.Join(root, rel)); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureSourceFileMissing,
				Path:    rel,
				Message: "referenced source file missing: " + rel,
			})
		}
	}
	for _, targetName := range manifest.EnabledTargets() {
		entry, ok := targetcontracts.Lookup(targetName)
		if !ok {
			continue
		}
		tc := graph.Targets[targetName]
		supportedPortable := setOf(entry.PortableComponentKinds)
		if len(graph.Portable.Skills) > 0 && !supportedPortable["skills"] {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureUnsupportedTargetKind,
				Path:    "skills",
				Target:  targetName,
				Message: fmt.Sprintf("target %s does not support portable component kind skills", targetName),
			})
		}
		if graph.Portable.MCP != nil && !supportedPortable["mcp_servers"] {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureUnsupportedTargetKind,
				Path:    "mcp",
				Target:  targetName,
				Message: fmt.Sprintf("target %s does not support portable component kind mcp_servers", targetName),
			})
		}
		if len(graph.Portable.Agents) > 0 && !supportedPortable["agents"] {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureUnsupportedTargetKind,
				Path:    "agents",
				Target:  targetName,
				Message: fmt.Sprintf("target %s does not support portable component kind agents", targetName),
			})
		}
		supportedNative := setOf(entry.TargetComponentKinds)
		for _, kind := range discoveredTargetKinds(tc) {
			if supportedNative[kind] {
				continue
			}
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureUnsupportedTargetKind,
				Path:    kind,
				Target:  targetName,
				Message: fmt.Sprintf("target %s does not support target-native component kind %s", targetName, kind),
			})
		}
		if targetName == "gemini" {
			validateGeminiTarget(root, manifest, graph, tc, &report)
		}
	}
	if drift, err := pluginmanifest.Drift(root, targetOrAll(platform)); err != nil {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureGeneratedContractInvalid,
			Message: err.Error(),
		})
	} else {
		for _, rel := range drift {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureGeneratedContractInvalid,
				Path:    rel,
				Message: "generated artifact drift: " + rel,
			})
		}
	}
	validatePluginRuntimeFiles(root, manifest, &report)
	return report, nil
}

func mapManifestWarningKind(kind pluginmanifest.WarningKind) WarningKind {
	switch kind {
	case pluginmanifest.WarningUnknownField:
		return WarningManifestUnknownField
	default:
		return WarningManifestUnknownField
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func validatePluginRuntimeFiles(root string, manifest pluginmanifest.Manifest, report *Report) {
	switch manifest.Runtime {
	case "go":
		validateRuntimeFileExists(root, "go.mod", report)
		cmd := exec.Command("go", "build", "./...")
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GOWORK=off")
		if out, err := cmd.CombinedOutput(); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureBuildFailed,
				Target:  "./...",
				Message: fmt.Sprintf("%v\n%s", err, out),
			})
		}
	case "python":
		validatePluginLauncher(root, manifest, report)
		validateRuntimeFileExists(root, "src/main.py", report)
		if err := validatePythonRuntime(root); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureRuntimeNotFound,
				Message: err.Error(),
			})
		}
	case "node":
		validatePluginLauncher(root, manifest, report)
		validateRuntimeFileExists(root, "package.json", report)
		validateNodeRuntimeTarget(root, manifest.Entrypoint, report)
		if err := validateNodeRuntime(); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureRuntimeNotFound,
				Message: err.Error(),
			})
		}
	case "shell":
		validatePluginLauncher(root, manifest, report)
		validateRuntimeTargetExecutable(root, "scripts/main.sh", report)
		if runtime.GOOS == "windows" {
			if _, err := exec.LookPath("bash"); err != nil {
				report.Failures = append(report.Failures, Failure{
					Kind:    FailureRuntimeNotFound,
					Message: "runtime not found: bash (shell runtime on Windows requires bash in PATH; install Git Bash or another bash-compatible shell)",
				})
			}
		}
	}
}

func validatePluginLauncher(root string, manifest pluginmanifest.Manifest, report *Report) {
	info, err := statLauncher(root, manifest.Entrypoint)
	if err != nil {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureLauncherInvalid,
			Message: "launcher invalid: missing " + manifest.Entrypoint,
		})
		return
	}
	if runtime.GOOS != "windows" && info.Mode()&0o111 == 0 {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureLauncherInvalid,
			Message: "launcher invalid: not executable " + manifest.Entrypoint,
		})
	}
}

func targetOrAll(platform string) string {
	if strings.TrimSpace(platform) == "" {
		return "all"
	}
	return platform
}

func discoveredTargetKinds(tc pluginmanifest.TargetComponents) []string {
	var kinds []string
	if tc.PackagePath != "" {
		kinds = append(kinds, "package_metadata")
	}
	if len(tc.Hooks) > 0 {
		kinds = append(kinds, "hooks")
	}
	if len(tc.Commands) > 0 {
		kinds = append(kinds, "commands")
	}
	if len(tc.Policies) > 0 {
		kinds = append(kinds, "policies")
	}
	if len(tc.Themes) > 0 {
		kinds = append(kinds, "themes")
	}
	if len(tc.Settings) > 0 {
		kinds = append(kinds, "settings")
	}
	if len(tc.Contexts) > 0 {
		kinds = append(kinds, "contexts")
	}
	if strings.TrimSpace(tc.ManifestExtra) != "" {
		kinds = append(kinds, "manifest_extra")
	}
	return kinds
}

func validateGeminiTarget(root string, manifest pluginmanifest.Manifest, graph pluginmanifest.PackageGraph, tc pluginmanifest.TargetComponents, report *Report) {
	if err := pluginmanifest.ValidateGeminiExtensionName(manifest.Name); err != nil {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureManifestInvalid,
			Path:    pluginmanifest.FileName,
			Target:  "gemini",
			Message: err.Error(),
		})
	}
	if base := filepath.Base(filepath.Clean(root)); base != manifest.Name {
		report.Warnings = append(report.Warnings, Warning{
			Kind:    WarningGeminiDirNameMismatch,
			Path:    root,
			Message: fmt.Sprintf("Gemini extension directory basename %q does not match extension name %q", base, manifest.Name),
		})
	}
	validateGeminiMCP(graph, report)
	validateGeminiContext(graph, tc, report)
	validateGeminiSettings(root, tc.Settings, report)
	validateGeminiThemes(root, tc.Themes, report)
	validateGeminiManifestExtra(root, tc, report)
	validateGeminiPolicies(root, tc.Policies, report)
	validateGeminiCommands(root, tc.Commands, report)
	validateGeminiJSONFileKinds(root, tc.Hooks, report)
}

func validateGeminiMCP(graph pluginmanifest.PackageGraph, report *Report) {
	if graph.Portable.MCP == nil {
		return
	}
	for serverName, raw := range graph.Portable.MCP.Servers {
		server, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		if _, blocked := server["trust"]; blocked {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    graph.Portable.MCP.Path,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini extension MCP server %q may not set trust", serverName),
			})
		}
		command, _ := server["command"].(string)
		_, hasArgs := server["args"]
		if strings.Contains(strings.TrimSpace(command), " ") && !hasArgs {
			report.Warnings = append(report.Warnings, Warning{
				Kind:    WarningGeminiMCPCommandStyle,
				Path:    graph.Portable.MCP.Path,
				Message: fmt.Sprintf("Gemini extension MCP server %q uses a space-delimited command string; prefer command plus args", serverName),
			})
		}
	}
}

func validateGeminiContext(graph pluginmanifest.PackageGraph, tc pluginmanifest.TargetComponents, report *Report) {
	selected := strings.TrimSpace(tc.Gemini.ContextFileName)
	candidates := geminiContextMatches(graph, tc, "")
	if selected != "" {
		matches := geminiContextMatches(graph, tc, selected)
		switch len(matches) {
		case 0:
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    tc.PackagePath,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini context_file_name %q does not resolve to a shared or Gemini-native context source", selected),
			})
		case 1:
			return
		default:
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    tc.PackagePath,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini context_file_name %q is ambiguous across multiple context sources", selected),
			})
		}
		return
	}
	geminiMD := geminiContextMatches(graph, tc, "GEMINI.md")
	if len(geminiMD) > 1 {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureManifestInvalid,
			Path:    "contexts",
			Target:  "gemini",
			Message: "Gemini primary context selection is ambiguous for GEMINI.md; keep one root context or set context_file_name explicitly",
		})
		return
	}
	if len(geminiMD) == 1 || len(candidates) <= 1 {
		return
	}
	report.Failures = append(report.Failures, Failure{
		Kind:    FailureManifestInvalid,
		Path:    "contexts",
		Target:  "gemini",
		Message: "Gemini primary context selection is ambiguous; set targets/gemini/package.yaml context_file_name explicitly",
	})
}

func geminiContextMatches(graph pluginmanifest.PackageGraph, tc pluginmanifest.TargetComponents, name string) []string {
	var matches []string
	seen := map[string]struct{}{}
	for _, rel := range append(append([]string{}, tc.Contexts...), graph.Portable.Contexts...) {
		rel = filepath.ToSlash(rel)
		if name == "" || filepath.Base(rel) == name {
			if _, ok := seen[rel]; ok {
				continue
			}
			seen[rel] = struct{}{}
			matches = append(matches, rel)
		}
	}
	slices.Sort(matches)
	return matches
}

func validateGeminiSettings(root string, rels []string, report *Report) {
	for _, rel := range rels {
		body, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini setting file %s is not readable: %v", rel, err),
			})
			continue
		}
		var raw map[string]any
		if err := yaml.Unmarshal(body, &raw); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini setting file %s is invalid YAML: %v", rel, err),
			})
			continue
		}
		var setting pluginmanifest.GeminiSetting
		if err := yaml.Unmarshal(body, &setting); err != nil {
			continue
		}
		_, hasSensitive := raw["sensitive"]
		if strings.TrimSpace(setting.Name) == "" || strings.TrimSpace(setting.Description) == "" || strings.TrimSpace(setting.EnvVar) == "" || !hasSensitive {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini setting file %s must define name, description, env_var, and sensitive", rel),
			})
		}
	}
}

func validateGeminiThemes(root string, rels []string, report *Report) {
	for _, rel := range rels {
		body, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini theme file %s is not readable: %v", rel, err),
			})
			continue
		}
		var raw map[string]any
		if err := yaml.Unmarshal(body, &raw); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini theme file %s is invalid YAML: %v", rel, err),
			})
			continue
		}
		name, _ := raw["name"].(string)
		if strings.TrimSpace(name) == "" {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini theme file %s must define name", rel),
			})
		}
	}
}

func validateGeminiManifestExtra(root string, tc pluginmanifest.TargetComponents, report *Report) {
	if strings.TrimSpace(tc.ManifestExtra) == "" {
		return
	}
	body, err := os.ReadFile(filepath.Join(root, tc.ManifestExtra))
	if err != nil {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureManifestInvalid,
			Path:    tc.ManifestExtra,
			Target:  "gemini",
			Message: fmt.Sprintf("Gemini manifest extra file %s is not readable: %v", tc.ManifestExtra, err),
		})
		return
	}
	var extra map[string]any
	if err := json.Unmarshal(body, &extra); err != nil {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureManifestInvalid,
			Path:    tc.ManifestExtra,
			Target:  "gemini",
			Message: fmt.Sprintf("Gemini manifest extra file %s must be a JSON object: %v", tc.ManifestExtra, err),
		})
		return
	}
	forbidden := map[string]struct{}{
		"name": {}, "version": {}, "description": {}, "mcpServers": {}, "contextFileName": {}, "excludeTools": {}, "migratedTo": {}, "settings": {}, "themes": {},
	}
	for key, value := range extra {
		if _, blocked := forbidden[key]; blocked {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    tc.ManifestExtra,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini manifest extra file may not override canonical field %q", key),
			})
		}
		if key == "plan" {
			plan, ok := value.(map[string]any)
			if !ok {
				report.Failures = append(report.Failures, Failure{
					Kind:    FailureManifestInvalid,
					Path:    tc.ManifestExtra,
					Target:  "gemini",
					Message: "Gemini manifest extra field plan must be a JSON object",
				})
				continue
			}
			if _, blocked := plan["directory"]; blocked {
				report.Failures = append(report.Failures, Failure{
					Kind:    FailureManifestInvalid,
					Path:    tc.ManifestExtra,
					Target:  "gemini",
					Message: "Gemini manifest extra file may not override canonical field \"plan.directory\"",
				})
			}
		}
	}
}

func validateGeminiPolicies(root string, rels []string, report *Report) {
	for _, rel := range rels {
		body, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini policy file %s is not readable: %v", rel, err),
			})
			continue
		}
		text := string(body)
		for _, key := range []string{"allow", "yolo"} {
			if strings.Contains(text, key+" =") {
				report.Warnings = append(report.Warnings, Warning{
					Kind:    WarningGeminiPolicyIgnored,
					Path:    rel,
					Message: fmt.Sprintf("Gemini extension policies ignore %q at extension tier", key),
				})
			}
		}
	}
}

func validateGeminiCommands(root string, rels []string, report *Report) {
	for _, rel := range rels {
		if filepath.Ext(rel) != ".toml" {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini command file %s must use the .toml extension", rel),
			})
			continue
		}
		body, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini command file %s is not readable: %v", rel, err),
			})
			continue
		}
		var discard map[string]any
		if err := toml.Unmarshal(body, &discard); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini command file %s is invalid TOML: %v", rel, err),
			})
		}
	}
}

func validateGeminiJSONFileKinds(root string, rels []string, report *Report) {
	for _, rel := range rels {
		body, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini JSON asset %s is not readable: %v", rel, err),
			})
			continue
		}
		var discard any
		if err := json.Unmarshal(body, &discard); err != nil {
			report.Failures = append(report.Failures, Failure{
				Kind:    FailureManifestInvalid,
				Path:    rel,
				Target:  "gemini",
				Message: fmt.Sprintf("Gemini JSON asset %s is invalid JSON: %v", rel, err),
			})
		}
	}
}

func setOf(values []string) map[string]bool {
	out := make(map[string]bool, len(values))
	for _, value := range values {
		out[value] = true
	}
	return out
}

func statLauncher(root, entrypoint string) (os.FileInfo, error) {
	rel := strings.TrimPrefix(filepath.Clean(entrypoint), "./")
	candidates := []string{filepath.Join(root, rel)}
	if runtime.GOOS == "windows" {
		candidates = append(candidates, filepath.Join(root, rel+".cmd"))
	}
	for _, candidate := range candidates {
		info, err := os.Stat(candidate)
		if err == nil {
			return info, nil
		}
		if !os.IsNotExist(err) {
			return nil, err
		}
	}
	return nil, os.ErrNotExist
}

func validateRuntimeFileExists(root, rel string, report *Report) {
	if _, err := os.Stat(filepath.Join(root, rel)); err != nil {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureRuntimeTargetMissing,
			Path:    rel,
			Message: "runtime target missing: " + rel,
		})
	}
}

func validateRuntimeTargetExecutable(root, rel string, report *Report) {
	info, err := os.Stat(filepath.Join(root, rel))
	if err != nil {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureRuntimeTargetMissing,
			Path:    rel,
			Message: "runtime target missing: " + rel,
		})
		return
	}
	if runtime.GOOS != "windows" && info.Mode()&0o111 == 0 {
		report.Failures = append(report.Failures, Failure{
			Kind:    FailureRuntimeTargetMissing,
			Path:    rel,
			Message: "runtime target missing: " + rel + " is not executable",
		})
	}
}

type pythonRuntimeResolution struct {
	Path   string
	Source string
}

func validatePythonRuntime(root string) error {
	resolution, err := findPython(root)
	if err != nil {
		return err
	}
	out, err := exec.Command(resolution.Path, "--version").CombinedOutput()
	if err != nil {
		switch resolution.Source {
		case "project-venv":
			return fmt.Errorf("runtime not found: found project virtualenv interpreter at %s but it is not runnable (%v); recreate .venv or install Python 3.10+", resolution.Path, err)
		default:
			return fmt.Errorf("runtime not found: found %s at %s but it is not runnable (%v); install Python 3.10+ or repair your PATH", resolution.Source, resolution.Path, err)
		}
	}
	if err := requireMinVersion("python", string(out), 3, 10); err != nil {
		switch resolution.Source {
		case "project-venv":
			return fmt.Errorf("runtime not found: found project virtualenv interpreter at %s but %v; recreate .venv with Python 3.10+ or repair the virtualenv", resolution.Path, err)
		default:
			return fmt.Errorf("runtime not found: found %s at %s but %v; install Python 3.10+ or repair your PATH", resolution.Source, resolution.Path, err)
		}
	}
	return nil
}

func findPython(root string) (pythonRuntimeResolution, error) {
	candidates := pythonCandidates(root)
	venvExists := fileExists(filepath.Join(root, ".venv")) || dirExists(filepath.Join(root, ".venv"))
	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			return pythonRuntimeResolution{Path: candidate, Source: "project-venv"}, nil
		}
	}
	checkedVenv := strings.Join(candidates, ", ")
	checkedPath := strings.Join(pythonPathNames(), ", ")
	if venvExists {
		return pythonRuntimeResolution{}, fmt.Errorf("runtime not found: python runtime required; checked project virtualenv (%s); found .venv but no runnable interpreter. Recreate .venv or install Python 3.10+", checkedVenv)
	}
	for _, name := range pythonPathNames() {
		path, err := exec.LookPath(name)
		if err == nil {
			return pythonRuntimeResolution{Path: path, Source: "system-path"}, nil
		}
	}
	return pythonRuntimeResolution{}, fmt.Errorf("runtime not found: python runtime required; checked PATH runtimes (%s). Install Python 3.10+ or create .venv with python3 -m venv .venv", checkedPath)
}

func pythonCandidates(root string) []string {
	if runtime.GOOS == "windows" {
		return []string{
			filepath.Join(root, ".venv", "Scripts", "python.exe"),
			filepath.Join(root, ".venv", "bin", "python3"),
		}
	}
	return []string{
		filepath.Join(root, ".venv", "bin", "python3"),
		filepath.Join(root, ".venv", "Scripts", "python.exe"),
	}
}

func pythonPathNames() []string {
	if runtime.GOOS == "windows" {
		return []string{"python", "python3"}
	}
	return []string{"python3"}
}

func validateNodeRuntime() error {
	path, err := exec.LookPath("node")
	if err != nil {
		return fmt.Errorf("runtime not found: node runtime required; checked PATH for node. Install Node.js 20+")
	}
	out, err := exec.Command(path, "--version").CombinedOutput()
	if err != nil {
		return fmt.Errorf("runtime not found: found node at %s but it is not runnable (%v); install or repair Node.js 20+", path, err)
	}
	if err := requireMinVersion("node", string(out), 20, 0); err != nil {
		return fmt.Errorf("runtime not found: found node at %s but %v; install or repair Node.js 20+", path, err)
	}
	return nil
}

func validateNodeRuntimeTarget(root, entrypoint string, report *Report) {
	rel := detectNodeRuntimeTarget(root, entrypoint)
	full := filepath.Join(root, filepath.FromSlash(rel))
	if _, err := os.Stat(full); err == nil {
		return
	}
	message := "runtime target missing: " + rel
	if strings.HasPrefix(rel, "dist/") || strings.HasPrefix(rel, "build/") {
		message += " (launcher points to built output; run npm install && npm run build, or restore the launcher target)"
	} else {
		message += " (restore the generated scaffold target or update the launcher)"
	}
	report.Failures = append(report.Failures, Failure{
		Kind:    FailureRuntimeTargetMissing,
		Path:    rel,
		Message: message,
	})
}

func detectNodeRuntimeTarget(root, entrypoint string) string {
	body, err := os.ReadFile(launcherPath(root, entrypoint))
	if err != nil {
		return "src/main.mjs"
	}
	text := filepath.ToSlash(string(body))
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`\$ROOT/([^"\s]+\.(?:mjs|js))`),
		regexp.MustCompile(`%ROOT%/([^"\r\n]+\.(?:mjs|js))`),
	}
	for _, pattern := range patterns {
		matches := pattern.FindStringSubmatch(text)
		if len(matches) == 2 {
			return matches[1]
		}
	}
	return "src/main.mjs"
}

func launcherPath(root, entrypoint string) string {
	rel := strings.TrimPrefix(filepath.Clean(entrypoint), "./")
	full := filepath.Join(root, rel)
	if runtime.GOOS == "windows" {
		if _, err := os.Stat(full + ".cmd"); err == nil {
			return full + ".cmd"
		}
	}
	return full
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

var versionPattern = regexp.MustCompile(`(\d+)\.(\d+)`)

func requireMinVersion(runtimeName, output string, wantMajor, wantMinor int) error {
	major, minor, err := parseMajorMinor(output)
	if err != nil {
		return fmt.Errorf("reported unsupported version output %q", strings.TrimSpace(output))
	}
	if major > wantMajor || (major == wantMajor && minor >= wantMinor) {
		return nil
	}
	return fmt.Errorf("reported version %d.%d is below the supported minimum %d.%d", major, minor, wantMajor, wantMinor)
}

func parseMajorMinor(output string) (int, int, error) {
	matches := versionPattern.FindStringSubmatch(strings.TrimSpace(output))
	if len(matches) != 3 {
		return 0, 0, fmt.Errorf("no major.minor version found")
	}
	major, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, err
	}
	minor, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, err
	}
	return major, minor, nil
}
