package platformexec

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/pluginmodel"
)

type opencodeAdapter struct{}

func (opencodeAdapter) ID() string { return "opencode" }

func (opencodeAdapter) DetectNative(root string) bool {
	return fileExists(filepath.Join(root, "opencode.json"))
}

func (opencodeAdapter) RefineDiscovery(root string, state *pluginmodel.TargetState) error {
	if rel := strings.TrimSpace(state.DocPath("package_metadata")); rel != "" {
		meta, ok, err := readYAMLDoc[opencodePackageMeta](root, rel)
		if err != nil {
			return fmt.Errorf("parse %s: %w", rel, err)
		}
		if ok {
			for i, plugin := range meta.Plugins {
				if strings.TrimSpace(plugin) == "" {
					return fmt.Errorf("%s plugin entry %d must be a non-empty string", rel, i)
				}
			}
		}
	}
	return nil
}

func (opencodeAdapter) Import(root string, seed ImportSeed) (ImportResult, error) {
	result := ImportResult{
		Manifest: seed.Manifest,
		Launcher: nil,
		Artifacts: []pluginmodel.Artifact{{
			RelPath: filepath.Join("targets", "opencode", "package.yaml"),
			Content: mustYAML(opencodePackageMeta{}),
		}},
	}
	config, ok, err := readImportedOpenCodeConfig(root)
	if err != nil {
		return ImportResult{}, err
	}
	skillArtifacts, err := importedOpenCodeSkillArtifacts(root)
	if err != nil {
		return ImportResult{}, err
	}
	if !ok && len(skillArtifacts) == 0 {
		return ImportResult{}, fmt.Errorf("OpenCode import requires opencode.json or .opencode/skills/**/SKILL.md")
	}
	if ok {
		if len(config.Plugins) > 0 {
			result.Artifacts[0].Content = mustYAML(opencodePackageMeta{Plugins: append([]string(nil), config.Plugins...)})
		}
		if len(config.MCP) > 0 {
			result.Artifacts = append(result.Artifacts, pluginmodel.Artifact{
				RelPath: filepath.Join("mcp", "servers.json"),
				Content: mustJSON(config.MCP),
			})
		}
		if len(config.Extra) > 0 {
			result.Artifacts = append(result.Artifacts, pluginmodel.Artifact{
				RelPath: filepath.Join("targets", "opencode", "config.extra.json"),
				Content: mustJSON(config.Extra),
			})
		}
	}
	result.Artifacts = append(result.Artifacts, skillArtifacts...)
	if fileExists(filepath.Join(root, ".opencode", "plugins")) {
		result.Warnings = append(result.Warnings, pluginmodel.Warning{
			Kind:    pluginmodel.WarningFidelity,
			Path:    filepath.ToSlash(filepath.Join(".opencode", "plugins")),
			Message: "ignored unsupported OpenCode local plugin code under .opencode/plugins; v1 keeps plugin code out of the package-standard contract",
		})
	}
	if fileExists(filepath.Join(root, ".opencode", "package.json")) {
		result.Warnings = append(result.Warnings, pluginmodel.Warning{
			Kind:    pluginmodel.WarningFidelity,
			Path:    filepath.ToSlash(filepath.Join(".opencode", "package.json")),
			Message: "ignored unsupported OpenCode local plugin dependency metadata under .opencode/package.json",
		})
	}
	result.Artifacts = compactArtifacts(result.Artifacts)
	return result, nil
}

func (opencodeAdapter) Render(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]pluginmodel.Artifact, error) {
	meta, _, err := readYAMLDoc[opencodePackageMeta](root, state.DocPath("package_metadata"))
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", state.DocPath("package_metadata"), err)
	}
	for i, plugin := range meta.Plugins {
		if strings.TrimSpace(plugin) == "" {
			return nil, fmt.Errorf("%s plugin entry %d must be a non-empty string", state.DocPath("package_metadata"), i)
		}
		meta.Plugins[i] = strings.TrimSpace(plugin)
	}
	extra, err := loadNativeExtraDoc(root, state, "config_extra", pluginmodel.NativeDocFormatJSON)
	if err != nil {
		return nil, err
	}
	managedPaths := []string{"$schema", "plugin", "mcp"}
	if err := pluginmodel.ValidateNativeExtraDocConflicts(extra, "opencode config.extra.json", managedPaths); err != nil {
		return nil, err
	}
	doc := map[string]any{
		"$schema": "https://opencode.ai/config.json",
	}
	if len(meta.Plugins) > 0 {
		doc["plugin"] = append([]string(nil), meta.Plugins...)
	}
	if graph.Portable.MCP != nil {
		doc["mcp"] = graph.Portable.MCP.Servers
	}
	if err := pluginmodel.MergeNativeExtraObject(doc, extra, "opencode config.extra.json", managedPaths); err != nil {
		return nil, err
	}
	body, err := marshalJSON(doc)
	if err != nil {
		return nil, err
	}
	artifacts := []pluginmodel.Artifact{{
		RelPath: "opencode.json",
		Content: body,
	}}
	skillArtifacts, err := renderPortableSkills(root, graph.Portable.Paths("skills"), ".opencode/skills")
	if err != nil {
		return nil, err
	}
	artifacts = append(artifacts, skillArtifacts...)
	return compactArtifacts(artifacts), nil
}

