package platformexec

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/pluginmodel"
)

type codexAdapter struct{}

func (codexAdapter) ID() string { return "codex" }

func (codexAdapter) DetectNative(root string) bool {
	return fileExists(filepath.Join(root, ".codex", "config.toml")) || fileExists(filepath.Join(root, ".codex-plugin", "plugin.json"))
}

func (codexAdapter) RefineDiscovery(root string, state *pluginmodel.TargetState) error {
	if rel := state.DocPath("package_metadata"); strings.TrimSpace(rel) != "" {
		if _, ok, err := readYAMLDoc[codexPackageMeta](root, rel); err != nil {
			return fmt.Errorf("parse %s: %w", rel, err)
		} else if !ok {
			return nil
		}
	}
	return nil
}

func (codexAdapter) Import(root string, seed ImportSeed) (ImportResult, error) {
	result := ImportResult{
		Manifest: seed.Manifest,
		Launcher: seed.Launcher,
	}
	config, _, err := readImportedCodexConfig(root)
	if err != nil {
		if !os.IsNotExist(err) {
			return ImportResult{}, err
		}
	} else {
		if strings.TrimSpace(config.Model) != "" {
			result.Artifacts = append(result.Artifacts, pluginmodel.Artifact{
				RelPath: filepath.Join("targets", "codex", "package.yaml"),
				Content: mustYAML(codexPackageMeta{ModelHint: config.Model}),
			})
		}
		if len(config.Extra) > 0 {
			body, err := toml.Marshal(config.Extra)
			if err != nil {
				return ImportResult{}, err
			}
			result.Artifacts = append(result.Artifacts, pluginmodel.Artifact{RelPath: filepath.Join("targets", "codex", "config.extra.toml"), Content: body})
			result.Warnings = append(result.Warnings, pluginmodel.Warning{
				Kind:    pluginmodel.WarningFidelity,
				Path:    filepath.ToSlash(filepath.Join("targets", "codex", "config.extra.toml")),
				Message: "preserved unsupported Codex config fields under targets/codex/config.extra.toml",
			})
		}
		if len(config.Notify) > 0 && result.Launcher != nil && strings.TrimSpace(config.Notify[0]) != "" {
			result.Launcher.Entrypoint = strings.TrimSpace(config.Notify[0])
		}
		if len(config.Notify) > 0 && !pluginmodel.IsCanonicalCodexNotify(config.Notify) {
			result.Warnings = append(result.Warnings, pluginmodel.Warning{
				Kind:    pluginmodel.WarningFidelity,
				Path:    filepath.ToSlash(filepath.Join(".codex", "config.toml")),
				Message: "normalized Codex notify argv to the managed [entrypoint, \"notify\"] shape",
			})
		}
	}
	pluginManifest, _, err := readImportedCodexPluginManifest(root)
	if err != nil {
		if !os.IsNotExist(err) {
			return ImportResult{}, err
		}
	} else {
		if strings.TrimSpace(pluginManifest.Name) != "" {
			result.Manifest.Name = pluginManifest.Name
		}
		if strings.TrimSpace(pluginManifest.Version) != "" {
			result.Manifest.Version = pluginManifest.Version
		}
		if strings.TrimSpace(pluginManifest.Description) != "" {
			result.Manifest.Description = pluginManifest.Description
		}
		if len(pluginManifest.Extra) > 0 {
			result.Artifacts = append(result.Artifacts, pluginmodel.Artifact{RelPath: filepath.Join("targets", "codex", "manifest.extra.json"), Content: mustJSON(pluginManifest.Extra)})
			result.Warnings = append(result.Warnings, pluginmodel.Warning{
				Kind:    pluginmodel.WarningFidelity,
				Path:    filepath.ToSlash(filepath.Join("targets", "codex", "manifest.extra.json")),
				Message: "preserved unsupported Codex plugin manifest fields under targets/codex/manifest.extra.json",
			})
		}
		if strings.TrimSpace(pluginManifest.SkillsPath) != "" && strings.TrimSpace(pluginManifest.SkillsPath) != "./skills/" {
			result.Warnings = append(result.Warnings, pluginmodel.Warning{
				Kind:    pluginmodel.WarningFidelity,
				Path:    filepath.ToSlash(filepath.Join(".codex-plugin", "plugin.json")),
				Message: "normalized Codex plugin skills path to the managed ./skills/ location",
			})
		}
		if strings.TrimSpace(pluginManifest.MCPServersRef) != "" && strings.TrimSpace(pluginManifest.MCPServersRef) != "./.mcp.json" {
			result.Warnings = append(result.Warnings, pluginmodel.Warning{
				Kind:    pluginmodel.WarningFidelity,
				Path:    filepath.ToSlash(filepath.Join(".codex-plugin", "plugin.json")),
				Message: "normalized Codex plugin mcpServers path to the managed ./.mcp.json location",
			})
		}
	}
	if fileExists(filepath.Join(root, "agents")) {
		result.Warnings = append(result.Warnings, pluginmodel.Warning{
			Kind:    pluginmodel.WarningIgnoredImport,
			Path:    "agents",
			Message: "ignored unsupported import asset: agents",
		})
	}
	result.Artifacts = compactArtifacts(result.Artifacts)
	return result, nil
}

