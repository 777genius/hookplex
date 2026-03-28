package platformexec

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/pluginmodel"
	skillfs "github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/skills/adapters/filesystem"
	skillsapp "github.com/plugin-kit-ai/plugin-kit-ai/cli/internal/skills/app"
	"github.com/tailscale/hujson"
)

type opencodeAdapter struct{}

func (opencodeAdapter) ID() string { return "opencode" }

func (opencodeAdapter) DetectNative(root string) bool {
	_, _, ok, err := resolveOpenCodeConfigPath(root)
	return err == nil && ok
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
	config, _, configWarnings, ok, err := readImportedOpenCodeConfig(root)
	if err != nil {
		return ImportResult{}, err
	}
	result.Warnings = append(result.Warnings, configWarnings...)
	skillArtifacts, skillWarnings, err := importedOpenCodeSkillArtifacts(root, seed.Explicit)
	if err != nil {
		return ImportResult{}, err
	}
	result.Warnings = append(result.Warnings, skillWarnings...)
	if !ok && len(skillArtifacts) == 0 {
		return ImportResult{}, fmt.Errorf("OpenCode import requires opencode.json, opencode.jsonc, or explicit --from opencode with compatible skill roots")
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
	configPath, warnings, ok, err := resolveOpenCodeConfigPath(root)
	if err != nil {
		return nil, err
	}
	if !ok {
		return []Diagnostic{{
			Severity: SeverityFailure,
			Code:     CodeGeneratedContractInvalid,
			Path:     "opencode.json",
			Target:   "opencode",
			Message:  "OpenCode config opencode.json or opencode.jsonc is required",
		}}, nil
	}
	for _, warning := range warnings {
		diagnostics = append(diagnostics, Diagnostic{
			Severity: SeverityWarning,
			Code:     CodeManifestInvalid,
			Path:     warning.Path,
			Target:   "opencode",
			Message:  warning.Message,
		})
	}
	body, err := os.ReadFile(filepath.Join(root, configPath))
	if err != nil {
		return []Diagnostic{{
			Severity: SeverityFailure,
			Code:     CodeGeneratedContractInvalid,
			Path:     configPath,
			Target:   "opencode",
			Message:  fmt.Sprintf("OpenCode config %s is not readable: %v", configPath, err),
		}}, nil
	}
	body, err = hujson.Standardize(body)
	if err != nil {
		return []Diagnostic{{
			Severity: SeverityFailure,
			Code:     CodeManifestInvalid,
			Path:     configPath,
			Target:   "opencode",
			Message:  fmt.Sprintf("OpenCode config %s is invalid JSON/JSONC: %v", configPath, err),
		}}, nil
	}
	var doc map[string]any
	if err := json.Unmarshal(body, &doc); err != nil {
		return []Diagnostic{{
			Severity: SeverityFailure,
			Code:     CodeManifestInvalid,
			Path:     configPath,
			Target:   "opencode",
			Message:  fmt.Sprintf("OpenCode config %s is invalid JSON/JSONC: %v", configPath, err),
		}}, nil
	}
	if schema, _ := doc["$schema"].(string); strings.TrimSpace(schema) != "https://opencode.ai/config.json" {
		diagnostics = append(diagnostics, Diagnostic{
			Severity: SeverityFailure,
			Code:     CodeGeneratedContractInvalid,
			Path:     configPath,
			Target:   "opencode",
			Message:  fmt.Sprintf("OpenCode config %s must declare $schema %q", configPath, "https://opencode.ai/config.json"),
		})
	}
	if raw, ok := doc["plugin"]; ok {
		values, ok := raw.([]any)
		if !ok {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: SeverityFailure,
				Code:     CodeManifestInvalid,
				Path:     configPath,
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
						Path:     configPath,
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
				Path:     configPath,
				Target:   "opencode",
				Message:  `OpenCode config field "mcp" must be a JSON object`,
			})
		}
	}
	if len(graph.Portable.Paths("skills")) > 0 {
		report, err := (skillsapp.Service{Repo: skillfs.Repository{}}).Validate(skillsapp.ValidateOptions{Root: root})
		if err != nil {
			return nil, err
		}
		for _, failure := range report.Failures {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: SeverityFailure,
				Code:     CodeManifestInvalid,
				Path:     filepath.ToSlash(failure.Path),
				Target:   "opencode",
				Message:  "OpenCode mirrored skill is incompatible with the shared SKILL.md contract: " + failure.Message,
			})
		}
	}
	return diagnostics, nil
}

func importedOpenCodeSkillArtifacts(root string, explicit bool) ([]pluginmodel.Artifact, []pluginmodel.Warning, error) {
	type importRoot struct {
		dir       string
		warnOnUse bool
		warnPath  string
		warnMsg   string
	}
	roots := []importRoot{
		{dir: filepath.Join(".opencode", "skills")},
	}
	if explicit {
		roots = append(roots,
			importRoot{
				dir:       filepath.Join(".claude", "skills"),
				warnOnUse: true,
				warnPath:  filepath.ToSlash(filepath.Join(".claude", "skills")),
				warnMsg:   "normalized OpenCode-compatible skills from .claude/skills into canonical portable skills/** during import",
			},
			importRoot{
				dir:       filepath.Join(".agents", "skills"),
				warnOnUse: true,
				warnPath:  filepath.ToSlash(filepath.Join(".agents", "skills")),
				warnMsg:   "normalized OpenCode-compatible skills from .agents/skills into canonical portable skills/** during import",
			},
		)
	}
	var (
		artifacts []pluginmodel.Artifact
		warnings  []pluginmodel.Warning
		seen      = map[string]struct{}{}
	)
	for _, source := range roots {
		paths := discoverFiles(root, source.dir, func(rel string) bool {
			return strings.HasSuffix(rel, "SKILL.md")
		})
		usedRoot := false
		for _, rel := range paths {
			child, err := filepath.Rel(filepath.FromSlash(source.dir), filepath.FromSlash(rel))
			if err != nil {
				return nil, nil, err
			}
			dstRel := filepath.ToSlash(filepath.Join("skills", child))
			if _, ok := seen[dstRel]; ok {
				continue
			}
			body, err := os.ReadFile(filepath.Join(root, rel))
			if err != nil {
				return nil, nil, err
			}
			seen[dstRel] = struct{}{}
			artifacts = append(artifacts, pluginmodel.Artifact{
				RelPath: dstRel,
				Content: body,
			})
			usedRoot = true
		}
		if source.warnOnUse && usedRoot {
			warnings = append(warnings, pluginmodel.Warning{
				Kind:    pluginmodel.WarningFidelity,
				Path:    source.warnPath,
				Message: source.warnMsg,
			})
		}
	}
	return compactArtifacts(artifacts), warnings, nil
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