func (opencodeAdapter) ManagedPaths(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]string, error) {
	return portableSkillManagedPaths(graph.Portable.Paths("skills"), ".opencode/skills"), nil
}

func (opencodeAdapter) Validate(root string, graph pluginmodel.PackageGraph, state pluginmodel.TargetState) ([]Diagnostic, error) {
	var diagnostics []Diagnostic
	meta, _, err := readYAMLDoc[opencodePackageMeta](root, state.DocPath("package_metadata"))
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", state.DocPath("package_metadata"), err)
	}
	for i, plugin := range meta.Plugins {
		if strings.TrimSpace(plugin) == "" {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: SeverityFailure,
				Code:     CodeManifestInvalid,
				Path:     state.DocPath("package_metadata"),
				Target:   "opencode",
				Message:  fmt.Sprintf("OpenCode package metadata plugin entry %d must be a non-empty string", i),
			})
		}
	}
	body, err := os.ReadFile(filepath.Join(root, "opencode.json"))
	if err != nil {
		return []Diagnostic{{
			Severity: SeverityFailure,
			Code:     CodeGeneratedContractInvalid,
			Path:     "opencode.json",
			Target:   "opencode",
			Message:  fmt.Sprintf("OpenCode config %s is not readable: %v", "opencode.json", err),
		}}, nil
	}
	var doc map[string]any
	if err := json.Unmarshal(body, &doc); err != nil {
		return []Diagnostic{{
			Severity: SeverityFailure,
			Code:     CodeManifestInvalid,
			Path:     "opencode.json",
			Target:   "opencode",
			Message:  fmt.Sprintf("OpenCode config %s is invalid JSON: %v", "opencode.json", err),
		}}, nil
	}
	if schema, _ := doc["$schema"].(string); strings.TrimSpace(schema) != "https://opencode.ai/config.json" {
		diagnostics = append(diagnostics, Diagnostic{
			Severity: SeverityFailure,
			Code:     CodeGeneratedContractInvalid,
			Path:     "opencode.json",
			Target:   "opencode",
			Message:  fmt.Sprintf("OpenCode config %s must declare $schema %q", "opencode.json", "https://opencode.ai/config.json"),
		})
	}
	if raw, ok := doc["plugin"]; ok {
		values, ok := raw.([]any)
		if !ok {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: SeverityFailure,
				Code:     CodeManifestInvalid,
				Path:     "opencode.json",
				Target:   "opencode",
				Message:  `OpenCode config field "plugin" must be an array of strings`,
			})
		} else {
			for i, value := range values {
				text, ok := value.(string)
				if !ok || strings.TrimSpace(text) == "" {
					diagnostics = append(diagnostics, Diagnostic{
						Severity: SeverityFailure,
						Code:     CodeManifestInvalid,
						Path:     "opencode.json",
						Target:   "opencode",
						Message:  fmt.Sprintf(`OpenCode config field "plugin" must contain non-empty strings (invalid entry at index %d)`, i),
					})
				}
			}
		}
	}
	if raw, ok := doc["mcp"]; ok {
		if _, ok := raw.(map[string]any); !ok {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: SeverityFailure,
				Code:     CodeManifestInvalid,
				Path:     "opencode.json",
				Target:   "opencode",
				Message:  `OpenCode config field "mcp" must be a JSON object`,
			})
		}
	}
	return diagnostics, nil
}

func importedOpenCodeSkillArtifacts(root string) ([]pluginmodel.Artifact, error) {
	paths := discoverFiles(root, filepath.Join(".opencode", "skills"), func(rel string) bool {
		return strings.HasSuffix(rel, "SKILL.md")
	})
	var artifacts []pluginmodel.Artifact
	for _, rel := range paths {
		body, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			return nil, err
		}
		child, err := filepath.Rel(filepath.FromSlash(filepath.Join(".opencode", "skills")), filepath.FromSlash(rel))
		if err != nil {
			return nil, err
		}
		artifacts = append(artifacts, pluginmodel.Artifact{
			RelPath: filepath.ToSlash(filepath.Join("skills", child)),
			Content: body,
		})
	}
	return compactArtifacts(artifacts), nil
}

func renderPortableSkills(root string, paths []string, outputRoot string) ([]pluginmodel.Artifact, error) {
	var artifacts []pluginmodel.Artifact
	for _, rel := range paths {
		body, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			return nil, err
		}
		child, err := filepath.Rel(filepath.FromSlash("skills"), filepath.FromSlash(rel))
		if err != nil {
			return nil, err
		}
		artifacts = append(artifacts, pluginmodel.Artifact{
			RelPath: filepath.ToSlash(filepath.Join(outputRoot, child)),
			Content: body,
		})
	}
	return compactArtifacts(artifacts), nil
}

func portableSkillManagedPaths(paths []string, outputRoot string) []string {
	out := make([]string, 0, len(paths))
	for _, rel := range paths {
		child, err := filepath.Rel(filepath.FromSlash("skills"), filepath.FromSlash(rel))
		if err != nil {
			continue
		}
		out = append(out, filepath.ToSlash(filepath.Join(outputRoot, child)))
	}
	slices.Sort(out)
	return slices.Compact(out)
}