func (codexAdapter) Render(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]pluginmodel.Artifact, error) {
	entrypoint := ""
	if graph.Launcher != nil {
		entrypoint = graph.Launcher.Entrypoint
	}
	if strings.TrimSpace(entrypoint) == "" {
		return nil, fmt.Errorf("required launcher missing: %s", pluginmodel.LauncherFileName)
	}
	manifestExtra, err := loadNativeExtraDoc(root, state, "manifest_extra", pluginmodel.NativeDocFormatJSON)
	if err != nil {
		return nil, err
	}
	artifacts, err := renderManagedPluginArtifacts(graph.Manifest.Name, graph.Manifest, graph.Portable, false, filepath.Join(".codex-plugin", "plugin.json"), manifestExtra, "codex manifest.extra.json", []string{"name", "version", "description", "skills", "mcpServers"})
	if err != nil {
		return nil, err
	}
	model := ""
	if meta, ok, err := readYAMLDoc[codexPackageMeta](root, state.DocPath("package_metadata")); err != nil {
		return nil, fmt.Errorf("parse %s: %w", state.DocPath("package_metadata"), err)
	} else if ok {
		model = strings.TrimSpace(meta.ModelHint)
	}
	if strings.TrimSpace(model) == "" {
		model = "gpt-5.4-mini"
	}
	configExtra, err := loadNativeExtraDoc(root, state, "config_extra", pluginmodel.NativeDocFormatTOML)
	if err != nil {
		return nil, err
	}
	if err := pluginmodel.ValidateNativeExtraDocConflicts(configExtra, "codex config.extra.toml", []string{"model", "notify"}); err != nil {
		return nil, err
	}
	var config bytes.Buffer
	config.WriteString("# Generated by plugin-kit-ai. DO NOT EDIT.\n")
	config.WriteString(fmt.Sprintf("model = %q\n", model))
	config.WriteString(fmt.Sprintf("notify = [%q, %q]\n", entrypoint, "notify"))
	if extraBody := pluginmodel.TrimmedExtraBody(configExtra); len(extraBody) > 0 {
		config.WriteByte('\n')
		config.Write(extraBody)
		config.WriteByte('\n')
	}
	artifacts = append(artifacts, pluginmodel.Artifact{RelPath: filepath.Join(".codex", "config.toml"), Content: config.Bytes()})
	copied, err := copyArtifactDirs(root,
		artifactDir{src: filepath.Join("targets", "codex", "commands"), dst: "commands"},
		artifactDir{src: filepath.Join("targets", "codex", "contexts"), dst: "contexts"},
	)
	if err != nil {
		return nil, err
	}
	return append(artifacts, copied...), nil
}

func (codexAdapter) ManagedPaths(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]string, error) {
	return nil, nil
}

func (codexAdapter) Validate(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]Diagnostic, error) {
	body, err := os.ReadFile(filepath.Join(root, ".codex", "config.toml"))
	if err != nil {
		return []Diagnostic{{
			Severity: SeverityFailure,
			Code:     CodeGeneratedContractInvalid,
			Path:     filepath.ToSlash(filepath.Join(".codex", "config.toml")),
			Target:   "codex",
			Message:  fmt.Sprintf("Codex config file %s is not readable: %v", filepath.ToSlash(filepath.Join(".codex", "config.toml")), err),
		}}, nil
	}
	var config struct {
		Model  string   `toml:"model"`
		Notify []string `toml:"notify"`
	}
	if err := toml.Unmarshal(body, &config); err != nil {
		return []Diagnostic{{
			Severity: SeverityFailure,
			Code:     CodeManifestInvalid,
			Path:     filepath.ToSlash(filepath.Join(".codex", "config.toml")),
			Target:   "codex",
			Message:  fmt.Sprintf("Codex config file %s is invalid TOML: %v", filepath.ToSlash(filepath.Join(".codex", "config.toml")), err),
		}}, nil
	}
	if graph.Launcher == nil {
		return nil, nil
	}
	var diagnostics []Diagnostic
	expectedNotify := []string{graph.Launcher.Entrypoint, "notify"}
	if len(config.Notify) != len(expectedNotify) || len(config.Notify) == 0 || strings.TrimSpace(config.Notify[0]) != expectedNotify[0] || (len(config.Notify) > 1 && strings.TrimSpace(config.Notify[1]) != expectedNotify[1]) {
		diagnostics = append(diagnostics, Diagnostic{
			Severity: SeverityFailure,
			Code:     CodeEntrypointMismatch,
			Path:     filepath.ToSlash(filepath.Join(".codex", "config.toml")),
			Target:   "codex",
			Message:  fmt.Sprintf("entrypoint mismatch: Codex notify argv uses %q; expected %q from launcher.yaml entrypoint", config.Notify, expectedNotify),
		})
	}
	if meta, ok, err := readYAMLDoc[codexPackageMeta](root, state.DocPath("package_metadata")); err != nil {
		return nil, fmt.Errorf("parse %s: %w", state.DocPath("package_metadata"), err)
	} else if ok && strings.TrimSpace(meta.ModelHint) != "" && strings.TrimSpace(config.Model) != strings.TrimSpace(meta.ModelHint) {
		diagnostics = append(diagnostics, Diagnostic{
			Severity: SeverityFailure,
			Code:     CodeGeneratedContractInvalid,
			Path:     filepath.ToSlash(filepath.Join(".codex", "config.toml")),
			Target:   "codex",
			Message:  fmt.Sprintf("Codex config model %q does not match targets/codex/package.yaml model_hint %q", strings.TrimSpace(config.Model), strings.TrimSpace(meta.ModelHint)),
		})
	}
	return diagnostics, nil
}
